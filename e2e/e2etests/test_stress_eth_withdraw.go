package e2etests

import (
	"fmt"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

// TestStressEtherWithdraw tests the stressing withdrawal of ether
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    // parse withdrawal amount and number of withdrawals
    withdrawAmount := utils.ParseBigInt(r, args[0])
    numWithdrawals := utils.ParseInt(r, args[1])

    r.Logger.Print("starting stress test of %d withdrawals", numWithdrawals)

    // Create an errgroup to wait for all the withdrawals to finish
    var eg errgroup.Group

    for i := 0; i < numWithdrawals; i++ {
        i := i

        // Call your actual withdrawal function here.
        // Example: LegacyWithdrawEther might return a *types.Transaction.
        tx := r.LegacyWithdrawEther(withdrawAmount)
        // Convert transaction to a common.Hash
        txHash := tx.Hash()

        r.Logger.Print("index %d: starting withdrawal, tx hash: %s", i, txHash.Hex())

        // Launch goroutine to monitor the withdrawal
        eg.Go(func() error {
            return monitorEtherWithdrawal(r, txHash, i, time.Now())
        })
    }

    // Wait for all goroutines (withdrawals) to finish or fail
    require.NoError(r, eg.Wait())

    r.Logger.Print("all withdrawals completed")
}

// monitorEtherWithdrawal waits for the cross-chain transaction to become OutboundMined
func monitorEtherWithdrawal(
    r *runner.E2ERunner,
    txHash ethcommon.Hash,
    index int,
    startTime time.Time,
) error {
    // Wait for the cross‑chain transaction to be observed and outbound mined
    cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
    if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
        return fmt.Errorf(
            "index %d: withdrawal cctx failed with status %s, message %s, cctx index %s",
            index,
            cctx.CctxStatus.Status,
            cctx.CctxStatus.StatusMessage,
            cctx.Index,
        )
    }
    timeToComplete := time.Since(startTime)
    r.Logger.Print("index %d: withdrawal cctx success in %s", index, timeToComplete.String())

    return nil
}
