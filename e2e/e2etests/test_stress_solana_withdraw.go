package e2etests

import (
	"fmt"
	"math/big"
	"sync"
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

	// load deployer private key
	privKey := r.GetSolanaPrivKey()

	r.Logger.Print("starting stress test of %d SOL withdrawals", numWithdrawalsSOL)

	tx, err := r.SOLZRC20.Approve(r.ZEVMAuth, r.SOLZRC20Addr, big.NewInt(1e18))
	require.NoError(r, err)
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "approve_sol")

	// create a wait group to wait for all the withdrawals to complete
	var eg errgroup.Group

	// store durations as float64 seconds like prometheus
	withdrawDurations := []float64{}
	withdrawDurationsLock := sync.Mutex{}

	// send the withdrawals SOL
	for i := 0; i < numWithdrawalsSOL; i++ {
    i := i // capture loop variable for goroutine

    eg.Go(func() error {
        // Start timing for this withdrawal
        startTime := time.Now()

        // Execute the withdraw SOL transaction concurrently
        tx, err := r.SOLZRC20.Withdraw(r.ZEVMAuth, []byte(privKey.PublicKey().String()), withdrawSOLAmount)
        if err != nil {
            return fmt.Errorf("index %d: failed to send SOL withdrawal transaction: %v", i, err)
        }

        r.Logger.Print("index %d: starting SOL withdraw, tx hash: %s", i, tx.Hash().Hex())

        // Wait for transaction receipt
        receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
        if receipt == nil {
            return fmt.Errorf("index %d: failed to get receipt for tx hash: %s", i, tx.Hash().Hex())
        }

        utils.RequireTxSuccessful(r, receipt, "withdraw_sol")

        // Wait for the cross-chain context (cctx) mining status
        cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
        if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
            return fmt.Errorf(
                "index %d: withdraw cctx failed with status %s, message %s, cctx index %s",
                i,
                cctx.CctxStatus.Status,
                cctx.CctxStatus.StatusMessage,
                cctx.Index,
            )
        }

        // Record the time it took to complete this transaction
        timeToComplete := time.Since(startTime)
        r.Logger.Print("index %d: withdraw SOL cctx success in %s", i, timeToComplete.String())

        // Store duration in a thread-safe way
        withdrawDurationsLock.Lock()
        withdrawDurations = append(withdrawDurations, timeToComplete.Seconds())
        withdrawDurationsLock.Unlock()

        return nil
    })
	}

	err = eg.Wait()

	desc, descErr := stats.Describe(withdrawDurations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
	if descErr != nil {
		r.Logger.Print("❌ failed to calculate latency report: %v", descErr)
	}

	r.Logger.Print("Latency report:")
	r.Logger.Print("min:  %.2f", desc.Min)
	r.Logger.Print("max:  %.2f", desc.Max)
	r.Logger.Print("mean: %.2f", desc.Mean)
	r.Logger.Print("std:  %.2f", desc.Std)
	for _, p := range desc.DescriptionPercentiles {
		r.Logger.Print("p%.0f:  %.2f", p.Percentile, p.Value)
	}

	require.NoError(r, err)
	r.Logger.Print("all SOL withdrawals completed")
}
