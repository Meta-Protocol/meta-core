package keeper

import (
	"fmt"

	cosmoserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zeta-chain/zetacore/pkg/chains"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// InitiateOutbound initiates the outbound for the CCTX depending on the CCTX gateway.
// It does a conditional dispatch to correct CCTX gateway based on the receiver chain
// which handles the state changes and error handling.
func (k Keeper) InitiateOutbound(ctx sdk.Context, cctx *types.CrossChainTx) (types.CctxStatus, error) {
	receiverChainID := cctx.GetCurrentOutboundParam().ReceiverChainId
	chainInfo := chains.GetChainFromChainID(receiverChainID)
	if chainInfo == nil {
		return cctx.CctxStatus.Status, cosmoserrors.Wrap(
			types.ErrInitiatitingOutbound,
			fmt.Sprintf(
				"chain info not found for %d", receiverChainID,
			),
		)
	}

	cctxGateway, ok := k.cctxGateways[chainInfo.CctxGateway]
	if !ok {
		return cctx.CctxStatus.Status, cosmoserrors.Wrap(
			types.ErrInitiatitingOutbound,
			fmt.Sprintf(
				"CCTXGateway not defined for receiver chain %d", receiverChainID,
			),
		)
	}

	return cctxGateway.InitiateOutbound(ctx, cctx), nil
}
