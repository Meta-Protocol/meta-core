package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
)

// TestKeeper_GetChainState tests get, and set chain state
func TestKeeper_GetChainState(t *testing.T) {
	k, ctx, _, _ := keepertest.LightclientKeeper(t)
	_, found := k.GetChainState(ctx, 42)
	require.False(t, found)

	k.SetChainState(ctx, sample.ChainState(42))
	_, found = k.GetChainState(ctx, 42)
	require.True(t, found)
}

func TestKeeper_GetAllChainStates(t *testing.T) {
	k, ctx, _, _ := keepertest.LightclientKeeper(t)
	c1 := sample.ChainState(42)
	c2 := sample.ChainState(43)
	c3 := sample.ChainState(44)

	k.SetChainState(ctx, c1)
	k.SetChainState(ctx, c2)
	k.SetChainState(ctx, c3)

	list := k.GetAllChainStates(ctx)
	require.Len(t, list, 3)
	require.EqualValues(t, c1, list[0])
	require.EqualValues(t, c2, list[1])
	require.EqualValues(t, c3, list[2])
}
