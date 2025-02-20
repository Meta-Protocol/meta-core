package observer

import (
	"bytes"
	"context"
	"encoding/hex"
	"sort"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayevm.sol"

	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/pkg/constant"
	"github.com/zeta-chain/node/pkg/crypto"
	"github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/node/zetaclient/chains/evm/common"
	"github.com/zeta-chain/node/zetaclient/compliance"
	"github.com/zeta-chain/node/zetaclient/config"
	"github.com/zeta-chain/node/zetaclient/logs"
	"github.com/zeta-chain/node/zetaclient/metrics"
	"github.com/zeta-chain/node/zetaclient/zetacore"
)

// isEventProcessable checks if the event is processable
func (ob *Observer) isEventProcessable(
	sender, receiver ethcommon.Address,
	txHash ethcommon.Hash,
	payload []byte,
) bool {
	// compliance check
	if config.ContainRestrictedAddress(sender.Hex(), receiver.Hex()) {
		compliance.PrintComplianceLog(
			ob.Logger().Inbound,
			ob.Logger().Compliance,
			false,
			ob.Chain().ChainId,
			txHash.Hex(),
			sender.Hex(),
			receiver.Hex(),
			"Deposit",
		)
		return false
	}

	// donation check
	if bytes.Equal(payload, []byte(constant.DonationMessage)) {
		logFields := map[string]any{
			"chain": ob.Chain().ChainId,
			"tx":    txHash.Hex(),
		}
		ob.Logger().Inbound.Info().Fields(logFields).
			Msgf("thank you rich folk for your donation!")
		return false
	}

	return true
}

// ObserveGatewayDeposit queries the gateway contract for deposit events
// returns the last block successfully scanned
func (ob *Observer) ObserveGatewayDeposit(ctx context.Context, startBlock, toBlock uint64) (uint64, error) {
	// filter ERC20CustodyDeposited logs
	gatewayAddr, gatewayContract, err := ob.GetGatewayContract()
	if err != nil {
		// lastScanned is startBlock - 1
		return startBlock - 1, errors.Wrap(err, "can't get gateway contract")
	}

	// get iterator for the events for the block range
	eventIterator, err := gatewayContract.FilterDeposited(&bind.FilterOpts{
		Start:   startBlock,
		End:     &toBlock,
		Context: ctx,
	}, []ethcommon.Address{}, []ethcommon.Address{})
	if err != nil {
		return startBlock - 1, errors.Wrapf(
			err,
			"error filtering deposits from block %d to %d for chain %d",
			startBlock,
			toBlock,
			ob.Chain().ChainId,
		)
	}

	// parse and validate events
	events := ob.parseAndValidateDepositEvents(eventIterator, gatewayAddr)

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
		if !ob.isEventProcessable(event.Sender, event.Receiver, event.Raw.TxHash, event.Payload) {
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
			return lastScanned - 1, errors.Wrap(err, "error posting vote inbound")
		}
	}

	// successfully processed all events in [startBlock, toBlock]
	return toBlock, nil
}

// parseAndValidateDepositEvents collects and sorts events by block number, tx index, and log index
func (ob *Observer) parseAndValidateDepositEvents(
	iterator *gatewayevm.GatewayEVMDepositedIterator,
	gatewayAddr ethcommon.Address,
) []*gatewayevm.GatewayEVMDeposited {
	// collect and sort validEvents by block number, then tx index, then log index (ascending)
	validEvents := make([]*gatewayevm.GatewayEVMDeposited, 0)
	for iterator.Next() {
		err := common.ValidateEvmTxLog(&iterator.Event.Raw, gatewayAddr, "", common.TopicsGatewayDeposit)
		if err == nil {
			validEvents = append(validEvents, iterator.Event)
			continue
		}
		ob.Logger().
			Inbound.Warn().
			Stringer(logs.FieldTx, iterator.Event.Raw.TxHash).
			Uint64(logs.FieldBlock, iterator.Event.Raw.BlockNumber).
			Msg("invalid Deposited event")
	}

	// order events by height, tx index and event index (ascending)
	// this ensures the first event is observed if there are multiple in the same tx
	sort.SliceStable(validEvents, func(i, j int) bool {
		if validEvents[i].Raw.BlockNumber == validEvents[j].Raw.BlockNumber {
			if validEvents[i].Raw.TxIndex == validEvents[j].Raw.TxIndex {
				return validEvents[i].Raw.Index < validEvents[j].Raw.Index
			}
			return validEvents[i].Raw.TxIndex < validEvents[j].Raw.TxIndex
		}
		return validEvents[i].Raw.BlockNumber < validEvents[j].Raw.BlockNumber
	})

	// filter events from same tx
	filtered := make([]*gatewayevm.GatewayEVMDeposited, 0)
	guard := make(map[string]bool)
	for _, event := range validEvents {
		// guard against multiple events in the same tx
		if guard[event.Raw.TxHash.Hex()] {
			ob.Logger().
				Inbound.Warn().
				Stringer(logs.FieldTx, event.Raw.TxHash).
				Msg("multiple Deposited events in same tx")
			continue
		}
		guard[event.Raw.TxHash.Hex()] = true
		filtered = append(filtered, event)
	}

	return filtered
}

// newDepositInboundVote creates a MsgVoteInbound message for a Deposit event
func (ob *Observer) newDepositInboundVote(event *gatewayevm.GatewayEVMDeposited) types.MsgVoteInbound {
	// if event.Asset is zero, it's a native token
	coinType := coin.CoinType_ERC20
	if crypto.IsEmptyAddress(event.Asset) {
		coinType = coin.CoinType_Gas
	}

	// to maintain compatibility with previous gateway version, deposit event with a non-empty payload is considered as a call
	isCrossChainCall := false
	if len(event.Payload) > 0 {
		isCrossChainCall = true
	}

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
		zetacore.PostVoteInboundCallOptionsGasLimit,
		coinType,
		event.Asset.Hex(),
		uint64(event.Raw.Index),
		types.ProtocolContractVersion_V2,
		false, // currently not relevant since calls are not arbitrary
		types.InboundStatus_SUCCESS,
		types.ConfirmationMode_SAFE,
		types.WithEVMRevertOptions(event.RevertOptions),
		types.WithCrossChainCall(isCrossChainCall),
	)
}

// ObserveGatewayCall queries the gateway contract for call events
// returns the last block successfully scanned
// TODO: there are lot of similarities between this function and ObserveGatewayDeposit
// logic should be factorized using interfaces and generics
// https://github.com/zeta-chain/node/issues/2493
func (ob *Observer) ObserveGatewayCall(ctx context.Context, startBlock, toBlock uint64) (uint64, error) {
	// filter ERC20CustodyDeposited logs
	gatewayAddr, gatewayContract, err := ob.GetGatewayContract()
	if err != nil {
		// lastScanned is startBlock - 1
		return startBlock - 1, errors.Wrap(err, "can't get gateway contract")
	}

	// get iterator for the events for the block range
	eventIterator, err := gatewayContract.FilterCalled(&bind.FilterOpts{
		Start:   startBlock,
		End:     &toBlock,
		Context: ctx,
	}, []ethcommon.Address{}, []ethcommon.Address{})
	if err != nil {
		return startBlock - 1, errors.Wrapf(
			err,
			"error filtering calls from block %d to %d for chain %d",
			startBlock,
			toBlock,
			ob.Chain().ChainId,
		)
	}

	// parse and validate events
	events := ob.parseAndValidateCallEvents(eventIterator, gatewayAddr)

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
		if !ob.isEventProcessable(event.Sender, event.Receiver, event.Raw.TxHash, event.Payload) {
			continue
		}

		msg := ob.newCallInboundVote(event)

		ob.Logger().Inbound.Info().
			Msgf("ObserveGateway: Call inbound detected on chain %d tx %s block %d from %s value message %s",
				ob.Chain().
					ChainId, event.Raw.TxHash.Hex(), event.Raw.BlockNumber, event.Sender.Hex(), hex.EncodeToString(event.Payload))

		_, err = ob.PostVoteInbound(ctx, &msg, zetacore.PostVoteInboundExecutionGasLimit)
		if err != nil {
			// decrement the last scanned block so we have to re-scan from this block next time
			return lastScanned - 1, errors.Wrap(err, "error posting vote inbound")
		}
	}

	// successfully processed all events in [startBlock, toBlock]
	return toBlock, nil
}

// parseAndValidateCallEvents collects and sorts events by block number, tx index, and log index
func (ob *Observer) parseAndValidateCallEvents(
	iterator *gatewayevm.GatewayEVMCalledIterator,
	gatewayAddr ethcommon.Address,
) []*gatewayevm.GatewayEVMCalled {
	// collect and sort validEvents by block number, then tx index, then log index (ascending)
	validEvents := make([]*gatewayevm.GatewayEVMCalled, 0)
	for iterator.Next() {
		err := common.ValidateEvmTxLog(&iterator.Event.Raw, gatewayAddr, "", common.TopicsGatewayCall)
		if err == nil {
			validEvents = append(validEvents, iterator.Event)
			continue
		}
		ob.Logger().
			Inbound.Warn().
			Stringer(logs.FieldTx, iterator.Event.Raw.TxHash).
			Uint64(logs.FieldBlock, iterator.Event.Raw.BlockNumber).
			Msg("invalid Called event")
	}

	// order events by height, tx index and event index (ascending)
	// this ensures the first event is observed if there are multiple in the same tx
	sort.SliceStable(validEvents, func(i, j int) bool {
		if validEvents[i].Raw.BlockNumber == validEvents[j].Raw.BlockNumber {
			if validEvents[i].Raw.TxIndex == validEvents[j].Raw.TxIndex {
				return validEvents[i].Raw.Index < validEvents[j].Raw.Index
			}
			return validEvents[i].Raw.TxIndex < validEvents[j].Raw.TxIndex
		}
		return validEvents[i].Raw.BlockNumber < validEvents[j].Raw.BlockNumber
	})

	// filter events from same tx
	filtered := make([]*gatewayevm.GatewayEVMCalled, 0)
	guard := make(map[string]bool)
	for _, event := range validEvents {
		// guard against multiple events in the same tx
		if guard[event.Raw.TxHash.Hex()] {
			ob.Logger().Inbound.Warn().Stringer(logs.FieldTx, event.Raw.TxHash).Msg("multiple Called events in same tx")
			continue
		}
		guard[event.Raw.TxHash.Hex()] = true
		filtered = append(filtered, event)
	}

	return filtered
}

// newCallInboundVote creates a MsgVoteInbound message for a Call event
func (ob *Observer) newCallInboundVote(event *gatewayevm.GatewayEVMCalled) types.MsgVoteInbound {
	return *types.NewMsgVoteInbound(
		ob.ZetacoreClient().GetKeys().GetOperatorAddress().String(),
		event.Sender.Hex(),
		ob.Chain().ChainId,
		"",
		event.Receiver.Hex(),
		ob.ZetacoreClient().Chain().ChainId,
		sdkmath.ZeroUint(),
		hex.EncodeToString(event.Payload),
		event.Raw.TxHash.Hex(),
		event.Raw.BlockNumber,
		zetacore.PostVoteInboundCallOptionsGasLimit,
		coin.CoinType_NoAssetCall,
		"",
		uint64(event.Raw.Index),
		types.ProtocolContractVersion_V2,
		false, // currently not relevant since calls are not arbitrary
		types.InboundStatus_SUCCESS,
		types.ConfirmationMode_SAFE,
		types.WithEVMRevertOptions(event.RevertOptions),
	)
}

// ObserveGatewayDepositAndCall queries the gateway contract for deposit and call events
// returns the last block successfully scanned
func (ob *Observer) ObserveGatewayDepositAndCall(ctx context.Context, startBlock, toBlock uint64) (uint64, error) {
	gatewayAddr, gatewayContract, err := ob.GetGatewayContract()
	if err != nil {
		// lastScanned is startBlock - 1
		return startBlock - 1, errors.Wrap(err, "can't get gateway contract")
	}

	// get iterator for the events for the block range
	eventIterator, err := gatewayContract.FilterDepositedAndCalled(&bind.FilterOpts{
		Start:   startBlock,
		End:     &toBlock,
		Context: ctx,
	}, []ethcommon.Address{}, []ethcommon.Address{})
	if err != nil {
		return startBlock - 1, errors.Wrapf(
			err,
			"error filtering deposits from block %d to %d for chain %d",
			startBlock,
			toBlock,
			ob.Chain().ChainId,
		)
	}

	// parse and validate events
	events := ob.parseAndValidateDepositAndCallEvents(eventIterator, gatewayAddr)

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
		if !ob.isEventProcessable(event.Sender, event.Receiver, event.Raw.TxHash, event.Payload) {
			continue
		}

		msg := ob.newDepositAndCallInboundVote(event)

		ob.Logger().Inbound.Info().
			Msgf("ObserveGateway: DepositAndCall inbound detected on chain %d tx %s block %d from %s value %s message %s",
				ob.Chain().
					ChainId, event.Raw.TxHash.Hex(), event.Raw.BlockNumber, event.Sender.Hex(), event.Amount.String(), hex.EncodeToString(event.Payload))

		_, err = ob.PostVoteInbound(ctx, &msg, zetacore.PostVoteInboundExecutionGasLimit)
		if err != nil {
			// decrement the last scanned block so we have to re-scan from this block next time
			return lastScanned - 1, errors.Wrap(err, "error posting vote inbound")
		}
	}

	// successfully processed all events in [startBlock, toBlock]
	return toBlock, nil
}

// parseAndValidateDepositAndCallEvents collects and sorts events by block number, tx index, and log index
func (ob *Observer) parseAndValidateDepositAndCallEvents(
	iterator *gatewayevm.GatewayEVMDepositedAndCalledIterator,
	gatewayAddr ethcommon.Address,
) []*gatewayevm.GatewayEVMDepositedAndCalled {
	// collect and sort validEvents by block number, then tx index, then log index (ascending)
	validEvents := make([]*gatewayevm.GatewayEVMDepositedAndCalled, 0)
	for iterator.Next() {
		err := common.ValidateEvmTxLog(&iterator.Event.Raw, gatewayAddr, "", common.TopicsGatewayDepositAndCall)
		if err == nil {
			validEvents = append(validEvents, iterator.Event)
			continue
		}
		ob.Logger().
			Inbound.Warn().
			Stringer(logs.FieldTx, iterator.Event.Raw.TxHash).
			Uint64(logs.FieldBlock, iterator.Event.Raw.BlockNumber).
			Msg("invalid DepositedAndCalled event")
	}

	// order events by height, tx index and event index (ascending)
	// this ensures the first event is observed if there are multiple in the same tx
	sort.SliceStable(validEvents, func(i, j int) bool {
		if validEvents[i].Raw.BlockNumber == validEvents[j].Raw.BlockNumber {
			if validEvents[i].Raw.TxIndex == validEvents[j].Raw.TxIndex {
				return validEvents[i].Raw.Index < validEvents[j].Raw.Index
			}
			return validEvents[i].Raw.TxIndex < validEvents[j].Raw.TxIndex
		}
		return validEvents[i].Raw.BlockNumber < validEvents[j].Raw.BlockNumber
	})

	// filter events from same tx
	filtered := make([]*gatewayevm.GatewayEVMDepositedAndCalled, 0)
	guard := make(map[string]bool)
	for _, event := range validEvents {
		// guard against multiple events in the same tx
		if guard[event.Raw.TxHash.Hex()] {
			ob.Logger().
				Inbound.Warn().
				Stringer(logs.FieldTx, event.Raw.TxHash).
				Msg("multiple DepositedAndCalled events in same tx")
			continue
		}
		guard[event.Raw.TxHash.Hex()] = true
		filtered = append(filtered, event)
	}

	return filtered
}

// newDepositAndCallInboundVote creates a MsgVoteInbound message for a Deposit event
func (ob *Observer) newDepositAndCallInboundVote(event *gatewayevm.GatewayEVMDepositedAndCalled) types.MsgVoteInbound {
	// if event.Asset is zero, it's a native token
	coinType := coin.CoinType_ERC20
	if crypto.IsEmptyAddress(event.Asset) {
		coinType = coin.CoinType_Gas
	}

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
		coinType,
		event.Asset.Hex(),
		uint64(event.Raw.Index),
		types.ProtocolContractVersion_V2,
		false, // currently not relevant since calls are not arbitrary
		types.InboundStatus_SUCCESS,
		types.ConfirmationMode_SAFE,
		types.WithEVMRevertOptions(event.RevertOptions),
		types.WithCrossChainCall(true),
	)
}
