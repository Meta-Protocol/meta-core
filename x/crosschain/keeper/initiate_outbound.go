package keeper

import (
	"fmt"

	cosmoserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zeta-chain/zetacore/pkg/chains"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// TODO: this is just a tmp solution, tbd if info can be passed to CCTX constructor somehow
// and not initialize CCTX using MsgVoteInbound but for example (InboundParams, OutboundParams)
// then PayGas can be decided based on GasPrice already presend in OutboundParams
// check if msg.Digest can be replaced to calculate index
type InitiateOutboundConfig struct {
	CCTX   *types.CrossChainTx
	PayGas bool
}

// InitiateOutbound initiates the outbound for the CCTX depending on the CCTX gateway.
// It does a conditional dispatch to correct CCTX gateway based on the receiver chain
// which handles the state changes and error handling.
func (k Keeper) InitiateOutbound(ctx sdk.Context, config InitiateOutboundConfig) (types.CctxStatus, error) {
	receiverChainID := config.CCTX.GetCurrentOutboundParam().ReceiverChainId
	chainInfo := chains.GetChainFromChainID(receiverChainID)
	if chainInfo == nil {
		return config.CCTX.CctxStatus.Status, cosmoserrors.Wrap(
			types.ErrInitiatitingOutbound,
			fmt.Sprintf(
				"chain info not found for %d", receiverChainID,
			),
		)
	}

	cctxGateway, found := ResolveCCTXGateway(chainInfo.CctxGateway, k)
	if !found {
		return config.CCTX.CctxStatus.Status, cosmoserrors.Wrap(
			types.ErrInitiatitingOutbound,
			fmt.Sprintf(
				"CCTXGateway not defined for receiver chain %d", receiverChainID,
			),
		)
	}

	return cctxGateway.InitiateOutbound(ctx, config)
}
