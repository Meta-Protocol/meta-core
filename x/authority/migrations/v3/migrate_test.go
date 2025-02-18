package v3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zeta-chain/node/testutil/keeper"
	v3 "github.com/zeta-chain/node/x/authority/migrations/v3"
	"github.com/zeta-chain/node/x/authority/types"
)

func TestMigrateStore(t *testing.T) {
	t.Run("update authorization list", func(t *testing.T) {
		// Arrange
		k, ctx := keepertest.AuthorityKeeper(t)

		list := types.DefaultAuthorizationsList()
		list.RemoveAuthorization("/zetachain.zetacore.fungible.MsgUpdateZRC20Name")
		list.RemoveAuthorization("/zetachain.zetacore.crosschain.MsgRemoveInboundTracker")
		list.RemoveAuthorization("/zetachain.zetacore.observer.MsgUpdateOperationalChainParams")
		list.RemoveAuthorization("/zetachain.zetacore.observer.MsgUpdateChainParams")
		k.SetAuthorizationList(ctx, list)

		// Act
		err := v3.MigrateStore(ctx, *k)

		// Assert
		require.NoError(t, err)
		list, found := k.GetAuthorizationList(ctx)
		require.True(t, found)

		require.ElementsMatch(t, types.DefaultAuthorizationsList().Authorizations, list.Authorizations)
	})

	t.Run("set default authorization list if list is not found", func(t *testing.T) {
		// Arrange
		k, ctx := keepertest.AuthorityKeeper(t)

		// Act
		err := v3.MigrateStore(ctx, *k)

		// Assert
		require.NoError(t, err)
		list, found := k.GetAuthorizationList(ctx)
		require.True(t, found)
		require.Equal(t, types.DefaultAuthorizationsList(), list)
	})

	t.Run("return error list is invalid", func(t *testing.T) {
		// Arrange
		k, ctx := keepertest.AuthorityKeeper(t)

		k.SetAuthorizationList(ctx, types.AuthorizationList{Authorizations: []types.Authorization{
			{
				MsgUrl:           "ABC",
				AuthorizedPolicy: types.PolicyType_groupEmergency,
			},
			{
				MsgUrl:           "ABC",
				AuthorizedPolicy: types.PolicyType_groupEmergency,
			},
		}})

		// Act
		err := v3.MigrateStore(ctx, *k)

		// Assert
		require.Error(t, err)
	})
}
