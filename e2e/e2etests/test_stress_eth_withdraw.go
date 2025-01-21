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

	// Import your generated binding for gateway
	gatewayzevm "github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"
)

// TestStressEtherWithdraw uses the Gateway contract's withdraw function to perform multiple withdrawals in parallel.
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    // parse withdrawal amount and number of withdrawals
    withdrawalAmount := utils.ParseBigInt(r, args[0])

    numWithdraws, err := strconv.Atoi(args[1])
    require.NoError(r, err)
    require.GreaterOrEqual(r, numWithdraws, 1)

    // 1) Approve the Gateway contract to spend your ZRC-20 tokens.
    //    Typically you want to allow enough to cover all your stress withdrawals.
    //    r.GatewayZEVMAddr should hold the gateway's contract address on ZEVM.
    approveAmount := big.NewInt(1e18)
    tx, err := r.ETHZRC20.Approve(r.ZEVMAuth, r.GatewayZEVMAddr, approveAmount)
    require.NoError(r, err, "approve transaction should succeed")

    // Wait for the approval tx to be mined on ZEVM
    r.WaitForTxReceiptOnZEVM(tx)

    r.Logger.Print("starting stress test of %d withdraws", numWithdraws)

    // 2) We'll collect latencies in a slice
    var (
        withdrawDurations     []float64
        withdrawDurationsLock sync.Mutex
    )

    // 3) Use an errgroup for concurrency
    var eg errgroup.Group

    // 4) Send the withdraws
    for i := 0; i < numWithdraws; i++ {
        // The index i must be "captured" in the loop properly
        i := i

        // Gateway expects:
        // withdraw(
        //   bytes memory receiver,
        //   uint256 amount,
        //   address zrc20,
        //   RevertOptions calldata revertOptions
        // )
        //
        // Convert EVM address to bytes for the "receiver"
        receiverBytes := r.EVMAddress().Bytes()

        // Optionally specify revert options if you want custom behavior;
        // here we pass empty
        revertOpts := gatewayzevm.RevertOptions{
            // e.g. RevertAddress: r.EVMAddress(), ...
        }

        // Create the withdraw transaction via Gateway
        tx, err := r.GatewayZEVM.Withdraw(
            r.ZEVMAuth,         // EVM signer opts
            receiverBytes,
            withdrawalAmount,
            r.ETHZRC20Addr,     // The ZRC-20 token for "ETH"
            revertOpts,
        )
        require.NoError(r, err, "gateway withdraw transaction should succeed")

        // Wait for the transaction to be confirmed on ZEVM
        receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
        utils.RequireTxSuccessful(r, receipt)

        r.Logger.Print("index %d: starting withdraw, tx hash: %s", i, tx.Hash().Hex())

        // Track the cross-chain transaction via errgroup
        eg.Go(func() error {
            startTime := time.Now()

            // 5) Wait until the cross-chain transaction is recognized as OutboundMined
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
            r.Logger.Print("index %d: withdraw cctx success in %s", i, timeToComplete.String())

            // 6) Record the latency
            withdrawDurationsLock.Lock()
            withdrawDurations = append(withdrawDurations, timeToComplete.Seconds())
            withdrawDurationsLock.Unlock()

            return nil
        })
    }

    // 7) Wait for all parallel withdrawals to finish
    err = eg.Wait()
    require.NoError(r, err, "some withdrawals failed")

    // 8) Collect statistics on the latencies
    desc, descErr := stats.Describe(withdrawDurations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
    if descErr != nil {
        r.Logger.Print("❌ failed to calculate latency report: %v", descErr)
    } else {
        r.Logger.Print("Latency report:")
        r.Logger.Print("min:  %.2f", desc.Min)
        r.Logger.Print("max:  %.2f", desc.Max)
        r.Logger.Print("mean: %.2f", desc.Mean)
        r.Logger.Print("std:  %.2f", desc.Std)
        for _, p := range desc.DescriptionPercentiles {
            r.Logger.Print("p%.0f:  %.2f", p.Percentile, p.Value)
        }
    }

    r.Logger.Print("all withdraws completed")
}
