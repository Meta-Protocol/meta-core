package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zeta-chain/zetacore/x/authority/types"
)

// SetChainInfo sets the chain info to the store
func (k Keeper) SetChainInfo(ctx sdk.Context, chainInfo types.ChainInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))
	b := k.cdc.MustMarshal(&chainInfo)
	store.Set([]byte{0}, b)
}

// GetChainInfo returns the policies from the store
func (k Keeper) GetChainInfo(ctx sdk.Context) (val types.ChainInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))
	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
