package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// CCTXGatewayZEVM is implementation of CCTXGateway interface for ZEVM
type CCTXGatewayZEVM struct {
	crosschainKeeper Keeper
}

// NewCCTXGatewayZEVM returns new instance of CCTXGatewayZEVM
func NewCCTXGatewayZEVM(crosschainKeeper Keeper) CCTXGatewayZEVM {
	return CCTXGatewayZEVM{
		crosschainKeeper: crosschainKeeper,
	}
}

/*
InitiateOutbound handles evm deposit and call ValidateOutbound.
TODO (https://github.com/zeta-chain/node/issues/2278): move remaining of this comment to ValidateOutbound once it's added.

  - If the deposit is successful, the CCTX status is changed to OutboundMined.

  - If the deposit returns an internal error i.e if HandleEVMDeposit() returns an error, but isContractReverted is false, the CCTX status is changed to Aborted.

  - If the deposit is reverted, the function tries to create a revert cctx with status PendingRevert.

  - If the creation of revert tx also fails it changes the status to Aborted.

Note : Aborted CCTXs are not refunded in this function. The refund is done using a separate refunding mechanism.
We do not return an error from this function , as all changes need to be persisted to the state.
Instead we use a temporary context to make changes and then commit the context on for the happy path ,i.e cctx is set to OutboundMined.
New CCTX status after preprocessing is returned.
*/
func (c CCTXGatewayZEVM) InitiateOutbound(ctx sdk.Context, cctx *types.CrossChainTx) (newCCTXStatus types.CctxStatus) {
	tmpCtx, commit := ctx.CacheContext()
	isContractReverted, err := c.crosschainKeeper.HandleEVMDeposit(tmpCtx, cctx)

	// TODO (https://github.com/zeta-chain/node/issues/2278): further processing will be in validateOutbound(...), for now keeping it here
	if err != nil && !isContractReverted {
		// exceptional case; internal error; should abort CCTX
		cctx.SetAbort(err.Error())
		return types.CctxStatus_Aborted
	} else if err != nil && isContractReverted {
		// contract call reverted; should refund via a revert tx
		revertMessage := err.Error()
		senderChain := c.crosschainKeeper.zetaObserverKeeper.GetSupportedChainFromChainID(ctx, cctx.InboundParams.SenderChainId)
		if senderChain == nil {
			cctx.SetAbort(fmt.Sprintf("invalid sender chain id %d", cctx.InboundParams.SenderChainId))
			return types.CctxStatus_Aborted
		}
		gasLimit, err := c.crosschainKeeper.GetRevertGasLimit(ctx, *cctx)
		if err != nil {
			cctx.SetAbort(fmt.Sprintf("revert gas limit error: %s", err.Error()))
			return types.CctxStatus_Aborted
		}
		if gasLimit == 0 {
			// use same gas limit of outbound as a fallback -- should not be required
			gasLimit = cctx.GetCurrentOutboundParam().GasLimit
		}

		err = cctx.AddRevertOutbound(gasLimit)
		if err != nil {
			cctx.SetAbort(fmt.Sprintf("revert outbound error: %s", err.Error()))
			return types.CctxStatus_Aborted
		}
		// we create a new cached context, and we don't commit the previous one with EVM deposit
		tmpCtxRevert, commitRevert := ctx.CacheContext()
		err = func() error {
			err := c.crosschainKeeper.PayGasAndUpdateCctx(
				tmpCtxRevert,
				senderChain.ChainId,
				cctx,
				cctx.InboundParams.Amount,
				false,
			)
			if err != nil {
				return err
			}
			// Update nonce using senderchain id as this is a revert tx and would go back to the original sender
			return c.crosschainKeeper.UpdateNonce(tmpCtxRevert, senderChain.ChainId, cctx)
		}()
		if err != nil {
			cctx.SetAbort(fmt.Sprintf("deposit revert message: %s err : %s", revertMessage, err.Error()))
			return types.CctxStatus_Aborted
		}
		commitRevert()
		cctx.SetPendingRevert(revertMessage)
		return types.CctxStatus_PendingRevert
	}
	// successful HandleEVMDeposit;
	commit()
	cctx.SetOutboundMined("Remote omnichain contract call completed")
	return types.CctxStatus_OutboundMined
}
