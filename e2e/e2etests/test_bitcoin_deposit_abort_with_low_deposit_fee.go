package e2etests

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	zetabitcoin "github.com/zeta-chain/node/zetaclient/chains/bitcoin/common"
)

func TestBitcoinDepositAndAbortWithLowDepositFee(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0)

	// Given small amount
	depositAmount := zetabitcoin.DefaultDepositorFee - float64(1)/btcutil.SatoshiPerBitcoin

	// ACT
	txHash := r.DepositBTCWithAmount(depositAmount, nil, false)

	// ASSERT
	// cctx status should be aborted
	cctx := utils.WaitCctxAbortedByInboundHash(r.Ctx, r, txHash.String(), r.CctxClient)
	r.Logger.CCTX(cctx, "deposit aborted")

	// check cctx details
	require.Equal(r, cctx.InboundParams.Amount.Uint64(), uint64(0))
	require.Equal(r, cctx.GetCurrentOutboundParam().Amount.Uint64(), uint64(0))

	// check cctx error message
	expectedError := fmt.Sprintf("deposited amount %v is less than depositor fee", depositAmount)
	require.Contains(r, cctx.CctxStatus.ErrorMessage, expectedError)
}
