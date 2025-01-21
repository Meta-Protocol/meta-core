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

	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestStressETHWithdrawAndCall(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    withdrawalAmount := utils.ParseBigInt(r, args[0])
    numWithdraws, err := strconv.Atoi(args[1])
    require.NoError(r, err)
    require.GreaterOrEqual(r, numWithdraws, 1)

    previousGasLimit := r.ZEVMAuth.GasLimit
    r.ZEVMAuth.GasLimit = 10_000_000
    defer func() {
        r.ZEVMAuth.GasLimit = previousGasLimit
    }()

    // Approve the ETH ZRC20 to Gateway, so we can withdraw
    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    r.Logger.Print("Starting stress test of %d withdraw-and-call operations", numWithdraws)

		var eg errgroup.Group
    var durationsLock sync.Mutex
    var withdrawDurations []float64

    for i := 0; i < numWithdraws; i++ {
        // It's nice to create a unique payload per call
        payload := randomPayload(r)
        i := i

        eg.Go(func() error {
            startTime := time.Now()

            tx := r.ETHWithdrawAndCall(
                r.TestDAppV2EVMAddr,
                withdrawalAmount,
                []byte(payload),
                gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
            )

            r.Logger.Print("index %d: starting withdraw-and-call, tx hash: %s", i, tx.Hash().Hex())

            cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
            if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
                return fmt.Errorf(
                    "index %d: cctx failed with status %s, message %s, cctx index %s",
                    i,
                    cctx.CctxStatus.Status,
                    cctx.CctxStatus.StatusMessage,
                    cctx.Index,
                )
            }

            timeToComplete := time.Since(startTime)
            r.Logger.Print("index %d: withdraw-and-call cctx success in %s", i, timeToComplete)

            r.AssertTestDAppEVMCalled(true, payload, withdrawalAmount)

            durationsLock.Lock()
            withdrawDurations = append(withdrawDurations, timeToComplete.Seconds())
            durationsLock.Unlock()

            return nil
        })
    }

    err = eg.Wait()
    require.NoError(r, err)

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

    r.Logger.Print("All withdraw-and-call operations completed successfully")
}
