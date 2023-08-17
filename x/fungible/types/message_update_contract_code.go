package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

const TypeMsgUpdateContractBytecode = "update_contract_bytecode"

var _ sdk.Msg = &MsgUpdateContractBytecode{}

func NewMsgUpdateContractBytecode(creator string, contractAddress ethcommon.Address, newBytecode []byte) *MsgUpdateContractBytecode {
	return &MsgUpdateContractBytecode{
		Creator:         creator,
		ContractAddress: contractAddress.Hex(),
		NewBytecode:     newBytecode,
	}
}

func (msg *MsgUpdateContractBytecode) Route() string {
	return RouterKey
}

func (msg *MsgUpdateContractBytecode) Type() string {
	return TypeMsgUpdateContractBytecode
}

func (msg *MsgUpdateContractBytecode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateContractBytecode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateContractBytecode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// check if the system contract address is valid
	if ethcommon.HexToAddress(msg.ContractAddress) == (ethcommon.Address{}) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", msg.ContractAddress)
	}

	return nil
}
