package e2etests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayevm.sol"

	"github.com/zeta-chain/node/e2e/contracts/testabort"
	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/testutil/sample"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestETHDepositRevertAndAbort(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	amount := utils.ParseBigInt(r, args[0])

	r.ApproveERC20OnEVM(r.GatewayEVMAddr)

	// deploy testabort contract
	testAbortAddr, _, testAbort, err := testabort.DeployTestAbort(r.ZEVMAuth, r.ZEVMClient)
	require.NoError(r, err)

	// perform the deposit
	tx := r.ETHDepositAndCall(
		r.TestDAppV2ZEVMAddr,
		amount,
		[]byte("revert"),
		gatewayevm.RevertOptions{
			RevertAddress:    r.TestDAppV2EVMAddr,
			CallOnRevert:     true,
			RevertMessage:    []byte("revert"),
			OnRevertGasLimit: big.NewInt(200000),
			AbortAddress:     testAbortAddr,
		},
	)

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "deposit")
	require.Equal(r, crosschaintypes.CctxStatus_Aborted, cctx.CctxStatus.Status)

	// check onAbort was called
	aborted, err := testAbort.IsAborted(&bind.CallOpts{})
	require.NoError(r, err)
	require.True(r, aborted)

	// check abort context was passed
	abortContext, err := testAbort.GetAbortedWithMessage(&bind.CallOpts{}, "revert")
	require.NoError(r, err)
	require.EqualValues(r, r.ETHZRC20Addr.Hex(), abortContext.Asset.Hex())

	// check abort contract received the tokens
	balance, err := r.ETHZRC20.BalanceOf(&bind.CallOpts{}, testAbortAddr)
	require.NoError(r, err)
	require.True(r, balance.Uint64() > 0)

	// Test 2: no contract ofr abort

	// check that funds are still received if onAbort is not called or fails
	abortAddressNoContract := sample.EthAddress()

	tx = r.ETHDepositAndCall(
		r.TestDAppV2ZEVMAddr,
		amount,
		[]byte("revert"),
		gatewayevm.RevertOptions{
			RevertAddress:    r.TestDAppV2EVMAddr,
			CallOnRevert:     true,
			RevertMessage:    []byte("revert"),
			OnRevertGasLimit: big.NewInt(200000),
			AbortAddress:     abortAddressNoContract,
		},
	)

	// wait for the cctx to be mined
	cctx = utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "deposit")
	require.Equal(r, crosschaintypes.CctxStatus_Aborted, cctx.CctxStatus.Status)

	// check abort contract received the tokens
	balance, err = r.ETHZRC20.BalanceOf(&bind.CallOpts{}, abortAddressNoContract)
	require.NoError(r, err)
	require.True(r, balance.Uint64() > 0)
}
