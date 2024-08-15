package e2etests

import (
	"math/big"

	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/zetacore/e2e/runner"
)

// TestOperationAddLiquidityERC20 is an operational test to add liquidity in erc20 token
func TestOperationAddLiquidityERC20(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 2)

	// #nosec G115 e2e - always in range
	liqZETA := big.NewInt(int64(parseInt(r, args[0])))
	// #nosec G115 e2e - always in range
	liqERC20 := big.NewInt(int64(parseInt(r, args[1])))

	// perform the add liquidity
	r.AddLiquidityERC20(liqZETA, liqERC20)
}
