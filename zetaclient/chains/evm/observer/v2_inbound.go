package observer

import (
	"bytes"
	"context"
	"encoding/hex"
	"sort"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/protocol-contracts/v2/pkg/gatewayevm.sol"

	"github.com/zeta-chain/zetacore/pkg/coin"
	"github.com/zeta-chain/zetacore/pkg/constant"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"github.com/zeta-chain/zetacore/zetaclient/compliance"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	"github.com/zeta-chain/zetacore/zetaclient/metrics"
	"github.com/zeta-chain/zetacore/zetaclient/zetacore"
)

// ObserveGateway queries the gateway contract for deposit/call events
// returns the last block successfully scanned
func (ob *Observer) ObserveGateway(ctx context.Context, startBlock, toBlock uint64) uint64 {
	// filter ERC20CustodyDeposited logs
	gatewayAddr, gatewayContract, err := ob.GetGatewayContract()
	if err != nil {
		ob.Logger().Inbound.Warn().Err(err).Msgf("ObserveGateway: can't get gateway contract")
		return startBlock - 1 // lastScanned
	}

	ob.Logger().Inbound.Info().Msgf("ObserveGateway: gatewayAddreth %s", gatewayAddr.Hex())

	// get iterator for the events for the block range
	eventIterator, err := gatewayContract.FilterDeposit(&bind.FilterOpts{
		Start:   startBlock,
		End:     &toBlock,
		Context: ctx,
	}, []ethcommon.Address{}, []ethcommon.Address{})
	if err != nil {
		ob.Logger().Inbound.Warn().Err(err).Msgf(
			"ObserveGateway: FilterDeposit error from block %d to %d for chain %d", startBlock, toBlock, ob.Chain().ChainId)
		return startBlock - 1 // lastScanned
	}

	// parse and validate events
	events := ob.parseAndValidateDepositEvents(eventIterator)

	// increment prom counter
	metrics.GetFilterLogsPerChain.WithLabelValues(ob.Chain().Name).Inc()

	// post to zetacore
	lastScanned := uint64(0)
	for _, event := range events {
		// remember which block we are scanning (there could be multiple events in the same block)
		if event.Raw.BlockNumber > lastScanned {
			lastScanned = event.Raw.BlockNumber
		}

		// check if the event is processable
		if !ob.checkEventProcessability(event) {
			continue
		}

		msg := ob.newDepositInboundVote(event)

		ob.Logger().Inbound.Info().
			Msgf("ObserveGateway: Deposit inbound detected on chain %d tx %s block %d from %s value %s message %s",
				ob.Chain().
					ChainId, event.Raw.TxHash.Hex(), event.Raw.BlockNumber, event.Sender.Hex(), event.Amount.String(), hex.EncodeToString(event.Payload))

		_, err = ob.PostVoteInbound(ctx, &msg, zetacore.PostVoteInboundExecutionGasLimit)
		if err != nil {
			// decrement the last scanned block so we have to re-scan from this block next time
			return lastScanned - 1
		}
	}

	// successfully processed all events in [startBlock, toBlock]
	return toBlock
}

// parseAndValidateEvents collects and sorts events by block number, tx index, and log index
func (ob *Observer) parseAndValidateDepositEvents(
	iterator *gatewayevm.GatewayEVMDepositIterator,
) []*gatewayevm.GatewayEVMDeposit {
	// collect and sort events by block number, then tx index, then log index (ascending)
	events := make([]*gatewayevm.GatewayEVMDeposit, 0)
	for iterator.Next() {
		// TODO: implement sanity check tx event
		events = append(events, iterator.Event)
		//err := evm.ValidateEvmTxLog(&eventIterator.Event.Raw, gatewayAddr, "", evm.TopicsDeposited)
		//if err == nil {
		//	events = append(events, eventIterator.Event)
		//	continue
		//}
		//ob.Logger().Inbound.Warn().
		//	Err(err).
		//	Msgf("ObserveGateway: invalid Deposited event in tx %s on chain %d at height %d",
		//		eventIterator.Event.Raw.TxHash.Hex(), ob.Chain().ChainId, eventIterator.Event.Raw.BlockNumber)
	}
	sort.SliceStable(events, func(i, j int) bool {
		if events[i].Raw.BlockNumber == events[j].Raw.BlockNumber {
			if events[i].Raw.TxIndex == events[j].Raw.TxIndex {
				return events[i].Raw.Index < events[j].Raw.Index
			}
			return events[i].Raw.TxIndex < events[j].Raw.TxIndex
		}
		return events[i].Raw.BlockNumber < events[j].Raw.BlockNumber
	})

	// filter events from same tx
	filtered := make([]*gatewayevm.GatewayEVMDeposit, 0)
	guard := make(map[string]bool)
	for _, event := range events {
		// guard against multiple events in the same tx
		if guard[event.Raw.TxHash.Hex()] {
			ob.Logger().Inbound.Warn().
				Msgf("ObserveGateway: multiple remote call events detected in same tx %s", event.Raw.TxHash)
			continue
		}
		guard[event.Raw.TxHash.Hex()] = true
		filtered = append(filtered, event)
	}

	return filtered
}

// checkEventProcessability checks if the event is processable
func (ob *Observer) checkEventProcessability(event *gatewayevm.GatewayEVMDeposit) bool {
	// compliance check
	if config.ContainRestrictedAddress(event.Sender.Hex(), event.Receiver.Hex()) {
		compliance.PrintComplianceLog(
			ob.Logger().Inbound,
			ob.Logger().Compliance,
			false,
			ob.Chain().ChainId,
			event.Raw.TxHash.Hex(),
			event.Sender.Hex(),
			event.Receiver.Hex(),
			"Deposit",
		)
		return false
	}

	// donation check
	if bytes.Equal(event.Payload, []byte(constant.DonationMessage)) {
		ob.Logger().Inbound.Info().
			Msgf("thank you rich folk for your donation! tx %s chain %d", event.Raw.TxHash.Hex(), ob.Chain().ChainId)
		return false
	}

	return true
}

// newDepositInboundVote creates a MsgVoteInbound message for a Deposit event
func (ob *Observer) newDepositInboundVote(event *gatewayevm.GatewayEVMDeposit) types.MsgVoteInbound {
	return *types.NewMsgVoteInbound(
		ob.ZetacoreClient().GetKeys().GetOperatorAddress().String(),
		event.Sender.Hex(),
		ob.Chain().ChainId,
		"",
		event.Receiver.Hex(),
		ob.ZetacoreClient().Chain().ChainId,
		sdkmath.NewUintFromBigInt(event.Amount),
		hex.EncodeToString(event.Payload),
		event.Raw.TxHash.Hex(),
		event.Raw.BlockNumber,
		1_500_000,
		coin.CoinType_Gas,
		"",
		event.Raw.Index,
		types.ProtocolContractVersion_V2,
	)
}
