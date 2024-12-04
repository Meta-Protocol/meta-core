package observer

import (
	"context"
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/onrik/ethrpc"
	"github.com/zeta-chain/protocol-contracts/v2/pkg/gatewayevm.sol"

	"github.com/zeta-chain/node/zetaclient/zetacore"
)

var errEventNotFound = errors.New("no gateway event found in inbound tracker")

// ProcessInboundTrackerV2 processes inbound tracker events from the gateway
func (ob *Observer) ProcessInboundTrackerV2(
	ctx context.Context,
	gateway *gatewayevm.GatewayEVM,
	gatewayAddr ethcommon.Address,
	tx *ethrpc.Transaction,
	receipt *ethtypes.Receipt,
) error {
	// check confirmations
	if confirmed := ob.HasEnoughConfirmations(receipt, ob.LastBlock()); !confirmed {
		return fmt.Errorf(
			"inbound %s has not been confirmed yet: receipt block %d",
			tx.Hash,
			receipt.BlockNumber.Uint64(),
		)
	}

	for _, log := range receipt.Logs {
		if log == nil && log.Address != gatewayAddr {
			continue
		}

		// try parsing deposit
		eventDeposit, err := gateway.ParseDeposited(*log)
		if err == nil {
			// check if the event is processable
			if !ob.isEventProcessable(
				eventDeposit.Sender,
				eventDeposit.Receiver,
				eventDeposit.Raw.TxHash,
				eventDeposit.Payload,
			) {
				return fmt.Errorf("event from inbound tracker %s is not processable", tx.Hash)
			}
			msg := ob.newDepositInboundVote(eventDeposit)
			_, err = ob.PostVoteInbound(ctx, &msg, zetacore.PostVoteInboundExecutionGasLimit)
			return err
		}

		// try parsing deposit and call
		eventDepositAndCall, err := gateway.ParseDepositedAndCalled(*log)
		if err == nil {
			// check if the event is processable
			if !ob.isEventProcessable(
				eventDepositAndCall.Sender,
				eventDepositAndCall.Receiver,
				eventDepositAndCall.Raw.TxHash,
				eventDepositAndCall.Payload,
			) {
				return fmt.Errorf("event from inbound tracker %s is not processable", tx.Hash)
			}
			msg := ob.newDepositAndCallInboundVote(eventDepositAndCall)
			_, err = ob.PostVoteInbound(ctx, &msg, zetacore.PostVoteInboundExecutionGasLimit)
			return err
		}

		// try parsing call
		eventCall, err := gateway.ParseCalled(*log)
		if err == nil {
			// check if the event is processable
			if !ob.isEventProcessable(
				eventCall.Sender,
				eventCall.Receiver,
				eventCall.Raw.TxHash,
				eventCall.Payload,
			) {
				return fmt.Errorf("event from inbound tracker %s is not processable", tx.Hash)
			}
			msg := ob.newCallInboundVote(eventCall)
			_, err = ob.PostVoteInbound(ctx, &msg, zetacore.PostVoteInboundExecutionGasLimit)
			return err
		}
	}

	return errors.Wrapf(ErrEventNotFound, "inbound tracker %s", tx.Hash)
}
