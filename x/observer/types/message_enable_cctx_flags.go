package types

import (
	cosmoserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgEnableCCTXFlags = "enable_crosschain_flags"
)

var _ sdk.Msg = &MsgEnableCCTXFlags{}

func NewMsgEnableCCTXFlags(creator string, enableInbound, enableOutbound bool) *MsgEnableCCTXFlags {
	return &MsgEnableCCTXFlags{
		Creator:        creator,
		EnableInbound:  enableInbound,
		EnableOutbound: enableOutbound,
	}
}

func (msg *MsgEnableCCTXFlags) Route() string {
	return RouterKey
}

func (msg *MsgEnableCCTXFlags) Type() string {
	return TypeMsgEnableCCTXFlags
}

func (msg *MsgEnableCCTXFlags) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnableCCTXFlags) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableCCTXFlags) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return cosmoserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !msg.EnableInbound && !msg.EnableOutbound {
		return cosmoserrors.Wrap(sdkerrors.ErrInvalidRequest, "at least one of EnableInbound or EnableOutbound must be true")
	}
	return nil
}
