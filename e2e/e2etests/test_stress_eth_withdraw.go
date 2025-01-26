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

func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    withdrawalAmount := utils.ParseBigInt(r, args[0])
    numWithdraws, err := strconv.Atoi(args[1])
    require.NoError(r, err)
    require.GreaterOrEqual(r, numWithdraws, 1, "Number of withdrawals must be >= 1")

    oldBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
    require.NoError(r, err)

    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    r.Logger.Print(fmt.Sprintf(
        "Starting ETH stress test with amount=%s, numWithdraws=%d",
        withdrawalAmount.String(),
        numWithdraws,
    ))

    var withdrawDurations []float64
    var withdrawDurationsLock sync.Mutex
    var eg errgroup.Group

    for i := 0; i < numWithdraws; i++ {
        i := i

        eg.Go(func() error {
            startTime := time.Now()

            tx := r.ETHWithdraw(
                r.EVMAddress(),
                withdrawalAmount,
                gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
            )

            r.Logger.Print(fmt.Sprintf("index=%d: Sent withdraw, txHash=%s", i, tx.Hash().Hex()))

            cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
            if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
                return fmt.Errorf(
                    "index=%d: withdraw CCTX failed. Status=%s, Msg=%s, CCTXIndex=%s",
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

            return nil
        })
    }

    err = eg.Wait()
    require.NoError(r, err, "One or more withdrawals failed")

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

    r.Logger.Print("All withdrawals completed successfully!")
}
