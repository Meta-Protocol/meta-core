package e2etests

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
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
	baseWithdrawAmount := utils.ParseBigInt(r, args[0])
	numWithdrawalsSOL := utils.ParseInt(r, args[1])

	// Calculate the total approval amount (add a buffer for all transactions)
	bufferFactor := 1.1 // Approve 10% more
	totalApprovalAmount := new(big.Int).Mul(baseWithdrawAmount, big.NewInt(int64(numWithdrawalsSOL)))
	totalApprovalAmountWithBuffer := new(big.Int).Mul(totalApprovalAmount, big.NewInt(int64(bufferFactor*1e9)))
	totalApprovalAmountWithBuffer.Div(totalApprovalAmountWithBuffer, big.NewInt(1e9)) // Avoid floating-point inaccuracies

	// Load deployer private key
	privKey := r.GetSolanaPrivKey()

	r.Logger.Print("starting stress test of %d SOL withdrawals with total approval amount of %s (including buffer)", numWithdrawalsSOL, totalApprovalAmountWithBuffer.String())

	// Approve the total amount with buffer
	tx, err := r.SOLZRC20.Approve(r.ZEVMAuth, r.SOLZRC20Addr, totalApprovalAmountWithBuffer)
	require.NoError(r, err, "failed to approve total amount with buffer")
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "approve_sol")

	// Store transaction objects
	txObjects := make([]*types.Transaction, numWithdrawalsSOL)
	txObjectsLock := sync.Mutex{}

	// Step 1: Send all transactions concurrently
	sendGroup := errgroup.Group{}
	for i := 0; i < numWithdrawalsSOL; i++ {
		i := i // Capture loop variable for goroutine
		sendGroup.Go(func() error {
			// Increment the withdrawal amount by 1 lamport for each transaction
			withdrawAmount := new(big.Int).Add(baseWithdrawAmount, big.NewInt(int64(i)))

			// Fetch the latest nonce dynamically from the Ethereum client
			nonce, err := r.ZEVMClient.PendingNonceAt(r.Ctx, r.ZEVMAuth.From)
			if err != nil {
				return fmt.Errorf("index %d: failed to fetch nonce: %v", i, err)
			}

			// Create a new transaction authorizer with the dynamic nonce
			auth := *r.ZEVMAuth // Copy the original authorizer
			auth.Nonce = big.NewInt(int64(nonce))

			// Send the transaction
			tx, err := r.SOLZRC20.Withdraw(&auth, []byte(privKey.PublicKey().String()), withdrawAmount)
			if err != nil {
				return fmt.Errorf("index %d: failed to send SOL withdrawal transaction: %v", i, err)
			}

			r.Logger.Print("index %d: sent SOL withdraw, tx hash: %s, amount: %s, nonce: %d", i, tx.Hash().Hex(), withdrawAmount.String(), nonce)

			// Store the transaction object
			txObjectsLock.Lock()
			txObjects[i] = tx
			txObjectsLock.Unlock()

			return nil
		})
	}

	// Wait for all send operations to complete
	err = sendGroup.Wait()
	require.NoError(r, err, "error during transaction sends")

	// Step 2: Wait for all transactions concurrently
	waitGroup := errgroup.Group{}
	withdrawDurations := []float64{}
	withdrawDurationsLock := sync.Mutex{}

	for i, tx := range txObjects {
		i := i
		tx := tx // Capture loop variables for goroutine
		waitGroup.Go(func() error {
			startTime := time.Now()

			// Wait for receipt
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

			timeToComplete := time.Since(startTime)
			r.Logger.Print("index %d: withdraw SOL cctx success in %s", i, timeToComplete.String())

			// Store duration in a thread-safe way
			withdrawDurationsLock.Lock()
			withdrawDurations = append(withdrawDurations, timeToComplete.Seconds())
			withdrawDurationsLock.Unlock()

			return nil
		})
	}

	// Wait for all wait operations to complete
	err = waitGroup.Wait()
	require.NoError(r, err, "error during transaction waits")

	// Generate and print latency report
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

	r.Logger.Print("all SOL withdrawals completed")
}
