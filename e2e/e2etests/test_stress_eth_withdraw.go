package e2etests

import (
	"fmt"
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"
)

// TestStressEtherWithdraw tests the stressing withdraw of ether
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 2)

	// parse withdraw amount and number of withdraws
	withdrawAmount := utils.ParseBigInt(r, args[0])
	numWithdrawals := utils.ParseInt(r, args[1])

	r.Logger.Print("starting stress test of %d withdrawals", numWithdrawals)

	// create a wait group to wait for all the withdrawals to complete
	var eg errgroup.Group

	// approve tokens for the Gateway contract
	r.ApproveETHZRC20(r.GatewayZEVMAddr)

	// send the withdrawals
	for i := 0; i < numWithdrawals; i++ {
		i := i
		oldBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
		require.NoError(r, err)

		tx := r.ETHWithdraw(r.EVMAddress(), withdrawAmount, gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)})
		r.Logger.Print("index %d: starting withdraw, tx hash: %s", i, tx.Hash().Hex())

		eg.Go(func() error { return monitorEtherWithdrawal(r, tx.Hash(), i, oldBalance, time.Now()) })
	}

	require.NoError(r, eg.Wait())

	r.Logger.Print("all withdrawals completed")
}

// monitorEtherWithdrawal monitors the withdrawal of ether, returns once the withdrawal is complete
func monitorEtherWithdrawal(r *runner.E2ERunner, hash ethcommon.Hash, index int, oldBalance *big.Int, startTime time.Time) error {
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, hash.Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
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

	newBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
	if err != nil {
		return fmt.Errorf("index %d: failed to fetch new balance: %v", index, err)
	}

	if newBalance.Uint64() <= oldBalance.Uint64() {
		return fmt.Errorf("index %d: new balance is not greater than old balance", index)
	}

	return nil
}
