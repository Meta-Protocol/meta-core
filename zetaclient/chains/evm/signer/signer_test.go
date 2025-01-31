package signer

import (
	"context"
	"math/big"
	"testing"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	observertypes "github.com/zeta-chain/node/x/observer/types"
	"github.com/zeta-chain/node/zetaclient/chains/evm/client"
	zctx "github.com/zeta-chain/node/zetaclient/context"
	"github.com/zeta-chain/node/zetaclient/keys"
	"github.com/zeta-chain/node/zetaclient/testutils/testlog"
	"github.com/zeta-chain/node/zetaclient/testutils/testrpc"

	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/testutil/sample"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/node/zetaclient/chains/base"
	"github.com/zeta-chain/node/zetaclient/config"
	"github.com/zeta-chain/node/zetaclient/testutils"
	"github.com/zeta-chain/node/zetaclient/testutils/mocks"
)

var (
	// Dummy addresses as they are just used as transaction data to be signed
	ConnectorAddress    = sample.EthAddress()
	ERC20CustodyAddress = sample.EthAddress()
)

type testSuite struct {
	*Signer
	tss       *mocks.TSS
	evmServer *testrpc.EVMServer
	client    *client.Client
}

func newTestSuite(t *testing.T) *testSuite {
	ctx := context.Background()

	chain := chains.BscMainnet

	evmServer := testrpc.NewEVMServer(t)

	evmServer.SetChainID(int(chain.ChainId))
	evmServer.MockSendTransaction()

	evmClient, err := client.NewFromEndpoint(ctx, evmServer.Endpoint)
	require.NoError(t, err)

	tss := mocks.NewTSS(t)

	logger := testlog.New(t)

	baseSigner := base.NewSigner(chain, tss, base.Logger{Std: logger.Logger, Compliance: logger.Logger})

	s, err := New(
		baseSigner,
		evmClient,
		ConnectorAddress,
		ERC20CustodyAddress,
		sample.EthAddress(),
	)
	require.NoError(t, err)

	return &testSuite{
		Signer:    s,
		tss:       tss,
		evmServer: evmServer,
		client:    evmClient,
	}
}

func (ts *testSuite) EvmSigner() ethtypes.Signer {
	return ts.client.Signer
}

func getCCTX(t *testing.T) *crosschaintypes.CrossChainTx {
	return testutils.LoadCctxByNonce(t, 56, 68270)
}

func getInvalidCCTX(t *testing.T) *crosschaintypes.CrossChainTx {
	cctx := getCCTX(t)
	// modify receiver chain id to make it invalid
	cctx.GetCurrentOutboundParam().ReceiverChainId = 13378337
	return cctx
}

// verifyTxSender is a helper function to verify the signature of a transaction
//
// signer.Sender() will ecrecover the public key of the transaction internally
// and will fail if the transaction is not valid or has been tampered with
func verifyTxSender(t *testing.T, tx *ethtypes.Transaction, expectedSender ethcommon.Address, signer ethtypes.Signer) {
	senderAddr, err := signer.Sender(tx)
	require.NoError(t, err)
	require.Equal(t, expectedSender.String(), senderAddr.String())
}

// verifyTxBodyBasics is a helper function to verify 'to', 'nonce' and 'amount' of a transaction
func verifyTxBodyBasics(
	t *testing.T,
	tx *ethtypes.Transaction,
	to ethcommon.Address,
	nonce uint64,
	amount *big.Int,
) {
	require.Equal(t, to, *tx.To())
	require.Equal(t, nonce, tx.Nonce())
	require.True(t, amount.Cmp(tx.Value()) == 0)
}

func TestSigner_SetGetConnectorAddress(t *testing.T) {
	evmSigner := newTestSuite(t)

	// Get and compare
	require.Equal(t, ConnectorAddress, evmSigner.GetZetaConnectorAddress())

	// Update and get again
	newConnector := sample.EthAddress()
	evmSigner.SetZetaConnectorAddress(newConnector)
	require.Equal(t, newConnector, evmSigner.GetZetaConnectorAddress())
}

func TestSigner_SetGetERC20CustodyAddress(t *testing.T) {
	evmSigner := newTestSuite(t)
	// Get and compare
	require.Equal(t, ERC20CustodyAddress, evmSigner.GetERC20CustodyAddress())

	// Update and get again
	newCustody := sample.EthAddress()
	evmSigner.SetERC20CustodyAddress(newCustody)
	require.Equal(t, newCustody, evmSigner.GetERC20CustodyAddress())
}

func TestSigner_TryProcessOutbound(t *testing.T) {
	ctx := makeCtx(t)

	// Setup evm signer
	evmSigner := newTestSuite(t)
	cctx := getCCTX(t)

	// Test with mock client that has keys
	client := mocks.NewZetacoreClient(t).
		WithKeys(&keys.Keys{}).
		WithZetaChain().
		WithPostVoteOutbound("", "")

	evmSigner.TryProcessOutbound(ctx, cctx, client, 123)

	// Check if cctx was signed and broadcasted
	list := evmSigner.GetReportedTxList()
	require.Len(t, *list, 1)
}

func TestSigner_BroadcastOutbound(t *testing.T) {
	ctx := makeCtx(t)

	// Setup evm signer
	evmSigner := newTestSuite(t)

	// Setup txData struct
	cctx := getCCTX(t)
	txData, skip, err := NewOutboundData(ctx, cctx, 123, zerolog.Logger{})
	require.NoError(t, err)
	require.False(t, skip)

	t.Run("BroadcastOutbound - should successfully broadcast", func(t *testing.T) {
		// Call SignERC20Withdraw
		tx, err := evmSigner.SignERC20Withdraw(ctx, txData)
		require.NoError(t, err)

		evmSigner.BroadcastOutbound(
			ctx,
			tx,
			cctx,
			zerolog.Logger{},
			sdktypes.AccAddress{},
			mocks.NewZetacoreClient(t),
			txData,
		)

		//Check if cctx was signed and broadcasted
		list := evmSigner.GetReportedTxList()
		require.Len(t, *list, 1)
	})
}

func TestSigner_SignerErrorMsg(t *testing.T) {
	cctx := getCCTX(t)

	msg := ErrorMsg(cctx)
	require.Contains(t, msg, "nonce 68270 chain 56")
}

func makeCtx(t *testing.T) context.Context {
	app := zctx.New(config.New(false), nil, zerolog.Nop())

	bscParams := mocks.MockChainParams(chains.BscMainnet.ChainId, 10)

	err := app.Update(
		[]chains.Chain{chains.BscMainnet, chains.ZetaChainMainnet},
		nil,
		map[int64]*observertypes.ChainParams{
			chains.BscMainnet.ChainId: &bscParams,
		},
		observertypes.CrosschainFlags{},
	)
	require.NoError(t, err, "unable to update app context")

	return zctx.WithAppContext(context.Background(), app)
}

func makeLogger(t *testing.T) zerolog.Logger {
	return zerolog.New(zerolog.NewTestWriter(t))
}
