package e2etests

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"
)

// TestStressEtherWithdraw tests the stressing withdraw of ether with payload
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 2)

	// Parse withdrawal amount and number of withdrawals
	withdrawalAmount := utils.ParseBigInt(r, args[0])
	numWithdrawals, err := strconv.Atoi(args[1])
	require.NoError(r, err)
	require.GreaterOrEqual(r, numWithdrawals, 1)

	// Approve sufficient tokens for the Gateway contract
	approvalAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil) // 1e20
	tx, err := r.ETHZRC20.Approve(r.ZEVMAuth, r.GatewayZEVMAddr, approvalAmount)
	require.NoError(r, err)

	r.Logger.Print("Approving tokens for Gateway contract, tx hash: %s", tx.Hash().Hex())
	r.WaitForTxReceiptOnZEVM(tx)

	r.Logger.Print("Starting stress test with %d withdrawals", numWithdrawals)

	// Store durations for reporting
	withdrawDurations := []float64{}
	withdrawDurationsLock := sync.Mutex{}

	// Error group for concurrent withdrawals
	var eg errgroup.Group

	for i := 0; i < numWithdrawals; i++ {
		i := i // Capture loop variable

		// Generate a random payload for each withdrawal
		payload := randomPayload(r)

		// Perform the withdrawal
		eg.Go(func() error {
			startTime := time.Now()

			r.Logger.Print("Starting withdrawal #%d with payload: %x", i, payload)

			tx := r.ETHWithdrawAndCall(
				r.ZEVMAuth.From, // Sender address
				withdrawalAmount,
				[]byte(payload),
				gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
			)

			r.Logger.Print("Withdrawal #%d initiated, tx hash: %s", i, tx.Hash().Hex())

			// Wait for the cctx to be mined
			cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
			if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
				return fmt.Errorf(
					"Withdrawal #%d failed: status=%s, message=%s, cctx index=%s",
					i,
					cctx.CctxStatus.Status,
					cctx.CctxStatus.StatusMessage,
					cctx.Index,
				)
			}

			timeToComplete := time.Since(startTime)
			r.Logger.Print("Withdrawal #%d succeeded in %s", i, timeToComplete.String())

			// Record the duration
			withdrawDurationsLock.Lock()
			withdrawDurations = append(withdrawDurations, timeToComplete.Seconds())
			withdrawDurationsLock.Unlock()

			return nil
		})
	}

	// Wait for all withdrawals to complete
	err = eg.Wait()
	require.NoError(r, err)

	// Calculate and log latency statistics
	desc, descErr := stats.Describe(withdrawDurations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
	if descErr != nil {
		r.Logger.Print("Failed to calculate latency report: %v", descErr)
		return
	}

	r.Logger.Print("Latency Report:")
	r.Logger.Print("Min:  %.2f", desc.Min)
	r.Logger.Print("Max:  %.2f", desc.Max)
	r.Logger.Print("Mean: %.2f", desc.Mean)
	r.Logger.Print("Std:  %.2f", desc.Std)
	for _, p := range desc.DescriptionPercentiles {
		r.Logger.Print("P%.0f:  %.2f", p.Percentile, p.Value)
	}

	r.Logger.Print("All withdrawals completed successfully")
}
