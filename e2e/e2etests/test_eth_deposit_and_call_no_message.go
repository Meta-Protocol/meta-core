package e2etests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/v2/pkg/gatewayevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestETHDepositAndCallNoMessage(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	amount := utils.ParseBigInt(r, args[0])

	oldBalance, err := r.ETHZRC20.BalanceOf(&bind.CallOpts{}, r.TestDAppV2ZEVMAddr)
	require.NoError(r, err)

	// perform the deposit and call to the TestDAppV2ZEVMAddr
	tx := r.ETHDepositAndCall(
		r.TestDAppV2ZEVMAddr,
		amount,
		[]byte{},
		gatewayevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
	)

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "deposit_and_call")
	require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)

	// check the payload was received on the contract
	messageIndex, err := r.TestDAppV2ZEVM.GetNoMessageIndex(&bind.CallOpts{}, r.EVMAddress())
	require.NoError(r, err)
	r.AssertTestDAppZEVMCalled(true, messageIndex, amount)

	// check the balance was updated
	newBalance, err := r.ETHZRC20.BalanceOf(&bind.CallOpts{}, r.TestDAppV2ZEVMAddr)
	require.NoError(r, err)
	require.Equal(r, new(big.Int).Add(oldBalance, amount), newBalance)
}
