package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
)

func TestKeeper_SetChainInfo(t *testing.T) {
	k, ctx := keepertest.AuthorityKeeper(t)
	chainInfo := sample.ChainInfo(42)

	_, found := k.GetChainInfo(ctx)
	require.False(t, found)

	k.SetChainInfo(ctx, chainInfo)

	// Check policy is set
	got, found := k.GetChainInfo(ctx)
	require.True(t, found)
	require.Equal(t, chainInfo, got)

	// Can set policies again
	newChainInfo := sample.ChainInfo(84)
	require.NotEqual(t, chainInfo, newChainInfo)
	k.SetChainInfo(ctx, newChainInfo)
	got, found = k.GetChainInfo(ctx)
	require.True(t, found)
	require.Equal(t, newChainInfo, got)
}
