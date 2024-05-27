package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	authoritytypes "github.com/zeta-chain/zetacore/x/authority/types"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

// DisableCCTXFlags disables the IsInboundEnabled and IsOutboundEnabled flags. These flags control the creation of inbounds and outbounds.
// The flags are disabled by the policy account with the groupEmergency policy type.
func (k msgServer) DisableCCTXFlags(
	goCtx context.Context,
	msg *types.MsgDisableCCTXFlags,
) (*types.MsgDisableCCTXFlagsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check permission
	if !k.GetAuthorityKeeper().IsAuthorized(ctx, msg.Creator, authoritytypes.PolicyType_groupEmergency) {
		return &types.MsgDisableCCTXFlagsResponse{}, authoritytypes.ErrUnauthorized.Wrap(
			"DisableCCTXFlags can only be executed by the correct policy account",
		)
	}

	// check if the value exists,
	// if not, set the default value for the Inbound and Outbound flags only
	flags, isFound := k.GetCrosschainFlags(ctx)
	if !isFound {
		flags = *types.DefaultCrosschainFlags()
		flags.GasPriceIncreaseFlags = nil
	}

	if msg.DisableInbound {
		flags.IsInboundEnabled = false
	}
	if msg.DisableOutbound {
		flags.IsOutboundEnabled = false
	}

	k.SetCrosschainFlags(ctx, flags)

	err := ctx.EventManager().EmitTypedEvents(&types.EventCCTXFlagsDisabled{
		MsgTypeUrl:        sdk.MsgTypeURL(&types.MsgDisableCCTXFlags{}),
		IsInboundEnabled:  flags.IsInboundEnabled,
		IsOutboundEnabled: flags.IsOutboundEnabled,
	})

	if err != nil {
		ctx.Logger().Error("Error emitting event EventCCTXFlagsDisabled :", err)
	}

	return &types.MsgDisableCCTXFlagsResponse{}, nil
}
