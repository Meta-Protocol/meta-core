package e2etests

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"
)

// TestStressEtherWithdraw sends multiple ETH withdrawals sequentially.
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    // Parse the withdrawal amount (args[0]) and the number of withdrawals (args[1])
    withdrawalAmount := utils.ParseBigInt(r, args[0])
    numWithdraws, err := strconv.Atoi(args[1])
    require.NoError(r, err)
    require.GreaterOrEqual(r, numWithdraws, 1, "Number of withdrawals must be >= 1")

    // Capture the old balance for a final comparison
    oldBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
    require.NoError(r, err)

    // Approve the Gateway to spend ETHZRC20 on our behalf
    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    // Log the start
    r.Logger.Print(fmt.Sprintf(
        "Starting sequential ETH stress test with amount=%s, numWithdraws=%d",
        withdrawalAmount.String(),
        numWithdraws,
    ))

    // We'll collect each withdrawal's duration (in seconds)
    var withdrawDurations []float64
    var withdrawDurationsLock sync.Mutex

    // Sequentially send each withdrawal
    for i := 0; i < numWithdraws; i++ {
        startTime := time.Now()

        // Create and broadcast the withdrawal transaction
        tx := r.ETHWithdraw(
            r.EVMAddress(),
            withdrawalAmount,
            gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
        )

        r.Logger.Print(fmt.Sprintf("index=%d: Sent withdraw, txHash=%s", i, tx.Hash().Hex()))

        // Wait for the CCTX to be mined
        cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
        if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
            require.Failf(r, "Withdraw CCTX failed",
                "index=%d: Status=%s, Msg=%s, CCTXIndex=%s",
                i,
                cctx.CctxStatus.Status,
                cctx.CctxStatus.StatusMessage,
                cctx.Index,
            )
        }

        elapsed := time.Since(startTime).Seconds()
        r.Logger.Print(fmt.Sprintf("index=%d: CCTX success, duration=%.2fs", i, elapsed))

        withdrawDurationsLock.Lock()
        withdrawDurations = append(withdrawDurations, elapsed)
        withdrawDurationsLock.Unlock()
    }

    // Perform basic latency stats
    desc, statsErr := stats.Describe(withdrawDurations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
    if statsErr != nil {
        r.Logger.Print(fmt.Sprintf("Failed to compute latency stats: %v", statsErr))
    } else {
        r.Logger.Print("Latency Report:")
        r.Logger.Print(fmt.Sprintf("  min=%.2fs", desc.Min))
        r.Logger.Print(fmt.Sprintf("  max=%.2fs", desc.Max))
        r.Logger.Print(fmt.Sprintf(" mean=%.2fs", desc.Mean))
        r.Logger.Print(fmt.Sprintf("  std=%.2fs", desc.Std))
        for _, p := range desc.DescriptionPercentiles {
            r.Logger.Print(fmt.Sprintf(" p%.0f=%.2fs", p.Percentile, p.Value))
        }
    }

    // Final check: ensure the new balance is greater than the old (minus gas usage).
    newBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
    require.NoError(r, err)

    r.Logger.Print(fmt.Sprintf(
        "Old balance=%s, New balance=%s",
        oldBalance.String(),
        newBalance.String(),
    ))
    require.Greater(
        r,
        newBalance.Uint64(),
        oldBalance.Uint64(),
        "Expected new balance to be greater than old balance (minus gas).",
    )

    r.Logger.Print("All withdrawals completed successfully (sequential)!")
}
