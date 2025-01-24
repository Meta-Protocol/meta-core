// Package interfaces provides interfaces for clients and signers for the chain to interact with
package interfaces

import (
	"context"
	"math/big"

	sdkmath "cosmossdk.io/math"
	cometbfttypes "github.com/cometbft/cometbft/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gagliardetto/solana-go"
	solrpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/onrik/ethrpc"
	"gitlab.com/thorchain/tss/go-tss/blame"

	"github.com/zeta-chain/node/pkg/chains"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	observertypes "github.com/zeta-chain/node/x/observer/types"
	keyinterfaces "github.com/zeta-chain/node/zetaclient/keys/interfaces"
	"github.com/zeta-chain/node/zetaclient/tss"
)

type Order string

const (
	NoOrder    Order = ""
	Ascending  Order = "ASC"
	Descending Order = "DESC"
)

// ChainObserver is the interface for chain observer
type ChainObserver interface {
	// Start starts the observer
	Start(ctx context.Context)

	// Stop stops the observer
	Stop()

	// ChainParams returns observer chain params (might be out of date with zetacore)
	ChainParams() observertypes.ChainParams

	// SetChainParams sets observer chain params
	SetChainParams(observertypes.ChainParams)

	// VoteOutboundIfConfirmed checks outbound status and returns (continueKeySign, error)
	// todo we should make this simpler.
	VoteOutboundIfConfirmed(ctx context.Context, cctx *crosschaintypes.CrossChainTx) (bool, error)
}

// ChainSigner is the interface to sign transactions for a chain
type ChainSigner interface {
	TryProcessOutbound(
		ctx context.Context,
		cctx *crosschaintypes.CrossChainTx,
		observer ChainObserver,
		zetacoreClient ZetacoreClient,
		height uint64,
	)
	SetZetaConnectorAddress(address ethcommon.Address)
	SetERC20CustodyAddress(address ethcommon.Address)
	GetZetaConnectorAddress() ethcommon.Address
	GetERC20CustodyAddress() ethcommon.Address
	SetGatewayAddress(address string)
	GetGatewayAddress() string
}

type ZetacoreVoter interface {
	PostVoteGasPrice(
		ctx context.Context,
		chain chains.Chain,
		gasPrice, priorityFee, blockNum uint64,
	) (string, error)
	PostVoteInbound(
		ctx context.Context,
		gasLimit, retryGasLimit uint64,
		msg *crosschaintypes.MsgVoteInbound,
	) (string, string, error)
	PostVoteOutbound(
		ctx context.Context,
		gasLimit, retryGasLimit uint64,
		msg *crosschaintypes.MsgVoteOutbound,
	) (string, string, error)
	PostVoteBlameData(ctx context.Context, blame *blame.Blame, chainID int64, index string) (string, error)
}

// ZetacoreClient is the client interface to interact with zetacore
//
//go:generate mockery --name ZetacoreClient --filename zetacore_client.go --case underscore --output ../../testutils/mocks
type ZetacoreClient interface {
	ZetacoreVoter

	Chain() chains.Chain
	GetKeys() keyinterfaces.ObserverKeys

	GetSupportedChains(ctx context.Context) ([]chains.Chain, error)
	GetAdditionalChains(ctx context.Context) ([]chains.Chain, error)
	GetChainParams(ctx context.Context) ([]*observertypes.ChainParams, error)

	GetKeyGen(ctx context.Context) (observertypes.Keygen, error)
	GetTSS(ctx context.Context) (observertypes.TSS, error)
	GetTSSHistory(ctx context.Context) ([]observertypes.TSS, error)
	PostVoteTSS(
		ctx context.Context,
		tssPubKey string,
		keyGenZetaHeight int64,
		status chains.ReceiveStatus,
	) (string, error)

	GetBlockHeight(ctx context.Context) (int64, error)

	ListPendingCCTX(ctx context.Context, chainID int64) ([]*crosschaintypes.CrossChainTx, uint64, error)
	ListPendingCCTXWithinRateLimit(
		ctx context.Context,
	) (*crosschaintypes.QueryListPendingCctxWithinRateLimitResponse, error)

	GetRateLimiterInput(ctx context.Context, window int64) (*crosschaintypes.QueryRateLimiterInputResponse, error)
	GetPendingNoncesByChain(ctx context.Context, chainID int64) (observertypes.PendingNonces, error)

	GetCctxByNonce(ctx context.Context, chainID int64, nonce uint64) (*crosschaintypes.CrossChainTx, error)
	GetOutboundTracker(ctx context.Context, chain chains.Chain, nonce uint64) (*crosschaintypes.OutboundTracker, error)
	GetAllOutboundTrackerByChain(
		ctx context.Context,
		chainID int64,
		order Order,
	) ([]crosschaintypes.OutboundTracker, error)
	GetCrosschainFlags(ctx context.Context) (observertypes.CrosschainFlags, error)
	GetRateLimiterFlags(ctx context.Context) (crosschaintypes.RateLimiterFlags, error)
	GetOperationalFlags(ctx context.Context) (observertypes.OperationalFlags, error)
	GetObserverList(ctx context.Context) ([]string, error)
	GetBTCTSSAddress(ctx context.Context, chainID int64) (string, error)
	GetZetaHotKeyBalance(ctx context.Context) (sdkmath.Int, error)
	GetInboundTrackersForChain(ctx context.Context, chainID int64) ([]crosschaintypes.InboundTracker, error)

	GetUpgradePlan(ctx context.Context) (*upgradetypes.Plan, error)

	PostOutboundTracker(ctx context.Context, chainID int64, nonce uint64, txHash string) (string, error)
	NewBlockSubscriber(ctx context.Context) (chan cometbfttypes.EventDataNewBlock, error)
}

// EVMRPCClient is the interface for EVM RPC client
type EVMRPCClient interface {
	bind.ContractBackend
	SendTransaction(ctx context.Context, tx *ethtypes.Transaction) error
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*ethtypes.Block, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*ethtypes.Header, error)
	TransactionByHash(ctx context.Context, hash ethcommon.Hash) (tx *ethtypes.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash ethcommon.Hash) (*ethtypes.Receipt, error)
	TransactionSender(
		ctx context.Context,
		tx *ethtypes.Transaction,
		block ethcommon.Hash,
		index uint,
	) (ethcommon.Address, error)
}

// SolanaRPCClient is the interface for Solana RPC client
type SolanaRPCClient interface {
	GetVersion(ctx context.Context) (*solrpc.GetVersionResult, error)
	GetHealth(ctx context.Context) (string, error)
	GetSlot(ctx context.Context, commitment solrpc.CommitmentType) (uint64, error)
	GetBlockTime(ctx context.Context, block uint64) (*solana.UnixTimeSeconds, error)
	GetAccountInfo(ctx context.Context, account solana.PublicKey) (*solrpc.GetAccountInfoResult, error)
	GetBalance(
		ctx context.Context,
		account solana.PublicKey,
		commitment solrpc.CommitmentType,
	) (*solrpc.GetBalanceResult, error)
	GetLatestBlockhash(ctx context.Context, commitment solrpc.CommitmentType) (*solrpc.GetLatestBlockhashResult, error)
	GetRecentPrioritizationFees(
		ctx context.Context,
		accounts solana.PublicKeySlice,
	) ([]solrpc.PriorizationFeeResult, error)
	GetTransaction(
		ctx context.Context,
		txSig solana.Signature, // transaction signature
		opts *solrpc.GetTransactionOpts,
	) (*solrpc.GetTransactionResult, error)
	GetConfirmedTransactionWithOpts(
		ctx context.Context,
		signature solana.Signature,
		opts *solrpc.GetTransactionOpts,
	) (*solrpc.TransactionWithMeta, error)
	GetSignaturesForAddressWithOpts(
		ctx context.Context,
		account solana.PublicKey,
		opts *solrpc.GetSignaturesForAddressOpts,
	) ([]*solrpc.TransactionSignature, error)
	SendTransactionWithOpts(
		ctx context.Context,
		transaction *solana.Transaction,
		opts solrpc.TransactionOpts,
	) (solana.Signature, error)
}

// EVMJSONRPCClient is the interface for EVM JSON RPC client
type EVMJSONRPCClient interface {
	EthGetBlockByNumber(number int, withTransactions bool) (*ethrpc.Block, error)
	EthGetTransactionByHash(hash string) (*ethrpc.Transaction, error)
}

// TSSSigner is the interface for TSS signer
type TSSSigner interface {
	PubKey() tss.PubKey
	Sign(ctx context.Context, data []byte, height, nonce uint64, chainID int64) ([65]byte, error)
	SignBatch(ctx context.Context, digests [][]byte, height, nonce uint64, chainID int64) ([][65]byte, error)
}
