package e2etests

import (
	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/pkg/coin"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestSuiDepositAndCallRevert(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	amount := utils.ParseBigInt(r, args[0])

	signer, err := r.Account.SuiSigner()
	require.NoError(r, err, "get deployer signer")
	balanceBefore := r.SuiGetSUIBalance(signer.Address())

	// make the deposit transaction
	resp := r.SuiDepositAndCallSUI(r.TestDAppV2ZEVMAddr, math.NewUintFromBigInt(amount), []byte("revert"))

	r.Logger.Info("Sui deposit and call tx: %s", resp.Digest)

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, resp.Digest, r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "deposit")
	require.EqualValues(r, crosschaintypes.CctxStatus_Reverted, cctx.CctxStatus.Status)
	require.EqualValues(r, coin.CoinType_Gas, cctx.InboundParams.CoinType)
	require.True(r, cctx.InboundParams.IsCrossChainCall)

	// check the balance after the failed deposit is higher than balance before - amount
	// reason it's not equal is because of the gas fee for revert
	balanceAfter := r.SuiGetSUIBalance(signer.Address())
	require.Greater(r, balanceAfter, balanceBefore-amount.Uint64())
}
