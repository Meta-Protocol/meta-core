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

// TxInfo holds metadata for each withdrawal transaction
type TxInfo struct {
    Index     int
    TxHash    string
    StartTime time.Time
}

// TestStressEtherWithdraw sends multiple ETH withdrawals sequentially on EVM,
// but waits for cross-chain confirmations concurrently.
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    // Parse the withdrawal amount and the number of withdrawals
    withdrawalAmount := utils.ParseBigInt(r, args[0])
    numWithdraws, err := strconv.Atoi(args[1])
    require.NoError(r, err, "Invalid number of withdrawals")
    require.GreaterOrEqual(r, numWithdraws, 1, "Number of withdrawals must be >= 1")

    // Capture the old balance for a final comparison
    oldBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
    require.NoError(r, err)

    // Approve Gateway to spend ETHZRC20 if necessary
    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    r.Logger.Print(fmt.Sprintf(
        "Starting sequential broadcast, concurrent wait. amount=%s, numWithdraws=%d",
        withdrawalAmount.String(),
        numWithdraws,
    ))

    // We'll collect each transaction's info (hash + startTime)
    txInfos := make([]TxInfo, 0, numWithdraws)

    // === 1) BROADCAST TXS SEQUENTIALLY (to avoid nonce conflicts) ===
    for i := 0; i < numWithdraws; i++ {
        start := time.Now()

        // Broadcast the withdrawal TX
        tx := r.ETHWithdraw(
            r.EVMAddress(),
            withdrawalAmount,
            gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
        )

        txHash := tx.Hash().Hex()
        r.Logger.Print(fmt.Sprintf("index=%d: Broadcast txHash=%s", i, txHash))

        // Wait for EVM receipt so the nonce is consumed before sending the next TX
        receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
        utils.RequireTxSuccessful(r, receipt)

        // Store the info for concurrent cross-chain waiting
        txInfos = append(txInfos, TxInfo{
            Index:     i,
            TxHash:    txHash,
            StartTime: start, // measure bridging latency from original broadcast
        })
    }

    // We'll collect each withdrawal's bridging duration (in seconds)
    var durations []float64
    var durationsLock sync.Mutex

    // === 2) WAIT FOR CCTXs CONCURRENTLY ===
    var eg errgroup.Group
    for _, info := range txInfos {
        info := info // capture loop variable
        eg.Go(func() error {
            cctx := utils.WaitCctxMinedByInboundHash(
                r.Ctx,
                info.TxHash,
                r.CctxClient,
                r.Logger,
                r.CctxTimeout,
            )
            if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
                return fmt.Errorf("index=%d: cctx failed. status=%s, msg=%s, cctxIndex=%s",
                    info.Index,
                    cctx.CctxStatus.Status,
                    cctx.CctxStatus.StatusMessage,
                    cctx.Index,
                )
            }

            elapsed := time.Since(info.StartTime).Seconds()
            r.Logger.Print(fmt.Sprintf("index=%d: CCTX success, bridgingDuration=%.2fs", info.Index, elapsed))

            durationsLock.Lock()
            durations = append(durations, elapsed)
            durationsLock.Unlock()

            return nil
        })
    }

    // Wait for all cross-chain confirmations
    require.NoError(r, eg.Wait(), "One or more cross-chain txs failed")

    // === 3) PERFORM LATENCY STATS ===
    desc, statsErr := stats.Describe(durations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
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

    // === 4) FINAL BALANCE CHECK ===
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
        "Expected new balance > old balance (minus gas).",
    )

    r.Logger.Print("All withdrawals completed successfully!")
}
