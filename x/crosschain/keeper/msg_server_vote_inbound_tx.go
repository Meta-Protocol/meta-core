package keeper

import (
	"context"
	"fmt"

	cosmoserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/pkg/coin"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// FIXME: use more specific error types & codes

// VoteOnObservedInboundTx casts a vote on an inbound transaction observed on a connected chain. If this
// is the first vote, a new ballot is created. When a threshold of votes is
// reached, the ballot is finalized. When a ballot is finalized, a new CCTX is
// created.
//
// If the receiver chain is ZetaChain, `HandleEVMDeposit` is called. If the
// tokens being deposited are ZETA, `MintZetaToEVMAccount` is called and the
// tokens are minted to the receiver account on ZetaChain. If the tokens being
// deposited are gas tokens or ERC20 of a connected chain, ZRC20's `deposit`
// method is called and the tokens are deposited to the receiver account on
// ZetaChain. If the message is not empty, system contract's `depositAndCall`
// method is also called and an omnichain contract on ZetaChain is executed.
// Omnichain contract address and arguments are passed as part of the message.
// If everything is successful, the CCTX status is changed to `OutboundMined`.
//
// If the receiver chain is a connected chain, the `FinalizeInbound` method is
// called to prepare the CCTX to be processed as an outbound transaction. To
// cover the outbound transaction fee, the required amount of tokens submitted
// with the CCTX are swapped using a Uniswap V2 contract instance on ZetaChain
// for the ZRC20 of the gas token of the receiver chain. The ZRC20 tokens are
// then burned. The nonce is updated. If everything is successful, the CCTX
// status is changed to `PendingOutbound`.
//
// ```mermaid
// stateDiagram-v2
//
//	state evm_deposit_success <<choice>>
//	state finalize_inbound <<choice>>
//	state evm_deposit_error <<choice>>
//	PendingInbound --> evm_deposit_success: Receiver is ZetaChain
//	evm_deposit_success --> OutboundMined: EVM deposit success
//	evm_deposit_success --> evm_deposit_error: EVM deposit error
//	evm_deposit_error --> PendingRevert: Contract error
//	evm_deposit_error --> Aborted: Internal error, invalid chain, gas, nonce
//	PendingInbound --> finalize_inbound: Receiver is connected chain
//	finalize_inbound --> Aborted: Finalize inbound error
//	finalize_inbound --> PendingOutbound: Finalize inbound success
//
// ```
//
// Only observer validators are authorized to broadcast this message.
func (k msgServer) VoteOnObservedInboundTx(goCtx context.Context, msg *types.MsgVoteOnObservedInboundTx) (*types.MsgVoteOnObservedInboundTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	index := msg.Digest()

	// vote on inbound ballot
	// use a temporary context to not commit any ballot state change in case of error
	tmpCtx, commit := ctx.CacheContext()
	finalized, isNew, err := k.zetaObserverKeeper.VoteOnInboundBallot(
		tmpCtx,
		msg.SenderChainId,
		msg.ReceiverChain,
		msg.CoinType,
		msg.Creator,
		index,
		msg.InTxHash,
	)
	if err != nil {
		return nil, err
	}

	// If it is a new ballot, check if an inbound with the same hash, sender chain and event index has already been finalized
	// This may happen if the same inbound is observed twice where msg.Digest gives a different index
	// This check prevents double spending
	if isNew {
		if k.IsFinalizedInbound(tmpCtx, msg.InTxHash, msg.SenderChainId, msg.EventIndex) {
			return nil, cosmoserrors.Wrap(
				types.ErrObservedTxAlreadyFinalized,
				fmt.Sprintf("InTxHash:%s, SenderChainID:%d, EventIndex:%d", msg.InTxHash, msg.SenderChainId, msg.EventIndex),
			)
		}
	}
	commit()
	// If the ballot is not finalized return nil here to add vote to commit state
	if !finalized {
		return &types.MsgVoteOnObservedInboundTxResponse{}, nil
	}
	tss, tssFound := k.zetaObserverKeeper.GetTSS(ctx)
	if !tssFound {
		return nil, types.ErrCannotFindTSSKeys
	}
	// create a new CCTX from the inbound message.The status of the new CCTX is set to PendingInbound.
	cctx, err := types.NewCCTX(ctx, *msg, tss.TssPubkey)
	if err != nil {
		return nil, err
	}
	// Process the inbound CCTX, the process function manages the state commit and cctx status change.
	//	If the process fails, the changes to the evm state are rolled back.
	k.ProcessInbound(ctx, &cctx)
	// Save the inbound CCTX to the store. This is called irrespective of the status of the CCTX or the outcome of the process function.
	k.SaveInbound(ctx, &cctx, msg.EventIndex)
	return &types.MsgVoteOnObservedInboundTxResponse{}, nil
}

/* SaveInbound saves the inbound CCTX to the store.It does the following:
    - Emits an event for the finalized inbound CCTX.
	- Adds the inbound CCTX to the finalized inbound CCTX store.This is done to prevent double spending, using the same inbound tx hash and event index.
	- Updates the CCTX with the finalized height and finalization status.
	- Removes the inbound CCTX from the inbound transaction tracker store.This is only for inbounds created via InTx tracker suggestions
	- Sets the CCTX and nonce to the CCTX and inbound transaction hash to CCTX store.
*/

func (k Keeper) SaveInbound(ctx sdk.Context, cctx *types.CrossChainTx, eventIndex uint64) {
	if cctx.InboundTxParams.CoinType == coin.CoinType_Zeta && cctx.CctxStatus.Status != types.CctxStatus_OutboundMined {
		ctx.Logger().Info(fmt.Sprintf("SaveInbound Starting: cctx: %s", cctx.Index))
	}

	cctx.InboundTxParams.InboundTxFinalizedZetaHeight = uint64(ctx.BlockHeight())
	cctx.InboundTxParams.TxFinalizationStatus = types.TxFinalizationStatus_Executed
	k.SetCctxAndNonceToCctxAndInTxHashToCctx(ctx, *cctx)

	if cctx.InboundTxParams.CoinType == coin.CoinType_Zeta && cctx.CctxStatus.Status != types.CctxStatus_OutboundMined {
		ctx.Logger().Info(fmt.Sprintf("SaveInbound After: cctx: %s", cctx.Index))
	}

	EmitEventInboundFinalized(ctx, cctx)
	if cctx.InboundTxParams.CoinType == coin.CoinType_Zeta && cctx.CctxStatus.Status != types.CctxStatus_OutboundMined {
		ctx.Logger().Info(fmt.Sprintf("SaveInbound EmitEventInboundFinalized: cctx: %s", cctx.Index))
	}
	k.AddFinalizedInbound(ctx,
		cctx.GetInboundTxParams().InboundTxObservedHash,
		cctx.GetInboundTxParams().SenderChainId,
		eventIndex)
	if cctx.InboundTxParams.CoinType == coin.CoinType_Zeta && cctx.CctxStatus.Status != types.CctxStatus_OutboundMined {
		ctx.Logger().Info(fmt.Sprintf("SaveInbound AddFinalizedInbound: cctx: %s", cctx.Index))
	}
	// #nosec G701 always positive

	k.RemoveInTxTrackerIfExists(ctx, cctx.InboundTxParams.SenderChainId, cctx.InboundTxParams.InboundTxObservedHash)
	if cctx.InboundTxParams.CoinType == coin.CoinType_Zeta && cctx.CctxStatus.Status != types.CctxStatus_OutboundMined {
		ctx.Logger().Info(fmt.Sprintf("SaveInbound Finished: cctx: %s", cctx.Index))
	}
}
