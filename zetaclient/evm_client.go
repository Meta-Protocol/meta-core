package zetaclient

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/zeta.non-eth.sol"
	zetaconnectoreth "github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/zetaconnector.eth.sol"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/erc20custody.sol"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/zetaconnector.non-eth.sol"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	metricsPkg "github.com/zeta-chain/zetacore/zetaclient/metrics"
	clienttypes "github.com/zeta-chain/zetacore/zetaclient/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TxHashEnvelope struct {
	TxHash string
	Done   chan struct{}
}

type OutTx struct {
	SendHash string
	TxHash   string
	Nonce    int64
}
type EVMLog struct {
	ChainLogger          zerolog.Logger // Parent logger
	ExternalChainWatcher zerolog.Logger // Observes external Chains for incoming trasnactions
	WatchGasPrice        zerolog.Logger // Observes external Chains for Gas prices and posts to core
	ObserveOutTx         zerolog.Logger // Observes external Chains for Outgoing transactions

}

const (
	DonationMessage = "I am rich!"
)

// EVMChainClient represents the chain configuration for an EVM chain
// Filled with above constants depending on chain
type EVMChainClient struct {
	*ChainMetrics
	chain                     common.Chain
	evmClient                 EVMRPCClient
	KlaytnClient              KlaytnRPCClient
	zetaClient                ZetaCoreBridger
	Tss                       TSSSigner
	lastBlockScanned          uint64
	lastBlock                 uint64
	BlockTimeExternalChain    uint64 // block time in seconds
	txWatchList               map[ethcommon.Hash]string
	Mu                        *sync.Mutex
	db                        *gorm.DB
	outTxPendingTransaction   map[string]*ethtypes.Transaction
	outTXConfirmedReceipts    map[string]*ethtypes.Receipt
	outTXConfirmedTransaction map[string]*ethtypes.Transaction
	MinNonce                  int64
	MaxNonce                  int64
	OutTxChan                 chan OutTx // send to this channel if you want something back!
	stop                      chan struct{}
	fileLogger                *zerolog.Logger // for critical info
	logger                    EVMLog
	cfg                       *config.Config
	params                    observertypes.CoreParams
	ts                        *TelemetryServer

	BlockCache *lru.Cache
}

var _ ChainClient = (*EVMChainClient)(nil)

// NewEVMChainClient returns a new configuration based on supplied target chain
func NewEVMChainClient(
	bridge ZetaCoreBridger,
	tss TSSSigner,
	dbpath string,
	metrics *metricsPkg.Metrics,
	logger zerolog.Logger,
	cfg *config.Config,
	evmCfg config.EVMConfig,
	ts *TelemetryServer,
) (*EVMChainClient, error) {
	ob := EVMChainClient{
		ChainMetrics: NewChainMetrics(evmCfg.Chain.ChainName.String(), metrics),
		ts:           ts,
	}
	chainLogger := logger.With().Str("chain", evmCfg.Chain.ChainName.String()).Logger()
	ob.logger = EVMLog{
		ChainLogger:          chainLogger,
		ExternalChainWatcher: chainLogger.With().Str("module", "ExternalChainWatcher").Logger(),
		WatchGasPrice:        chainLogger.With().Str("module", "WatchGasPrice").Logger(),
		ObserveOutTx:         chainLogger.With().Str("module", "ObserveOutTx").Logger(),
	}
	ob.cfg = cfg
	ob.params = evmCfg.CoreParams
	ob.stop = make(chan struct{})
	ob.chain = evmCfg.Chain
	ob.Mu = &sync.Mutex{}
	ob.zetaClient = bridge
	ob.txWatchList = make(map[ethcommon.Hash]string)
	ob.Tss = tss
	ob.outTxPendingTransaction = make(map[string]*ethtypes.Transaction)
	ob.outTXConfirmedReceipts = make(map[string]*ethtypes.Receipt)
	ob.outTXConfirmedTransaction = make(map[string]*ethtypes.Transaction)
	ob.OutTxChan = make(chan OutTx, 100)

	logFile, err := os.OpenFile(ob.chain.ChainName.String()+"_debug.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Error().Err(err).Msgf("there was an error creating a logFile chain %s", ob.chain.ChainName.String())
	}
	fileLogger := zerolog.New(logFile).With().Logger()
	ob.fileLogger = &fileLogger

	ob.logger.ChainLogger.Info().Msgf("Chain %s endpoint %s", ob.chain.ChainName.String(), evmCfg.Endpoint)
	client, err := ethclient.Dial(evmCfg.Endpoint)
	if err != nil {
		ob.logger.ChainLogger.Error().Err(err).Msg("eth Client Dial")
		return nil, err
	}
	ob.evmClient = client

	ob.BlockCache, err = lru.New(1000)
	if err != nil {
		ob.logger.ChainLogger.Error().Err(err).Msg("failed to create block cache")
		return nil, err
	}

	if ob.chain.IsKlaytnChain() {
		client, err := Dial(evmCfg.Endpoint)
		if err != nil {
			ob.logger.ChainLogger.Err(err).Msg("klaytn Client Dial")
			return nil, err
		}
		ob.KlaytnClient = client
	}

	// create metric counters
	err = ob.RegisterPromCounter("rpc_getLogs_count", "Number of getLogs")
	if err != nil {
		return nil, err
	}
	err = ob.RegisterPromCounter("rpc_getBlockByNumber_count", "Number of getBlockByNumber")
	if err != nil {
		return nil, err
	}
	err = ob.RegisterPromGauge(metricsPkg.PendingTxs, "Number of pending transactions")
	if err != nil {
		return nil, err
	}

	err = ob.LoadDB(dbpath, ob.chain)
	if err != nil {
		return nil, err
	}

	ob.logger.ChainLogger.Info().Msgf("%s: start scanning from block %d", ob.chain.String(), ob.GetLastBlockHeightScanned())

	return &ob, nil
}
func (ob *EVMChainClient) WithChain(chain common.Chain) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.chain = chain
}
func (ob *EVMChainClient) WithLogger(logger zerolog.Logger) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.logger = EVMLog{
		ChainLogger:          logger,
		ExternalChainWatcher: logger.With().Str("module", "ExternalChainWatcher").Logger(),
		WatchGasPrice:        logger.With().Str("module", "WatchGasPrice").Logger(),
		ObserveOutTx:         logger.With().Str("module", "ObserveOutTx").Logger(),
	}
}

func (ob *EVMChainClient) WithEvmClient(client *ethclient.Client) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.evmClient = client
}

func (ob *EVMChainClient) WithZetaClient(bridge *ZetaCoreBridge) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.zetaClient = bridge
}

func (ob *EVMChainClient) WithParams(params observertypes.CoreParams) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.params = params
}

func (ob *EVMChainClient) SetConfig(cfg *config.Config) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.cfg = cfg
}

func (ob *EVMChainClient) SetCoreParams(params observertypes.CoreParams) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	ob.params = params
}

func (ob *EVMChainClient) GetCoreParams() observertypes.CoreParams {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()
	return ob.params
}

func (ob *EVMChainClient) GetConnectorContract() (*zetaconnector.ZetaConnectorNonEth, error) {
	addr := ethcommon.HexToAddress(ob.GetCoreParams().ConnectorContractAddress)
	return FetchConnectorContract(addr, ob.evmClient)
}

func (ob *EVMChainClient) GetConnectorContractEth() (*zetaconnectoreth.ZetaConnectorEth, error) {
	addr := ethcommon.HexToAddress(ob.GetCoreParams().ConnectorContractAddress)
	return FetchConnectorContractEth(addr, ob.evmClient)
}

func (ob *EVMChainClient) GetZetaTokenNonEthContract() (*zeta.ZetaNonEth, error) {
	addr := ethcommon.HexToAddress(ob.GetCoreParams().ZetaTokenContractAddress)
	return FetchZetaZetaNonEthTokenContract(addr, ob.evmClient)
}

func (ob *EVMChainClient) GetERC20CustodyContract() (*erc20custody.ERC20Custody, error) {
	addr := ethcommon.HexToAddress(ob.GetCoreParams().Erc20CustodyContractAddress)
	return FetchERC20CustodyContract(addr, ob.evmClient)
}

func FetchConnectorContract(addr ethcommon.Address, client EVMRPCClient) (*zetaconnector.ZetaConnectorNonEth, error) {
	return zetaconnector.NewZetaConnectorNonEth(addr, client)
}

func FetchConnectorContractEth(addr ethcommon.Address, client EVMRPCClient) (*zetaconnectoreth.ZetaConnectorEth, error) {
	return zetaconnectoreth.NewZetaConnectorEth(addr, client)
}

func FetchZetaZetaNonEthTokenContract(addr ethcommon.Address, client EVMRPCClient) (*zeta.ZetaNonEth, error) {
	return zeta.NewZetaNonEth(addr, client)
}

func FetchERC20CustodyContract(addr ethcommon.Address, client EVMRPCClient) (*erc20custody.ERC20Custody, error) {
	return erc20custody.NewERC20Custody(addr, client)
}

func (ob *EVMChainClient) Start() {
	go ob.ExternalChainWatcherForNewInboundTrackerSuggestions()
	go ob.ExternalChainWatcher() // Observes external Chains for incoming trasnactions
	go ob.WatchGasPrice()        // Observes external Chains for Gas prices and posts to core
	go ob.observeOutTx()         // Populates receipts and confirmed outbound transactions
}

func (ob *EVMChainClient) Stop() {
	ob.logger.ChainLogger.Info().Msgf("ob %s is stopping", ob.chain.String())
	close(ob.stop) // this notifies all goroutines to stop

	ob.logger.ChainLogger.Info().Msg("closing ob.db")
	dbInst, err := ob.db.DB()
	if err != nil {
		ob.logger.ChainLogger.Info().Msg("error getting database instance")
	}
	err = dbInst.Close()
	if err != nil {
		ob.logger.ChainLogger.Error().Err(err).Msg("error closing database")
	}

	ob.logger.ChainLogger.Info().Msgf("%s observer stopped", ob.chain.String())
}

// returns: isIncluded, isConfirmed, Error
// If isConfirmed, it also post to ZetaCore
func (ob *EVMChainClient) IsSendOutTxProcessed(sendHash string, nonce uint64, cointype common.CoinType, logger zerolog.Logger) (bool, bool, error) {
	if !ob.hasTxConfirmed(nonce) {
		return false, false, nil
	}
	params := ob.GetCoreParams()
	receipt, transaction := ob.GetTxNReceipt(nonce)

	sendID := fmt.Sprintf("%s-%d", ob.chain.String(), nonce)
	logger = logger.With().Str("sendID", sendID).Logger()
	if cointype == common.CoinType_Cmd {
		recvStatus := common.ReceiveStatus_Failed
		if receipt.Status == 1 {
			recvStatus = common.ReceiveStatus_Success
		}
		zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
			sendHash,
			receipt.TxHash.Hex(),
			receipt.BlockNumber.Uint64(),
			receipt.GasUsed,
			transaction.GasPrice(),
			transaction.Gas(),
			transaction.Value(),
			recvStatus,
			ob.chain,
			nonce,
			common.CoinType_Cmd,
		)
		if err != nil {
			logger.Error().Err(err).Msgf("error posting confirmation to meta core for cctx %s nonce %d", sendHash, nonce)
		} else if zetaTxHash != "" {
			logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
		}
		return true, true, nil

	} else if cointype == common.CoinType_Gas { // the outbound is a regular Ether/BNB/Matic transfer; no need to check events
		if receipt.Status == 1 {
			zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
				sendHash,
				receipt.TxHash.Hex(),
				receipt.BlockNumber.Uint64(),
				receipt.GasUsed,
				transaction.GasPrice(),
				transaction.Gas(),
				transaction.Value(),
				common.ReceiveStatus_Success,
				ob.chain,
				nonce,
				common.CoinType_Gas,
			)
			if err != nil {
				logger.Error().Err(err).Msgf("error posting confirmation to meta core for cctx %s nonce %d", sendHash, nonce)
			} else if zetaTxHash != "" {
				logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
			}
			return true, true, nil
		} else if receipt.Status == 0 { // the same as below events flow
			logger.Info().Msgf("Found (failed tx) sendHash %s on chain %s txhash %s", sendHash, ob.chain.String(), receipt.TxHash.Hex())
			zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
				sendHash,
				receipt.TxHash.Hex(),
				receipt.BlockNumber.Uint64(),
				receipt.GasUsed,
				transaction.GasPrice(),
				transaction.Gas(),
				big.NewInt(0),
				common.ReceiveStatus_Failed,
				ob.chain,
				nonce,
				common.CoinType_Gas,
			)
			if err != nil {
				logger.Error().Err(err).Msgf("PostReceiveConfirmation error in WatchTxHashWithTimeout; zeta tx hash %s cctx %s nonce %d", zetaTxHash, sendHash, nonce)
			} else if zetaTxHash != "" {
				logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
			}
			return true, true, nil
		}
	} else if cointype == common.CoinType_Zeta { // the outbound is a Zeta transfer; need to check events ZetaReceived
		if receipt.Status == 1 {
			logs := receipt.Logs
			for _, vLog := range logs {
				confHeight := vLog.BlockNumber + params.ConfirmationCount
				// TODO rewrite this to return early if not confirmed
				connector, err := ob.GetConnectorContract()
				if err != nil {
					return false, false, fmt.Errorf("error getting connector contract: %w", err)
				}
				receivedLog, err := connector.ZetaConnectorNonEthFilterer.ParseZetaReceived(*vLog)
				if err == nil {
					logger.Info().Msgf("Found (outTx) sendHash %s on chain %s txhash %s", sendHash, ob.chain.String(), vLog.TxHash.Hex())
					if confHeight <= ob.GetLastBlockHeight() {
						logger.Info().Msg("Confirmed! Sending PostConfirmation to zetacore...")
						if len(vLog.Topics) != 4 {
							logger.Error().Msgf("wrong number of topics in log %d", len(vLog.Topics))
							return false, false, fmt.Errorf("wrong number of topics in log %d", len(vLog.Topics))
						}
						sendhash := vLog.Topics[3].Hex()
						//var rxAddress string = ethcommon.HexToAddress(vLog.Topics[1].Hex()).Hex()
						mMint := receivedLog.ZetaValue
						zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
							sendhash,
							vLog.TxHash.Hex(),
							vLog.BlockNumber,
							receipt.GasUsed,
							transaction.GasPrice(),
							transaction.Gas(),
							mMint,
							common.ReceiveStatus_Success,
							ob.chain,
							nonce,
							common.CoinType_Zeta,
						)
						if err != nil {
							logger.Error().Err(err).Msgf("error posting confirmation to meta core for cctx %s nonce %d", sendHash, nonce)
							continue
						} else if zetaTxHash != "" {
							logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
						}
						return true, true, nil
					}
					logger.Info().Msgf("Included; %d blocks before confirmed! chain %s nonce %d", confHeight-ob.GetLastBlockHeight(), ob.chain.String(), nonce)
					return true, false, nil
				}
				revertedLog, err := connector.ZetaConnectorNonEthFilterer.ParseZetaReverted(*vLog)
				if err == nil {
					logger.Info().Msgf("Found (revertTx) sendHash %s on chain %s txhash %s", sendHash, ob.chain.String(), vLog.TxHash.Hex())
					if confHeight <= ob.GetLastBlockHeight() {
						logger.Info().Msg("Confirmed! Sending PostConfirmation to zetacore...")
						if len(vLog.Topics) != 3 {
							logger.Error().Msgf("wrong number of topics in log %d", len(vLog.Topics))
							return false, false, fmt.Errorf("wrong number of topics in log %d", len(vLog.Topics))
						}
						sendhash := vLog.Topics[2].Hex()
						mMint := revertedLog.RemainingZetaValue
						zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
							sendhash,
							vLog.TxHash.Hex(),
							vLog.BlockNumber,
							receipt.GasUsed,
							transaction.GasPrice(),
							transaction.Gas(),
							mMint,
							common.ReceiveStatus_Success,
							ob.chain,
							nonce,
							common.CoinType_Zeta,
						)
						if err != nil {
							logger.Err(err).Msgf("error posting confirmation to meta core for cctx %s nonce %d", sendHash, nonce)
							continue
						} else if zetaTxHash != "" {
							logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
						}
						return true, true, nil
					}
					logger.Info().Msgf("Included; %d blocks before confirmed! chain %s nonce %d", confHeight-ob.GetLastBlockHeight(), ob.chain.String(), nonce)
					return true, false, nil
				}
			}
		} else if receipt.Status == 0 {
			//FIXME: check nonce here by getTransaction RPC
			logger.Info().Msgf("Found (failed tx) sendHash %s on chain %s txhash %s", sendHash, ob.chain.String(), receipt.TxHash.Hex())
			zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
				sendHash,
				receipt.TxHash.Hex(),
				receipt.BlockNumber.Uint64(),
				receipt.GasUsed,
				transaction.GasPrice(),
				transaction.Gas(),
				big.NewInt(0),
				common.ReceiveStatus_Failed,
				ob.chain,
				nonce,
				common.CoinType_Zeta,
			)
			if err != nil {
				logger.Error().Err(err).Msgf("error posting confirmation to meta core for cctx %s nonce %d", sendHash, nonce)
			} else if zetaTxHash != "" {
				logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
			}
			return true, true, nil
		}
	} else if cointype == common.CoinType_ERC20 {
		if receipt.Status == 1 {
			logs := receipt.Logs
			ERC20Custody, err := ob.GetERC20CustodyContract()
			if err != nil {
				logger.Warn().Msgf("NewERC20Custody err: %s", err)
			}
			for _, vLog := range logs {
				event, err := ERC20Custody.ParseWithdrawn(*vLog)
				confHeight := vLog.BlockNumber + params.ConfirmationCount
				if err == nil {
					logger.Info().Msgf("Found (ERC20Custody.Withdrawn Event) sendHash %s on chain %s txhash %s", sendHash, ob.chain.String(), vLog.TxHash.Hex())
					if confHeight <= ob.GetLastBlockHeight() {
						logger.Info().Msg("Confirmed! Sending PostConfirmation to zetacore...")
						zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
							sendHash,
							vLog.TxHash.Hex(),
							vLog.BlockNumber,
							receipt.GasUsed,
							transaction.GasPrice(),
							transaction.Gas(),
							event.Amount,
							common.ReceiveStatus_Success,
							ob.chain,
							nonce,
							common.CoinType_ERC20,
						)
						if err != nil {
							logger.Error().Err(err).Msgf("error posting confirmation to meta core for cctx %s nonce %d", sendHash, nonce)
							continue
						} else if zetaTxHash != "" {
							logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
						}
						return true, true, nil
					}
					logger.Info().Msgf("Included; %d blocks before confirmed! chain %s nonce %d", confHeight-ob.GetLastBlockHeight(), ob.chain.String(), nonce)
					return true, false, nil
				}
			}
		} else {
			logger.Info().Msgf("Found (failed tx) sendHash %s on chain %s txhash %s", sendHash, ob.chain.String(), receipt.TxHash.Hex())
			zetaTxHash, ballot, err := ob.zetaClient.PostReceiveConfirmation(
				sendHash,
				receipt.TxHash.Hex(),
				receipt.BlockNumber.Uint64(),
				receipt.GasUsed,
				transaction.GasPrice(),
				transaction.Gas(),
				big.NewInt(0),
				common.ReceiveStatus_Failed,
				ob.chain,
				nonce,
				common.CoinType_ERC20,
			)
			if err != nil {
				logger.Error().Err(err).Msgf("PostReceiveConfirmation error in WatchTxHashWithTimeout; zeta tx hash %s", zetaTxHash)
			} else if zetaTxHash != "" {
				logger.Info().Msgf("Zeta tx hash: %s cctx %s nonce %d ballot %s", zetaTxHash, sendHash, nonce, ballot)
			}
			return true, true, nil
		}
	}

	return false, false, nil
}

// The lowest nonce we observe outTx for each chain
var lowestOutTxNonceToObserve = map[int64]uint64{
	5:     113000, // Goerli
	97:    102600, // BSC testnet
	80001: 154500, // Mumbai
}

// FIXME: there's a chance that a txhash in OutTxChan may not deliver when Stop() is called
// observeOutTx periodically checks all the txhash in potential outbound txs
func (ob *EVMChainClient) observeOutTx() {
	// read env variables if set
	timeoutNonce, err := strconv.Atoi(os.Getenv("OS_TIMEOUT_NONCE"))
	if err != nil || timeoutNonce <= 0 {
		timeoutNonce = 100 * 3 // process up to 100 hashes
	}
	rpcRestTime, err := strconv.Atoi(os.Getenv("OS_RPC_REST_TIME"))
	if err != nil || rpcRestTime <= 0 {
		rpcRestTime = 20 // 20ms
	}
	ob.logger.ObserveOutTx.Info().Msgf("observeOutTx using timeoutNonce %d seconds, rpcRestTime %d ms", timeoutNonce, rpcRestTime)

	ticker, err := NewDynamicTicker(fmt.Sprintf("EVM_observeOutTx_%d", ob.chain.ChainId), ob.GetCoreParams().OutTxTicker)
	if err != nil {
		ob.logger.ObserveOutTx.Error().Err(err).Msg("failed to create ticker")
		return
	}

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C():
			trackers, err := ob.zetaClient.GetAllOutTxTrackerByChain(ob.chain.ChainId, Ascending)
			if err != nil {
				continue
			}
			//FIXME: remove this timeout here to ensure that all trackers are queried
			outTimeout := time.After(time.Duration(timeoutNonce) * time.Second)
		TRACKERLOOP:
			// Skip old gabbage trackers as we spent too much time on querying them
			for _, tracker := range trackers {
				nonceInt := tracker.Nonce
				if nonceInt < lowestOutTxNonceToObserve[ob.chain.ChainId] {
					continue
				}
				if ob.hasTxConfirmed(nonceInt) { // Go to next tracker if this one already has a confirmed tx
					continue
				}
				for _, txHash := range tracker.HashList {
					select {
					case <-outTimeout:
						ob.logger.ObserveOutTx.Warn().Msgf("observeOutTx timeout on chain %d nonce %d", ob.chain.ChainId, nonceInt)
						break TRACKERLOOP
					default:
						err := ob.confirmTxByHash(txHash.TxHash, nonceInt)
						time.Sleep(time.Duration(rpcRestTime) * time.Millisecond)
						if err == nil { // confirmed
							ob.logger.ObserveOutTx.Info().Msgf("observeOutTx confirmed outTx %s for chain %d nonce %d", txHash.TxHash, ob.chain.ChainId, nonceInt)
							break
						}
						ob.logger.ObserveOutTx.Debug().Err(err).Msgf("error confirmTxByHash: chain %s hash %s", ob.chain.String(), txHash.TxHash)
					}
				}
			}
			ticker.UpdateInterval(ob.GetCoreParams().OutTxTicker, ob.logger.ObserveOutTx)
		case <-ob.stop:
			ob.logger.ObserveOutTx.Info().Msg("observeOutTx: stopped")
			return
		}
	}
}

// SetPendingTx sets the pending transaction in memory
func (ob *EVMChainClient) SetPendingTx(nonce uint64, transaction *ethtypes.Transaction) {
	ob.Mu.Lock()
	ob.outTxPendingTransaction[ob.GetTxID(nonce)] = transaction
	ob.Mu.Unlock()
}

// GetPendingTx gets the pending transaction from memory
func (ob *EVMChainClient) GetPendingTx(nonce uint64) *ethtypes.Transaction {
	ob.Mu.Lock()
	transaction := ob.outTxPendingTransaction[ob.GetTxID(nonce)]
	ob.Mu.Unlock()
	return transaction
}

// SetTxNReceipt sets the receipt and transaction in memory
func (ob *EVMChainClient) SetTxNReceipt(nonce uint64, receipt *ethtypes.Receipt, transaction *ethtypes.Transaction) {
	ob.Mu.Lock()
	delete(ob.outTxPendingTransaction, ob.GetTxID(nonce)) // remove pending transaction, if any
	ob.outTXConfirmedReceipts[ob.GetTxID(nonce)] = receipt
	ob.outTXConfirmedTransaction[ob.GetTxID(nonce)] = transaction
	ob.Mu.Unlock()
}

// getTxNReceipt gets the receipt and transaction from memory
func (ob *EVMChainClient) GetTxNReceipt(nonce uint64) (*ethtypes.Receipt, *ethtypes.Transaction) {
	ob.Mu.Lock()
	receipt := ob.outTXConfirmedReceipts[ob.GetTxID(nonce)]
	transaction := ob.outTXConfirmedTransaction[ob.GetTxID(nonce)]
	ob.Mu.Unlock()
	return receipt, transaction
}

// hasTxConfirmed returns true if there is a confirmed tx for 'nonce'
func (ob *EVMChainClient) hasTxConfirmed(nonce uint64) bool {
	ob.Mu.Lock()
	confirmed := ob.outTXConfirmedReceipts[ob.GetTxID(nonce)] != nil && ob.outTXConfirmedTransaction[ob.GetTxID(nonce)] != nil
	ob.Mu.Unlock()
	return confirmed
}

// confirmTxByHash checks if a txHash is confirmed and saves transaction and receipt in memory
// returns nil if confirmed or error otherwise
func (ob *EVMChainClient) confirmTxByHash(txHash string, nonce uint64) error {
	ctxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query transaction
	transaction, isPending, err := ob.evmClient.TransactionByHash(ctxt, ethcommon.HexToHash(txHash))
	if err != nil {
		return err
	}
	if transaction == nil { // should not happen
		log.Error().Msgf("confirmTxByHash: transaction is nil for txHash %s nonce %d", txHash, nonce)
		return fmt.Errorf("transaction is nil")
	}
	if isPending { // save pending transaction
		ob.SetPendingTx(nonce, transaction)
		return fmt.Errorf("confirmTxByHash: txHash %s nonce %d is pending", txHash, nonce)
	}

	// query receipt
	receipt, err := ob.evmClient.TransactionReceipt(ctxt, ethcommon.HexToHash(txHash))
	if err != nil {
		if err != ethereum.NotFound {
			log.Warn().Err(err).Msgf("confirmTxByHash: TransactionReceipt error, txHash %s nonce %d", txHash, nonce)
		}
		return err
	}
	if receipt == nil { // should not happen
		log.Error().Msgf("confirmTxByHash: receipt is nil for txHash %s nonce %d", txHash, nonce)
		return fmt.Errorf("receipt is nil")
	}

	// check nonce and confirmations
	if transaction.Nonce() != nonce {
		log.Error().Msgf("confirmTxByHash: txHash %s nonce mismatch: wanted %d, got tx nonce %d", txHash, nonce, transaction.Nonce())
		return fmt.Errorf("outtx nonce mismatch")
	}
	confHeight := receipt.BlockNumber.Uint64() + ob.GetCoreParams().ConfirmationCount
	if confHeight >= math.MaxInt64 {
		return fmt.Errorf("confirmTxByHash: confHeight is out of range")
	}
	if confHeight > ob.GetLastBlockHeight() {
		log.Info().Msgf("confirmTxByHash: txHash %s nonce %d included but not confirmed: receipt block %d, current block %d",
			txHash, nonce, receipt.BlockNumber, ob.GetLastBlockHeight())
		return fmt.Errorf("included but not confirmed")
	}

	// confirmed, save receipt and transaction
	ob.SetTxNReceipt(nonce, receipt, transaction)

	return nil
}

// SetLastBlockHeightScanned set last block height scanned (not necessarily caught up with external block; could be slow/paused)
func (ob *EVMChainClient) SetLastBlockHeightScanned(height uint64) {
	atomic.StoreUint64(&ob.lastBlockScanned, height)
	ob.ts.SetLastScannedBlockNumber(ob.chain.ChainId, height)
}

// GetLastBlockHeightScanned get last block height scanned (not necessarily caught up with external block; could be slow/paused)
func (ob *EVMChainClient) GetLastBlockHeightScanned() uint64 {
	height := atomic.LoadUint64(&ob.lastBlockScanned)
	return height
}

// SetLastBlockHeight set external last block height
func (ob *EVMChainClient) SetLastBlockHeight(height uint64) {
	if height >= math.MaxInt64 {
		panic("lastBlock is too large")
	}
	atomic.StoreUint64(&ob.lastBlock, height)
}

// GetLastBlockHeight get external last block height
func (ob *EVMChainClient) GetLastBlockHeight() uint64 {
	height := atomic.LoadUint64(&ob.lastBlock)
	if height >= math.MaxInt64 {
		panic("lastBlock is too large")
	}
	return height
}

func (ob *EVMChainClient) ExternalChainWatcher() {
	// At each tick, query the Connector contract
	ticker, err := NewDynamicTicker(fmt.Sprintf("EVM_ExternalChainWatcher_%d", ob.chain.ChainId), ob.GetCoreParams().InTxTicker)
	if err != nil {
		ob.logger.ExternalChainWatcher.Error().Err(err).Msg("NewDynamicTicker error")
		return
	}

	defer ticker.Stop()
	ob.logger.ExternalChainWatcher.Info().Msg("ExternalChainWatcher started")
	for {
		select {
		case <-ticker.C():
			err := ob.observeInTX()
			if err != nil {
				ob.logger.ExternalChainWatcher.Err(err).Msg("observeInTX error")
			}
			ticker.UpdateInterval(ob.GetCoreParams().InTxTicker, ob.logger.ExternalChainWatcher)
		case <-ob.stop:
			ob.logger.ExternalChainWatcher.Info().Msg("ExternalChainWatcher stopped")
			return
		}
	}
}

func (ob *EVMChainClient) postBlockHeader(tip uint64) error {
	bn := tip

	res, err := ob.zetaClient.GetBlockHeaderStateByChain(ob.chain.ChainId)
	if err == nil && res.BlockHeaderState != nil && res.BlockHeaderState.EarliestHeight > 0 {
		// #nosec G701 always positive
		bn = uint64(res.BlockHeaderState.LatestHeight)
	}

	if bn > tip {
		return fmt.Errorf("postBlockHeader: must post block confirmed block header: %d > %d", bn, tip)
	}

	block, err := ob.GetBlockByNumberCached(bn)
	if err != nil {
		ob.logger.ExternalChainWatcher.Error().Err(err).Msgf("error getting block: %d", bn)
		return err
	}
	headerRLP, err := rlp.EncodeToBytes(block.Header())
	if err != nil {
		ob.logger.ExternalChainWatcher.Error().Err(err).Msgf("error encoding block header: %d", bn)
		return err
	}

	_, err = ob.zetaClient.PostAddBlockHeader(
		ob.chain.ChainId,
		block.Hash().Bytes(),
		block.Number().Int64(),
		common.NewEthereumHeader(headerRLP),
	)
	if err != nil {
		ob.logger.ExternalChainWatcher.Error().Err(err).Msgf("error posting block header: %d", bn)
		return err
	}
	return nil
}

func (ob *EVMChainClient) observeInTX() error {
	header, err := ob.evmClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	// update last block height
	ob.SetLastBlockHeight(header.Number.Uint64())
	confirmedBlockNum := header.Number.Uint64() - ob.GetCoreParams().ConfirmationCount

	crosschainFlags, err := ob.zetaClient.GetCrosschainFlags()
	if err != nil {
		return err
	}
	if !crosschainFlags.IsInboundEnabled {
		return errors.New("inbound TXS / Send has been disabled by the protocol")
	}
	counter, err := ob.GetPromCounter("rpc_getBlockByNumber_count")
	if err != nil {
		ob.logger.ExternalChainWatcher.Error().Err(err).Msg("GetPromCounter:")
	}
	counter.Inc()

	// skip if no new block is produced.
	sampledLogger := ob.logger.ExternalChainWatcher.Sample(&zerolog.BasicSampler{N: 10})
	if confirmedBlockNum <= ob.GetLastBlockHeightScanned() {
		sampledLogger.Debug().Msg("Skipping observer , No new block is produced")
		return nil
	}
	lastBlock := ob.GetLastBlockHeightScanned()
	startBlock := lastBlock + 1
	toBlock := lastBlock + config.MaxBlocksPerPeriod // read at most 100 blocks in one go
	if toBlock > confirmedBlockNum {
		toBlock = confirmedBlockNum
	}
	ob.logger.ExternalChainWatcher.Info().Msgf("Checking for all inTX : startBlock %d, toBlock %d", startBlock, toBlock)
	//task 1:  Query evm chain for zeta sent logs
	func() {
		toB := toBlock
		connector, err := ob.GetConnectorContract()
		if err != nil {
			ob.logger.ChainLogger.Warn().Err(err).Msgf("observeInTx: GetConnectorContract error:")
			return
		}
		cnt, err := ob.GetPromCounter("rpc_getLogs_count")
		if err != nil {
			ob.logger.ExternalChainWatcher.Error().Err(err).Msg("GetPromCounter:")
		} else {
			cnt.Inc()
		}
		logs, err := connector.FilterZetaSent(&bind.FilterOpts{
			Start:   startBlock,
			End:     &toB,
			Context: context.TODO(),
		}, []ethcommon.Address{}, []*big.Int{})
		if err != nil {
			ob.logger.ChainLogger.Warn().Err(err).Msgf("observeInTx: FilterZetaSent error:")
			return
		}
		// Pull out arguments from logs
		for logs.Next() {
			msg, err := ob.GetInboundVoteMsgForZetaSentEvent(logs.Event)
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msg("error getting inbound vote msg")
				continue
			}

			zetaHash, err := ob.zetaClient.PostSend(PostSendNonEVMGasLimit, &msg)
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msg("error posting to zeta core")
				return
			}
			ob.logger.ExternalChainWatcher.Info().Msgf("ZetaSent event detected and reported: PostSend zeta tx: %s", zetaHash)
		}
	}()

	// task 2: Query evm chain for deposited logs
	func() {
		toB := toBlock
		erc20custody, err := ob.GetERC20CustodyContract()
		if err != nil {
			ob.logger.ExternalChainWatcher.Warn().Err(err).Msgf("observeInTx: GetERC20CustodyContract error:")
			return
		}
		depositedLogs, err := erc20custody.FilterDeposited(&bind.FilterOpts{
			Start:   startBlock,
			End:     &toB,
			Context: context.TODO(),
		}, []ethcommon.Address{})

		if err != nil {
			ob.logger.ExternalChainWatcher.Warn().Err(err).Msgf("observeInTx: FilterDeposited error:")
			return
		}
		cnt, err := ob.GetPromCounter("rpc_getLogs_count")
		if err != nil {
			ob.logger.ExternalChainWatcher.Error().Err(err).Msg("GetPromCounter:")
		} else {
			cnt.Inc()
		}

		// Pull out arguments from logs
		for depositedLogs.Next() {
			msg, err := ob.GetInboundVoteMsgForDepositedEvent(depositedLogs.Event)
			if err != nil {
				continue
			}
			zetaHash, err := ob.zetaClient.PostSend(PostSendEVMGasLimit, &msg)
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msg("error posting to zeta core")
				return
			}
			ob.logger.ExternalChainWatcher.Info().Msgf("ZRC20Custody Deposited event detected and reported: PostSend zeta tx: %s", zetaHash)
		}
	}()

	// task 3: query the incoming tx to TSS address ==============
	func() {
		tssAddress := ob.Tss.EVMAddress() // after keygen, ob.Tss.pubkey will be updated
		if tssAddress == (ethcommon.Address{}) {
			ob.logger.ExternalChainWatcher.Warn().Msgf("observeInTx: TSS address not set")
			return
		}

		// query incoming gas asset
		for bn := startBlock; bn <= toBlock; bn++ {
			err = ob.postBlockHeader(toBlock)
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msg("error posting block header")
			}
			block, err := ob.GetBlockByNumberCached(bn)
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msgf("error getting block: %d", bn)
				continue
			}
			headerRLP, err := rlp.EncodeToBytes(block.Header())
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msgf("error encoding block header: %d", bn)
				continue
			}

			_, err = ob.zetaClient.PostAddBlockHeader(
				ob.chain.ChainId,
				block.Hash().Bytes(),
				block.Number().Int64(),
				common.NewEthereumHeader(headerRLP),
			)
			if err != nil {
				ob.logger.ExternalChainWatcher.Error().Err(err).Msgf("error posting block header: %d", bn)
				continue
			}

			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					continue
				}
				if bytes.Equal(tx.Data(), []byte(DonationMessage)) {
					ob.logger.ExternalChainWatcher.Info().Msgf("thank you rich folk for your donation!: %s", tx.Hash().Hex())
					continue
				}

				if *tx.To() == tssAddress {
					receipt, err := ob.evmClient.TransactionReceipt(context.Background(), tx.Hash())
					if err != nil {
						ob.logger.ExternalChainWatcher.Err(err).Msg("TransactionReceipt error")
						continue
					}
					if receipt.Status != 1 { // 1: successful, 0: failed
						ob.logger.ExternalChainWatcher.Info().Msgf("tx %s failed; don't act", tx.Hash())
						continue
					}

					from, err := ob.evmClient.TransactionSender(context.Background(), tx, block.Hash(), receipt.TransactionIndex)
					if err != nil {
						ob.logger.ExternalChainWatcher.Err(err).Msg("TransactionSender error; trying local recovery (assuming LondonSigner dynamic fee tx type) of sender address")
						signer := ethtypes.NewLondonSigner(big.NewInt(ob.chain.ChainId))
						from, err = signer.Sender(tx)
						if err != nil {
							ob.logger.ExternalChainWatcher.Err(err).Msg("local recovery of sender address failed")
							continue
						}
					}
					msg := ob.GetInboundVoteMsgForTokenSentToTSS(tx.Hash(), tx.Value(), receipt, from, tx.Data())
					if msg == nil {
						continue
					}
					zetaHash, err := ob.zetaClient.PostSend(PostSendEVMGasLimit, msg)
					if err != nil {
						ob.logger.ExternalChainWatcher.Error().Err(err).Msg("error posting to zeta core")
						continue
					}
					ob.logger.ExternalChainWatcher.Info().Msgf("Gas Deposit detected and reported: PostSend zeta tx: %s", zetaHash)
				}
			}
		}
	}()
	// ============= end of query the incoming tx to TSS address ==============
	ob.SetLastBlockHeightScanned(toBlock)
	if err := ob.db.Save(clienttypes.ToLastBlockSQLType(ob.GetLastBlockHeightScanned())).Error; err != nil {
		ob.logger.ExternalChainWatcher.Error().Err(err).Msg("error writing toBlock to db")
	}
	return nil
}

func (ob *EVMChainClient) WatchGasPrice() {

	err := ob.PostGasPrice()
	if err != nil {
		height, err := ob.zetaClient.GetBlockHeight()
		if err != nil {
			ob.logger.WatchGasPrice.Error().Err(err).Msg("GetBlockHeight error")
		} else {
			ob.logger.WatchGasPrice.Error().Err(err).Msgf("PostGasPrice error at zeta block : %d  ", height)
		}
	}

	ticker, err := NewDynamicTicker(fmt.Sprintf("EVM_WatchGasPrice_%d", ob.chain.ChainId), ob.GetCoreParams().GasPriceTicker)
	if err != nil {
		ob.logger.WatchGasPrice.Error().Err(err).Msg("NewDynamicTicker error")
		return
	}

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C():
			err = ob.PostGasPrice()
			if err != nil {
				height, err := ob.zetaClient.GetBlockHeight()
				if err != nil {
					ob.logger.WatchGasPrice.Error().Err(err).Msg("GetBlockHeight error")
				} else {
					ob.logger.WatchGasPrice.Error().Err(err).Msgf("PostGasPrice error at zeta block : %d  ", height)
				}
			}
			ticker.UpdateInterval(ob.GetCoreParams().GasPriceTicker, ob.logger.WatchGasPrice)
		case <-ob.stop:
			ob.logger.WatchGasPrice.Info().Msg("WatchGasPrice stopped")
			return
		}
	}
}

func (ob *EVMChainClient) PostGasPrice() error {
	// GAS PRICE
	gasPrice, err := ob.evmClient.SuggestGasPrice(context.TODO())
	if err != nil {
		ob.logger.WatchGasPrice.Err(err).Msg("Err SuggestGasPrice:")
		return err
	}
	blockNum, err := ob.evmClient.BlockNumber(context.TODO())
	if err != nil {
		ob.logger.WatchGasPrice.Err(err).Msg("Err Fetching Most recent Block : ")
		return err
	}

	// SUPPLY
	supply := "100" // lockedAmount on ETH, totalSupply on other chains

	zetaHash, err := ob.zetaClient.PostGasPrice(ob.chain, gasPrice.Uint64(), supply, blockNum)
	if err != nil {
		ob.logger.WatchGasPrice.Err(err).Msg("PostGasPrice to zetacore failed")
		return err
	}
	_ = zetaHash
	//ob.logger.WatchGasPrice.Debug().Msgf("PostGasPrice zeta tx: %s", zetaHash)

	return nil
}

// query ZetaCore about the last block that it has heard from a specific chain.
// return 0 if not existent.
func (ob *EVMChainClient) getLastHeight() (uint64, error) {
	lastheight, err := ob.zetaClient.GetLastBlockHeightByChain(ob.chain)
	if err != nil {
		return 0, errors.Wrap(err, "getLastHeight")
	}
	return lastheight.LastSendHeight, nil
}

func (ob *EVMChainClient) BuildBlockIndex() error {
	logger := ob.logger.ChainLogger.With().Str("module", "BuildBlockIndex").Logger()
	envvar := ob.chain.ChainName.String() + "_SCAN_FROM"
	scanFromBlock := os.Getenv(envvar)
	if scanFromBlock != "" {
		logger.Info().Msgf("envvar %s is set; scan from  block %s", envvar, scanFromBlock)
		if scanFromBlock == clienttypes.EnvVarLatest {
			header, err := ob.evmClient.HeaderByNumber(context.Background(), nil)
			if err != nil {
				return err
			}
			ob.SetLastBlockHeightScanned(header.Number.Uint64())
		} else {
			scanFromBlockInt, err := strconv.ParseUint(scanFromBlock, 10, 64)
			if err != nil {
				return err
			}
			ob.SetLastBlockHeightScanned(scanFromBlockInt)
		}
	} else { // last observed block
		var lastBlockNum clienttypes.LastBlockSQLType
		if err := ob.db.First(&lastBlockNum, clienttypes.LastBlockNumID).Error; err != nil {
			logger.Info().Msg("db PosKey does not exist; read from ZetaCore")
			lastheight, err := ob.getLastHeight()
			if err != nil {
				logger.Warn().Err(err).Msg("getLastHeight error")
			}
			ob.SetLastBlockHeightScanned(lastheight)
			// if ZetaCore does not have last heard block height, then use current
			if ob.GetLastBlockHeightScanned() == 0 {
				header, err := ob.evmClient.HeaderByNumber(context.Background(), nil)
				if err != nil {
					return err
				}
				ob.SetLastBlockHeightScanned(header.Number.Uint64())
			}
			if dbc := ob.db.Save(clienttypes.ToLastBlockSQLType(ob.GetLastBlockHeightScanned())); dbc.Error != nil {
				logger.Error().Err(dbc.Error).Msg("error writing ob.LastBlock to db: ")
			}
		} else {
			ob.SetLastBlockHeightScanned(lastBlockNum.Num)
		}
	}
	return nil
}

func (ob *EVMChainClient) BuildReceiptsMap() error {
	logger := ob.logger
	var receipts []clienttypes.ReceiptSQLType
	if err := ob.db.Find(&receipts).Error; err != nil {
		logger.ChainLogger.Error().Err(err).Msg("error iterating over db")
		return err
	}
	for _, receipt := range receipts {
		r, err := clienttypes.FromReceiptDBType(receipt.Receipt)
		if err != nil {
			return err
		}
		ob.outTXConfirmedReceipts[receipt.Identifier] = r
	}

	return nil
}

func (ob *EVMChainClient) BuildTransactionsMap() error {
	logger := ob.logger
	var transactions []clienttypes.TransactionSQLType
	if err := ob.db.Find(&transactions).Error; err != nil {
		logger.ChainLogger.Error().Err(err).Msg("error iterating over db")
		return err
	}
	for _, transaction := range transactions {
		trans, err := clienttypes.FromTransactionDBType(transaction.Transaction)
		if err != nil {
			return err
		}
		ob.outTXConfirmedTransaction[transaction.Identifier] = trans
	}
	return nil
}

// LoadDB open sql database and load data into EVMChainClient
func (ob *EVMChainClient) LoadDB(dbPath string, chain common.Chain) error {
	if dbPath != "" {
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			err := os.MkdirAll(dbPath, os.ModePerm)
			if err != nil {
				return err
			}
		}
		path := fmt.Sprintf("%s/%s", dbPath, chain.ChainName.String()) //Use "file::memory:?cache=shared" for temp db
		db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		err = db.AutoMigrate(&clienttypes.ReceiptSQLType{},
			&clienttypes.TransactionSQLType{},
			&clienttypes.LastBlockSQLType{})
		if err != nil {
			ob.logger.ChainLogger.Error().Err(err).Msg("error migrating db")
			return err
		}

		ob.db = db
		err = ob.BuildBlockIndex()
		if err != nil {
			return err
		}

		//DISABLING RECEIPT AND TRANSACTION PERSISTENCE
		//err = ob.BuildReceiptsMap()
		//if err != nil {
		//	return err
		//}
		//
		//err = ob.BuildTransactionsMap()
		//if err != nil {
		//	return err
		//}

	}
	return nil
}

func (ob *EVMChainClient) SetMinAndMaxNonce(trackers []types.OutTxTracker) error {
	minNonce, maxNonce := int64(-1), int64(0)
	for _, tracker := range trackers {
		conv := tracker.Nonce
		// #nosec G701 always in range
		intNonce := int64(conv)
		if minNonce == -1 {
			minNonce = intNonce
		}
		if intNonce < minNonce {
			minNonce = intNonce
		}
		if intNonce > maxNonce {
			maxNonce = intNonce
		}
	}
	if minNonce != -1 {
		atomic.StoreInt64(&ob.MinNonce, minNonce)
	}
	if maxNonce > 0 {
		atomic.StoreInt64(&ob.MaxNonce, maxNonce)
	}
	return nil
}

func (ob *EVMChainClient) GetTxID(nonce uint64) string {
	tssAddr := ob.Tss.EVMAddress().String()
	return fmt.Sprintf("%d-%s-%d", ob.chain.ChainId, tssAddr, nonce)
}

func (ob *EVMChainClient) GetBlockByNumberCached(blockNumber uint64) (*ethtypes.Block, error) {
	if block, ok := ob.BlockCache.Get(blockNumber); ok {
		return block.(*ethtypes.Block), nil
	}
	block, err := ob.evmClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, err
	}
	ob.BlockCache.Add(blockNumber, block)
	ob.BlockCache.Add(block.Hash(), block)
	return block, nil
}
