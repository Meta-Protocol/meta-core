package e2etests

import (
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestSolanaDepositSPL2(r *runner.E2ERunner, _ []string) {
	// require.Len(r, args, 1)

	// load deployer private key
	privKey, err := solana.PrivateKeyFromBase58(r.Account.SolanaPrivateKey.String())
	require.NoError(r, err)

	wall, err := solana.WalletFromPrivateKeyBase58(r.SPLPrivateKey.String())
	require.NoError(r, err)
	sig := r.DepositSPL(&privKey, *wall, r.EVMAddress(), nil)

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, sig.String(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "solana_deposit")
	utils.RequireCCTXStatus(r, cctx, crosschaintypes.CctxStatus_OutboundMined)
}
