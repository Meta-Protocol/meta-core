package e2etests

import (
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/txserver"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/pkg/chains"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestSolanaWhitelistSPL(r *runner.E2ERunner, _ []string) {
	// Deploy a new SPL
	r.Logger.Info("Deploying new SPL")

	// load deployer private key
	privkey, err := solana.PrivateKeyFromBase58(r.Account.SolanaPrivateKey.String())
	require.NoError(r, err)

	spl := r.DeploySPL(&privkey)

	// check that whitelist entry doesn't exist for this spl
	seed := [][]byte{[]byte("whitelist"), spl.PublicKey().Bytes()}
	whitelistEntryPDA, _, err := solana.FindProgramAddress(seed, r.GatewayProgram)
	require.NoError(r, err)

	whitelistEntryInfo, err := r.SolanaClient.GetAccountInfo(r.Ctx, whitelistEntryPDA)
	require.Error(r, err)
	require.Nil(r, whitelistEntryInfo)

	// whitelist sol zrc20
	r.Logger.Info("whitelisting spl on new network")
	res, err := r.ZetaTxServer.BroadcastTx(utils.AdminPolicyName, crosschaintypes.NewMsgWhitelistERC20(
		r.ZetaTxServer.MustGetAccountAddressFromName(utils.AdminPolicyName),
		spl.PublicKey().String(),
		chains.SolanaLocalnet.ChainId,
		"TESTSPL",
		"TESTSPL",
		6,
		100000,
	))
	require.NoError(r, err)

	event, ok := txserver.EventOfType[*crosschaintypes.EventERC20Whitelist](res.Events)
	require.True(r, ok, "no EventERC20Whitelist in %s", res.TxHash)
	erc20zrc20Addr := event.Zrc20Address
	whitelistCCTXIndex := event.WhitelistCctxIndex

	err = r.ZetaTxServer.InitializeLiquidityCaps(erc20zrc20Addr)
	require.NoError(r, err)

	// ensure CCTX created
	resCCTX, err := r.CctxClient.Cctx(r.Ctx, &crosschaintypes.QueryGetCctxRequest{Index: whitelistCCTXIndex})
	require.NoError(r, err)

	cctx := resCCTX.CrossChainTx
	r.Logger.CCTX(*cctx, "whitelist_cctx")

	// wait for the whitelist cctx to be mined
	r.WaitForMinedCCTXFromIndex(whitelistCCTXIndex)

	// check that whitelist entry exists for this spl
	whitelistEntryInfo, err = r.SolanaClient.GetAccountInfo(r.Ctx, whitelistEntryPDA)
	require.NoError(r, err)
	require.NotNil(r, whitelistEntryInfo)
}
