package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/fungible/types"
)

func TestMsgRemoveForeignCoin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRemoveForeignCoin
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRemoveForeignCoin{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: types.MsgRemoveForeignCoin{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
