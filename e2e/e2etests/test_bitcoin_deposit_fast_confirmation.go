package e2etests

import (
	"math/big"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/pkg/constant"
	mathpkg "github.com/zeta-chain/node/pkg/math"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	observertypes "github.com/zeta-chain/node/x/observer/types"
)

// TestBitcoinDepositFastConfirmation tests the fast confirmation of Bitcoin deposits
func TestBitcoinDepositFastConfirmation(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0)

	// ARRANGE
	// enable inbound fast confirmation by updating the chain params
	chainID := r.GetBitcoinChainID()
	reqQuery := &observertypes.QueryGetChainParamsForChainRequest{ChainId: chainID}
	resOldChainParams, err := r.ObserverClient.GetChainParamsForChain(r.Ctx, reqQuery)
	require.NoError(r, err)

	// define new confirmation params
	chainParams := *resOldChainParams.ChainParams
	chainParams.ConfirmationParams = &observertypes.ConfirmationParams{
		SafeInboundCount:  5, // approx 30 seconds, much longer than Fast confirmation time (6 second)
		FastInboundCount:  1,
		SafeOutboundCount: 1,
		FastOutboundCount: 1,
	}
	err = r.ZetaTxServer.UpdateChainParams(&chainParams)
	require.NoError(r, err, "failed to enable inbound fast confirmation")

	// it takes 1 Zeta block time for zetaclient to pick up the new chain params
	time.Sleep(constant.ZetaBlockTime)
	r.Logger.Info("enabled inbound fast confirmation")

	// query current BTC ZRC20 supply
	supply, err := r.BTCZRC20.TotalSupply(&bind.CallOpts{})
	supplyUint := sdkmath.NewUintFromBigInt(supply)
	require.NoError(r, err)

	// set ZRC20 liquidity cap to 150% of the current supply
	// note: the percentage should not be too small as it may block other tests
	liquidityCap, _ := mathpkg.IncreaseUintByPercent(supplyUint, 50)
	require.True(r, liquidityCap.GT(sdkmath.ZeroUint()))
	res, err := r.ZetaTxServer.SetZRC20LiquidityCap(r.BTCZRC20Addr.Hex(), liquidityCap)
	require.NoError(r, err)
	r.Logger.Info("set liquidity cap to %s tx hash: %s", liquidityCap.String(), res.TxHash)

	// ACT-1
	// deposit with exactly fast amount cap, should be fast confirmed
	multiplier, enabled := constant.GetInboundFastConfirmationLiquidityMultiplier(chainID)
	require.True(r, enabled)
	fastAmountCap := constant.CalcInboundFastAmountCap(liquidityCap, multiplier)
	fastAmountCapFloat := float64(fastAmountCap.Uint64()) / btcutil.SatoshiPerBitcoin
	txHash := r.DepositBTCWithExactAmount(fastAmountCapFloat, nil)
	r.Logger.Info("deposited exactly fast amount %d cap tx hash: %s", fastAmountCap, txHash)

	// ASSERT-1
	// wait for the cctx to be FAST confirmed
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.String(), r.CctxClient, r.Logger, r.CctxTimeout)
	require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)
	require.Equal(r, crosschaintypes.ConfirmationMode_FAST, cctx.InboundParams.ConfirmationMode)

	r.Logger.Info("FAST confirmed deposit succeeded")

	// ACT-2
	// deposit with amount more than fast amount cap
	amountMoreThanCap := big.NewInt(0).Add(fastAmountCap, big.NewInt(1))
	amountMoreThanCapFloat := float64(amountMoreThanCap.Uint64()) / btcutil.SatoshiPerBitcoin
	txHash = r.DepositBTCWithExactAmount(amountMoreThanCapFloat, nil)
	r.Logger.Info("deposited more than fast amount cap %d tx hash: %s", amountMoreThanCap, txHash)

	// mine blocks at normal speed
	stop := r.MineBlocksIfLocalBitcoin()
	defer stop()

	// ASSERT-2
	// wait for the cctx to be SAFE confirmed
	cctx = utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.String(), r.CctxClient, r.Logger, r.CctxTimeout)
	require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)
	require.Equal(r, crosschaintypes.ConfirmationMode_SAFE, cctx.InboundParams.ConfirmationMode)

	r.Logger.Info("SAFE confirmed deposit succeeded")

	// TEARDOWN
	// restore old chain params
	err = r.ZetaTxServer.UpdateChainParams(resOldChainParams.ChainParams)
	require.NoError(r, err, "failed to restore chain params")

	// remove the liquidity cap
	_, err = r.ZetaTxServer.RemoveZRC20LiquidityCap(r.ETHZRC20Addr.Hex())
	require.NoError(r, err)
}
