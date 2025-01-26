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

type TxInfo struct {
    Index     int
    TxHash    string
    StartTime time.Time
}

// TestStressEtherWithdraw sends multiple ETH withdrawals sequentially on EVM,
// but waits for cross-chain confirmations concurrently.
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    withdrawalAmount := utils.ParseBigInt(r, args[0])
    numWithdraws, err := strconv.Atoi(args[1])
    require.NoError(r, err, "Invalid number of withdrawals")
    require.GreaterOrEqual(r, numWithdraws, 1, "Number of withdrawals must be >= 1")

    require.NoError(r, err)

    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    txInfos := make([]TxInfo, 0, numWithdraws)

    for i := 0; i < numWithdraws; i++ {
        tx := r.ETHWithdraw(
            r.EVMAddress(),
            withdrawalAmount,
            gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
        )

        txHash := tx.Hash().Hex()
        r.Logger.Print(fmt.Sprintf("index=%d: start ETH withdraw, tx hash: %s", i, txHash))

        receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
        utils.RequireTxSuccessful(r, receipt)

        txInfos = append(txInfos, TxInfo{
            Index:     i,
            TxHash:    txHash,
        })
    }

    var durations []float64
    var durationsLock sync.Mutex

    var eg errgroup.Group
    for _, info := range txInfos {
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


						elapsed := time.Since(info.StartTime)
						r.Logger.Print(fmt.Sprintf("index=%d: withdraw SPL cctx success in %s", info.Index, elapsed.String()))

            durationsLock.Lock()
            durations = append(durations, elapsed.Seconds())
            durationsLock.Unlock()

            return nil
        })
    }

    require.NoError(r, eg.Wait(), "One or more cross-chain txs failed")

    desc, statsErr := stats.Describe(durations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
    if statsErr != nil {
        r.Logger.Print(fmt.Sprintf("Failed to compute latency stats: %v", statsErr))
    } else {
        r.Logger.Print("Latency Report:")
        r.Logger.Print(fmt.Sprintf(" min=%.2fs", desc.Min))
        r.Logger.Print(fmt.Sprintf(" max=%.2fs", desc.Max))
        r.Logger.Print(fmt.Sprintf(" mean=%.2fs", desc.Mean))
        r.Logger.Print(fmt.Sprintf(" std=%.2fs", desc.Std))
        for _, p := range desc.DescriptionPercentiles {
            r.Logger.Print(fmt.Sprintf(" p%.0f=%.2fs", p.Percentile, p.Value))
        }
    }

    require.NoError(r, err)

    r.Logger.Print("all ETH withdrawals completed")
}
