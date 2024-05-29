package authority_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/nullify"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/authority"
	"github.com/zeta-chain/zetacore/x/authority/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Policies:  sample.Policies(),
		ChainInfo: sample.ChainInfo(42),
	}

	// Init
	k, ctx := keepertest.AuthorityKeeper(t)
	authority.InitGenesis(ctx, *k, genesisState)

	// Check policy is set
	policies, found := k.GetPolicies(ctx)
	require.True(t, found)
	require.Equal(t, genesisState.Policies, policies)

	// Check chain info is set
	chainInfo, found := k.GetChainInfo(ctx)
	require.True(t, found)
	require.Equal(t, genesisState.ChainInfo, chainInfo)

	// Export
	got := authority.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// Compare genesis after init and export
	nullify.Fill(&genesisState)
	nullify.Fill(got)
	require.Equal(t, genesisState, *got)
}
