package e2etests

import (
	"math/big"

	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/pkg/constant"
	"github.com/zeta-chain/node/testutil/sample"
)

// TestSolanaDepositAndCallRevertWithDust tests Solana deposit and call that reverts with a dust amount in the revert outbound.
func TestSolanaDepositAndCallRevertWithDust(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0)

	// deposit the rent exempt amount which will result in a dust amount (after fee deduction) in the revert outbound
	depositAmount := big.NewInt(constant.SolanaWalletRentExempt)

	// ACT
	// execute the deposit and call transaction
	nonExistReceiver := sample.EthAddress()
	data := []byte("dust lamports should abort cctx")
	sig := r.SOLDepositAndCall(nil, nonExistReceiver, depositAmount, data)

	// ASSERT
	// Now we want to make sure cctx is aborted.
	utils.WaitCctxAbortedByInboundHash(r.Ctx, r, sig.String(), r.CctxClient)
}
