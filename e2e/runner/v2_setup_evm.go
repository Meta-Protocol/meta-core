package runner

import (
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/prototypes/evm/erc20custodynew.sol"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/prototypes/evm/gatewayevm.sol"

	"github.com/zeta-chain/zetacore/e2e/contracts/erc20"
	"github.com/zeta-chain/zetacore/e2e/utils"
	"github.com/zeta-chain/zetacore/pkg/constant"
)

var (
	zeroAddress = ethcommon.HexToAddress("0x0000000000000000000000000000000000000000")
)

// SetupEVMV2 setup contracts on EVM with v2 contracts
func (r *E2ERunner) SetupEVMV2() {
	r.Logger.Print("⚙️ setting up EVM v2 network")
	startTime := time.Now()
	defer func() {
		r.Logger.Info("EVM v2 setup took %s\n", time.Since(startTime))
	}()

	r.Logger.InfoLoud("Deploy Gateway ERC20Custody ERC20\n")

	// donate to the TSS address to avoid account errors because deploying gas token ZRC20 will automatically mint
	// gas token on ZetaChain to initialize the pool
	txDonation, err := r.SendEther(r.TSSAddress, big.NewInt(101000000000000000), []byte(constant.DonationMessage))
	require.NoError(r, err)

	r.Logger.Info("Deploying Gateway EVM")
	gatewayEVMAddr, txGateway, gatewayEVM, err := gatewayevm.DeployGatewayEVM(r.EVMAuth, r.EVMClient)
	require.NoError(r, err)

	r.GatewayEVMAddr = gatewayEVMAddr
	r.GatewayEVM = gatewayEVM
	r.Logger.Info("Gateway EVM contract address: %s, tx hash: %s", gatewayEVMAddr.Hex(), txGateway.Hash().Hex())

	r.Logger.Info("Deploying ERC20Custody contract")
	erc20CustodyNewAddr, txCustody, erc20CustodyNew, err := erc20custodynew.DeployERC20CustodyNew(
		r.EVMAuth,
		r.EVMClient,
		r.GatewayEVMAddr,
		r.TSSAddress,
	)
	require.NoError(r, err)

	r.ERC20CustodyAddr = erc20CustodyNewAddr
	r.ERC20CustodyNew = erc20CustodyNew
	r.Logger.Info(
		"ERC20CustodyNew contract address: %s, tx hash: %s",
		erc20CustodyNewAddr.Hex(),
		txCustody.Hash().Hex(),
	)

	r.Logger.Info("Deploying ERC20 contract")
	erc20Addr, txERC20, erc20, err := erc20.DeployERC20(r.EVMAuth, r.EVMClient, "TESTERC20", "TESTERC20", 6)
	require.NoError(r, err)

	r.ERC20 = erc20
	r.ERC20Addr = erc20Addr
	r.Logger.Info("ERC20 contract address: %s, tx hash: %s", erc20Addr.Hex(), txERC20.Hash().Hex())

	ensureTxReceipt := func(tx *ethtypes.Transaction, failMessage string) {
		receipt := utils.MustWaitForTxReceipt(r.Ctx, r.EVMClient, tx, r.Logger, r.ReceiptTimeout)
		r.requireTxSuccessful(receipt, failMessage)
	}

	// check contract deployment receipt
	ensureTxReceipt(txDonation, "EVM donation tx failed")
	ensureTxReceipt(txGateway, "GatewayEVM deployment failed")
	ensureTxReceipt(txCustody, "ERC20CustodyNew deployment failed")
	ensureTxReceipt(txERC20, "ERC20 deployment failed")

	// initialize contracts
	r.Logger.Info("Initialize Gateway EVM")
	txCustody, err = r.GatewayEVM.Initialize(r.EVMAuth, r.TSSAddress, zeroAddress)
	require.NoError(r, err)

	ensureTxReceipt(txCustody, "ERC20 update TSS address failed")

	r.Logger.Info("TSS set receipt tx hash: %s", txCustody.Hash().Hex())
}
