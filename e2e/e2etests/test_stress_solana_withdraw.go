package e2etests

import (
	"fmt"
	"math/big"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

// TestStressSolanaWithdraw tests the stressing withdrawal of SOL
func TestStressSolanaWithdraw(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 2)

	withdrawSOLAmount := utils.ParseBigInt(r, args[0])
	numWithdrawalsSOL := utils.ParseInt(r, args[1])

	// Load deployer private key
	privKey := r.GetSolanaPrivKey()

	r.Logger.Print("Starting stress test of %d SOL withdrawals", numWithdrawalsSOL)

	// Approve sufficient amount for all withdrawals
	totalApproveAmount := new(big.Int).Mul(withdrawSOLAmount, big.NewInt(int64(numWithdrawalsSOL)))
	tx, err := r.SOLZRC20.Approve(r.ZEVMAuth, r.SOLZRC20Addr, totalApproveAmount)
	require.NoError(r, err)
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "approve_sol")

	// Store durations for latency analysis
	durationsChan := make(chan float64, numWithdrawalsSOL)
	defer close(durationsChan)

	// Create a group to manage goroutines
	var eg errgroup.Group

	for i := 0; i < numWithdrawalsSOL; i++ {
		i := i
		eg.Go(func() error {
			startTime := time.Now()

			// Execute the withdraw SOL transaction
			tx, err := r.SOLZRC20.Withdraw(r.ZEVMAuth, []byte(privKey.PublicKey().String()), withdrawSOLAmount)
			if err != nil {
				r.Logger.Print("Error in withdrawal %d: %v", i, err)
				return nil
			}

			receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
			utils.RequireTxSuccessful(r, receipt)

			r.Logger.Print("Index %d: Starting SOL withdraw, tx hash: %s", i, tx.Hash().Hex())

			// Wait for the cctx to be mined
			cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
			if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
				return fmt.Errorf(
					"Index %d: Withdraw cctx failed with status %s, message %s, cctx index %s",
					i,
					cctx.CctxStatus.Status,
					cctx.CctxStatus.StatusMessage,
					cctx.Index,
				)
			}

			timeToComplete := time.Since(startTime)
			r.Logger.Print("Index %d: Withdraw SOL cctx success in %s", i, timeToComplete.String())

			durationsChan <- timeToComplete.Seconds()
			return nil
		})
	}

	// Wait for all goroutines to complete
	err = eg.Wait()
	require.NoError(r, err)

	// Collect durations
	var withdrawDurations []float64
	for duration := range durationsChan {
		withdrawDurations = append(withdrawDurations, duration)
	}

	// Generate latency report
	desc, descErr := stats.Describe(withdrawDurations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
	if descErr != nil {
		r.Logger.Print("❌ Failed to calculate latency report: %v", descErr)
	}

	r.Logger.Print("Latency report:")
	r.Logger.Print("Min:  %.2f", desc.Min)
	r.Logger.Print("Max:  %.2f", desc.Max)
	r.Logger.Print("Mean: %.2f", desc.Mean)
	r.Logger.Print("Std:  %.2f", desc.Std)
	for _, p := range desc.DescriptionPercentiles {
		r.Logger.Print("P%.0f:  %.2f", p.Percentile, p.Value)
	}

	r.Logger.Print("All SOL withdrawals completed")
}
