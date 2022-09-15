package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/zetacore/types"
)

func EmitEventSendFinalized(ctx sdk.Context, cctx *types.CrossChainTx) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.InboundFinalized,
			sdk.NewAttribute(types.CctxIndex, cctx.Index),
			sdk.NewAttribute(types.Sender, cctx.InBoundTxParams.Sender),
			sdk.NewAttribute(types.SenderChain, cctx.InBoundTxParams.SenderChain),
			sdk.NewAttribute(types.InTxHash, cctx.InBoundTxParams.InBoundTxObservedHash),
			sdk.NewAttribute(types.InBlockHeight, fmt.Sprintf("%d", cctx.InBoundTxParams.InBoundTxObservedHeight)),
			sdk.NewAttribute(types.Receiver, cctx.OutBoundTxParams.Receiver),
			sdk.NewAttribute(types.ReceiverChain, cctx.OutBoundTxParams.ReceiverChain),
			sdk.NewAttribute(types.ZetaBurnt, cctx.ZetaBurnt.String()),
			sdk.NewAttribute(types.ZetaMint, cctx.ZetaMint.String()),
			sdk.NewAttribute(types.RelayedMessage, cctx.RelayedMessage),
			sdk.NewAttribute(types.NewStatus, cctx.CctxStatus.Status.String()),
			sdk.NewAttribute(types.StatusMessage, cctx.CctxStatus.StatusMessage),
			sdk.NewAttribute(types.Identifiers, cctx.LogIdentifierForCCTX()),
		),
	)
}

func EmitEventCCTXCreated(ctx sdk.Context, cctx *types.CrossChainTx) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.InboundCreated,
			sdk.NewAttribute(types.CctxIndex, cctx.Index),
			sdk.NewAttribute(types.Sender, cctx.InBoundTxParams.Sender),
			sdk.NewAttribute(types.SenderChain, cctx.InBoundTxParams.SenderChain),
			sdk.NewAttribute(types.InTxHash, cctx.InBoundTxParams.InBoundTxObservedHash),
			sdk.NewAttribute(types.Receiver, cctx.OutBoundTxParams.Receiver),
			sdk.NewAttribute(types.ReceiverChain, cctx.OutBoundTxParams.ReceiverChain),
			sdk.NewAttribute(types.ZetaBurnt, cctx.ZetaBurnt.String()),
			sdk.NewAttribute(types.NewStatus, cctx.CctxStatus.String()),
			sdk.NewAttribute(types.Identifiers, cctx.LogIdentifierForCCTX()),
		),
	)
}

func EmitReceiveSuccess(ctx sdk.Context, msg *types.MsgVoteOnObservedOutboundTx, receive *types.Receive, oldStatus, newStatus, logIdentifiers string) {
	event := sdk.NewEvent(types.OutboundTxSuccessful,
		sdk.NewAttribute(types.CctxIndex, receive.SendHash),
		sdk.NewAttribute(types.OutTxHash, receive.OutTxHash),
		sdk.NewAttribute(types.ZetaMint, msg.MMint.String()),
		sdk.NewAttribute(types.OutBoundChain, msg.OutTxChain),
		sdk.NewAttribute(types.OldStatus, oldStatus),
		sdk.NewAttribute(types.NewStatus, newStatus),
		sdk.NewAttribute(types.Identifiers, logIdentifiers),
	)
	ctx.EventManager().EmitEvent(event)
}

func EmitReceiveFailure(ctx sdk.Context, msg *types.MsgVoteOnObservedOutboundTx, receive *types.Receive, oldStatus, newStatus, logIdentifiers string) {
	event := sdk.NewEvent(types.OutboundTxFailed,
		sdk.NewAttribute(types.CctxIndex, receive.SendHash),
		sdk.NewAttribute(types.OutTxHash, receive.OutTxHash),
		sdk.NewAttribute(types.ZetaMint, msg.MMint.String()),
		sdk.NewAttribute(types.OutBoundChain, msg.OutTxChain),
		sdk.NewAttribute(types.OldStatus, oldStatus),
		sdk.NewAttribute(types.NewStatus, newStatus),
		sdk.NewAttribute(types.Identifiers, logIdentifiers),
	)
	ctx.EventManager().EmitEvent(event)
}
