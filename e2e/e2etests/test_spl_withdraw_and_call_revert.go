package e2etests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	solanacontract "github.com/zeta-chain/node/pkg/contracts/solana"
	"github.com/zeta-chain/node/testutil/sample"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"
)

// TestSPLWithdrawAndCall executes withdrawAndCall on zevm and calls connected program on solana
// execution is reverted in connected program on_call function
func TestSPLWithdrawAndCallRevert(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	withdrawAmount := utils.ParseBigInt(r, args[0])

	// get SPL ZRC20 balance before withdraw
	zrc20BalanceBefore, err := r.SPLZRC20.BalanceOf(&bind.CallOpts{}, r.EVMAddress())
	require.NoError(r, err)
	r.Logger.Info("runner balance of SPL before withdraw: %d", zrc20BalanceBefore)

	require.Equal(r, 1, zrc20BalanceBefore.Cmp(withdrawAmount), "Insufficient balance for withdrawal")

	// parse withdraw amount (in lamports), approve amount is 1 SOL
	approvedAmount := new(big.Int).SetUint64(solana.LAMPORTS_PER_SOL)
	require.Equal(
		r,
		-1,
		withdrawAmount.Cmp(approvedAmount),
		"Withdrawal amount must be less than the %v",
		approvedAmount,
	)

	// load deployer private key
	privkey := r.GetSolanaPrivKey()

	// get receiver ata balance before withdraw
	receiverAta := r.ResolveSolanaATA(privkey, privkey.PublicKey(), r.SPLAddr)
	receiverBalanceBefore, err := r.SolanaClient.GetTokenAccountBalance(r.Ctx, receiverAta, rpc.CommitmentConfirmed)
	require.NoError(r, err)
	r.Logger.Info("receiver balance of SPL before withdraw: %s", receiverBalanceBefore.Value.Amount)

	connected := solana.MustPublicKeyFromBase58(runner.ConnectedSPLProgramID.String())
	connectedPda, err := solanacontract.ComputeConnectedSPLPdaAddress(connected)
	require.NoError(r, err)

	connectedPdaAta := r.ResolveSolanaATA(privkey, connectedPda, r.SPLAddr)
	connectedPdaBalanceBefore, err := r.SolanaClient.GetTokenAccountBalance(
		r.Ctx,
		connectedPdaAta,
		rpc.CommitmentConfirmed,
	)
	require.NoError(r, err)
	r.Logger.Info("connected pda balance of SPL before withdraw: %s", connectedPdaBalanceBefore.Value.Amount)

	// use a random address to get the revert amount
	revertAddress := sample.EthAddress()
	balance, err := r.SPLZRC20.BalanceOf(&bind.CallOpts{}, revertAddress)
	require.NoError(r, err)
	require.EqualValues(r, int64(0), balance.Int64())

	// withdraw
	tx := r.WithdrawAndCallSPLZRC20(runner.ConnectedSPLProgramID, withdrawAmount, approvedAmount, []byte("revert"), gatewayzevm.RevertOptions{
		RevertAddress:    revertAddress,
		OnRevertGasLimit: big.NewInt(0),
	})

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	utils.RequireCCTXStatus(r, cctx, crosschaintypes.CctxStatus_Reverted)

	// get SPL ZRC20 balance after withdraw
	zrc20BalanceAfter, err := r.SPLZRC20.BalanceOf(&bind.CallOpts{}, r.EVMAddress())
	require.NoError(r, err)
	r.Logger.Info("runner balance of SPL after withdraw: %d", zrc20BalanceAfter)

	balanceAfter, err := r.SPLZRC20.BalanceOf(&bind.CallOpts{}, r.EVMAddress())
	require.NoError(r, err)
	r.Logger.Info("runner balance of SOL after withdraw: %d", balanceAfter)

	// check the balance of revert address is equal to withdraw amount
	balance, err = r.SPLZRC20.BalanceOf(&bind.CallOpts{}, revertAddress)
	require.NoError(r, err)

	require.True(r, balance.Cmp(withdrawAmount) == 0)
}
