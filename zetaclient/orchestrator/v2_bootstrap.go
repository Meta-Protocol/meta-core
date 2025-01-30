package orchestrator

import (
	"context"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	solrpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/onrik/ethrpc"
	"github.com/pkg/errors"
	tontools "github.com/tonkeeper/tongo/ton"

	"github.com/zeta-chain/node/pkg/chains"
	toncontracts "github.com/zeta-chain/node/pkg/contracts/ton"
	"github.com/zeta-chain/node/zetaclient/chains/base"
	"github.com/zeta-chain/node/zetaclient/chains/bitcoin"
	"github.com/zeta-chain/node/zetaclient/chains/bitcoin/client"
	btcobserver "github.com/zeta-chain/node/zetaclient/chains/bitcoin/observer"
	btcsigner "github.com/zeta-chain/node/zetaclient/chains/bitcoin/signer"
	"github.com/zeta-chain/node/zetaclient/chains/evm"
	evmclient "github.com/zeta-chain/node/zetaclient/chains/evm/client"
	evmobserver "github.com/zeta-chain/node/zetaclient/chains/evm/observer"
	evmsigner "github.com/zeta-chain/node/zetaclient/chains/evm/signer"
	"github.com/zeta-chain/node/zetaclient/chains/solana"
	solbserver "github.com/zeta-chain/node/zetaclient/chains/solana/observer"
	solanasigner "github.com/zeta-chain/node/zetaclient/chains/solana/signer"
	"github.com/zeta-chain/node/zetaclient/chains/ton"
	"github.com/zeta-chain/node/zetaclient/chains/ton/liteapi"
	tonobserver "github.com/zeta-chain/node/zetaclient/chains/ton/observer"
	tonsigner "github.com/zeta-chain/node/zetaclient/chains/ton/signer"
	zctx "github.com/zeta-chain/node/zetaclient/context"
	"github.com/zeta-chain/node/zetaclient/db"
	"github.com/zeta-chain/node/zetaclient/keys"
	"github.com/zeta-chain/node/zetaclient/metrics"
)

const btcBlocksPerDay = 144

func (oc *V2) bootstrapBitcoin(ctx context.Context, chain zctx.Chain) (*bitcoin.Bitcoin, error) {
	// should not happen
	if !chain.IsBitcoin() {
		return nil, errors.New("chain is not bitcoin")
	}

	app, err := zctx.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	cfg, found := app.Config().GetBTCConfig(chain.ID())
	if !found {
		return nil, errors.Wrap(errSkipChain, "unable to find btc config")
	}

	rpcClient, err := client.New(cfg, chain.ID(), oc.logger.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create rpc client")
	}

	var (
		rawChain = chain.RawChain()
		dbName   = btcDatabaseFileName(*rawChain)
	)

	baseObserver, err := oc.newBaseObserver(chain, dbName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create base observer")
	}

	observer, err := btcobserver.New(*rawChain, baseObserver, rpcClient)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create observer")
	}

	baseSigner := oc.newBaseSigner(chain)
	signer := btcsigner.New(baseSigner, rpcClient)

	return bitcoin.New(oc.scheduler, observer, signer), nil
}

func (oc *V2) bootstrapEVM(ctx context.Context, chain zctx.Chain) (*evm.EVM, error) {
	// should not happen
	if !chain.IsEVM() {
		return nil, errors.New("chain is not EVM")
	}

	app, err := zctx.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	cfg, found := app.Config().GetEVMConfig(chain.ID())
	if !found || cfg.Empty() {
		return nil, errors.Wrap(errSkipChain, "unable to find evm config")
	}

	httpClient, err := metrics.GetInstrumentedHTTPClient(cfg.Endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create http client (%s)", cfg.Endpoint)
	}

	evmClient, err := evmclient.NewFromEndpoint(ctx, cfg.Endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create evm client (%s)", cfg.Endpoint)
	}

	evmJSONRPCClient := ethrpc.NewEthRPC(cfg.Endpoint, ethrpc.WithHttpClient(httpClient))

	baseObserver, err := oc.newBaseObserver(chain, chain.Name())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create base observer")
	}

	observer, err := evmobserver.New(baseObserver, evmClient, evmJSONRPCClient)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create observer")
	}

	var (
		zetaConnectorAddress = ethcommon.HexToAddress(chain.Params().ConnectorContractAddress)
		erc20CustodyAddress  = ethcommon.HexToAddress(chain.Params().Erc20CustodyContractAddress)
		gatewayAddress       = ethcommon.HexToAddress(chain.Params().GatewayAddress)
	)

	signer, err := evmsigner.New(
		oc.newBaseSigner(chain),
		evmClient,
		zetaConnectorAddress,
		erc20CustodyAddress,
		gatewayAddress,
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create signer")
	}

	return evm.New(oc.scheduler, observer, signer), nil
}

func (oc *V2) bootstrapSolana(ctx context.Context, chain zctx.Chain) (*solana.Solana, error) {
	// should not happen
	if !chain.IsSolana() {
		return nil, errors.New("chain is not Solana")
	}

	app, err := zctx.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	cfg, found := app.Config().GetSolanaConfig()
	if !found {
		return nil, errors.Wrap(errSkipChain, "unable to find solana config")
	}

	baseObserver, err := oc.newBaseObserver(chain, chain.Name())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create base observer")
	}

	gwAddress := chain.Params().GatewayAddress

	rpcClient := solrpc.New(cfg.Endpoint)
	if rpcClient == nil {
		return nil, errors.New("unable to create rpc client")
	}

	observer, err := solbserver.NewObserver(baseObserver, rpcClient, gwAddress)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create observer")
	}

	// try loading Solana relayer key if present
	password := chain.RelayerKeyPassword()
	relayerKey, err := keys.LoadRelayerKey(app.Config().GetRelayerKeyPath(), chain.RawChain().Network, password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load relayer key")
	}

	baseSigner := oc.newBaseSigner(chain)

	// create Solana signer
	signer, err := solanasigner.New(baseSigner, rpcClient, gwAddress, relayerKey)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create signer")
	}

	return solana.New(oc.scheduler, observer, signer), nil
}

func (oc *V2) bootstrapTON(ctx context.Context, chain zctx.Chain) (*ton.TON, error) {
	// should not happen
	if !chain.IsTON() {
		return nil, errors.New("chain is not TON")
	}

	app, err := zctx.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	cfg, found := app.Config().GetTONConfig()
	if !found {
		return nil, errors.Wrap(errSkipChain, "unable to find TON config")
	}

	gatewayID, err := tontools.ParseAccountID(chain.Params().GatewayAddress)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse gateway address %q", chain.Params().GatewayAddress)
	}

	gw := toncontracts.NewGateway(gatewayID)

	lightClient, err := liteapi.NewFromSource(ctx, cfg.LiteClientConfigURL)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create TON liteapi")
	}

	baseObserver, err := oc.newBaseObserver(chain, chain.Name())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create base observer")
	}

	observer, err := tonobserver.New(baseObserver, lightClient, gw)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create observer")
	}

	signer := tonsigner.New(oc.newBaseSigner(chain), lightClient, gw)

	return ton.New(oc.scheduler, observer, signer), nil
}

func (oc *V2) newBaseObserver(chain zctx.Chain, dbName string) (*base.Observer, error) {
	var (
		rawChain       = chain.RawChain()
		rawChainParams = chain.Params()
	)

	database, err := db.NewFromSqlite(oc.deps.DBPath, dbName, true)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open database %s", dbName)
	}

	blocksCacheSize := base.DefaultBlockCacheSize
	if chain.IsBitcoin() {
		blocksCacheSize = btcBlocksPerDay
	}

	return base.NewObserver(
		*rawChain,
		*rawChainParams,
		oc.deps.Zetacore,
		oc.deps.TSS,
		blocksCacheSize,
		oc.deps.Telemetry,
		database,
		oc.logger.base,
	)
}

func (oc *V2) newBaseSigner(chain zctx.Chain) *base.Signer {
	return base.NewSigner(*chain.RawChain(), oc.deps.TSS, oc.logger.base)
}

func btcDatabaseFileName(chain chains.Chain) string {
	// legacyBTCDatabaseFilename is the Bitcoin database file name now used in mainnet and testnet3
	// so we keep using it here for backward compatibility
	const legacyBTCDatabaseFilename = "btc_chain_client"

	// For additional bitcoin networks, we use the chain name as the database file name
	switch chain.ChainId {
	case chains.BitcoinMainnet.ChainId, chains.BitcoinTestnet.ChainId:
		return legacyBTCDatabaseFilename
	default:
		return fmt.Sprintf("%s_%s", legacyBTCDatabaseFilename, chain.Name)
	}
}
