package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	authoritytypes "github.com/zeta-chain/zetacore/x/authority/types"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

// EnableCCTXFlags enables the IsInboundEnabled and IsOutboundEnabled flags.These flags control the creation of inbounds and outbounds.
// The flags are enabled by the policy account with the groupOperational policy type.
func (k msgServer) EnableCCTXFlags(
	goCtx context.Context,
	msg *types.MsgEnableCCTXFlags,
) (*types.MsgEnableCCTXFlagsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check permission
	if !k.GetAuthorityKeeper().IsAuthorized(ctx, msg.Creator, authoritytypes.PolicyType_groupOperational) {
		return &types.MsgEnableCCTXFlagsResponse{}, authoritytypes.ErrUnauthorized.Wrap(
			"EnableCCTXFlags can only be executed by the correct policy account",
		)
	}

	// check if the value exists,
	// if not, set the default value for the Inbound and Outbound flags only
	flags, isFound := k.GetCrosschainFlags(ctx)
	if !isFound {
		flags = *types.DefaultCrosschainFlags()
		flags.GasPriceIncreaseFlags = nil
	}

	if msg.EnableInbound {
		flags.IsInboundEnabled = true
	}
	if msg.EnableOutbound {
		flags.IsOutboundEnabled = true
	}

	k.SetCrosschainFlags(ctx, flags)

	err := ctx.EventManager().EmitTypedEvents(&types.EventCCTXFlagsEnabled{
		MsgTypeUrl:        sdk.MsgTypeURL(&types.MsgEnableCCTXFlags{}),
		IsInboundEnabled:  flags.IsInboundEnabled,
		IsOutboundEnabled: flags.IsOutboundEnabled,
	})

	if err != nil {
		ctx.Logger().Error("Error emitting EventCCTXFlagsEnabled :", err)
	}

	return &types.MsgEnableCCTXFlagsResponse{}, nil
}
