package base_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/testutil/sample"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	observertypes "github.com/zeta-chain/node/x/observer/types"
	"github.com/zeta-chain/node/zetaclient/chains/base"
	"github.com/zeta-chain/node/zetaclient/chains/interfaces"
	"github.com/zeta-chain/node/zetaclient/config"
	zctx "github.com/zeta-chain/node/zetaclient/context"
	"github.com/zeta-chain/node/zetaclient/db"
	"github.com/zeta-chain/node/zetaclient/metrics"
	"github.com/zeta-chain/node/zetaclient/testutils/mocks"
)

const (
	// defaultAlertLatency is the default alert latency (in seconds) for unit tests
	defaultAlertLatency = 60

	// defaultConfirmationCount is the default confirmation count for unit tests
	defaultConfirmationCount = 2
)

// createObserver creates a new observer for testing
func createObserver(t *testing.T, chain chains.Chain, alertLatency int64) *base.Observer {
	// constructor parameters
	chainParams := *sample.ChainParams(chain.ChainId)
	chainParams.ConfirmationCount = defaultConfirmationCount
	zetacoreClient := mocks.NewZetacoreClient(t)
	tss := mocks.NewTSS(t)

	database := createDatabase(t)

	// create observer
	logger := base.DefaultLogger()
	ob, err := base.NewObserver(
		chain,
		chainParams,
		zetacoreClient,
		tss,
		base.DefaultBlockCacheSize,
		alertLatency,
		nil,
		database,
		logger,
	)
	require.NoError(t, err)

	return ob
}

func TestNewObserver(t *testing.T) {
	// constructor parameters
	chain := chains.Ethereum
	chainParams := *sample.ChainParams(chain.ChainId)
	appContext := zctx.New(config.New(false), nil, zerolog.Nop())
	zetacoreClient := mocks.NewZetacoreClient(t)
	tss := mocks.NewTSS(t)
	blockCacheSize := base.DefaultBlockCacheSize

	database := createDatabase(t)

	// test cases
	tests := []struct {
		name           string
		chain          chains.Chain
		chainParams    observertypes.ChainParams
		appContext     *zctx.AppContext
		zetacoreClient interfaces.ZetacoreClient
		tss            interfaces.TSSSigner
		blockCacheSize int
		fail           bool
		message        string
	}{
		{
			name:           "should be able to create new observer",
			chain:          chain,
			chainParams:    chainParams,
			appContext:     appContext,
			zetacoreClient: zetacoreClient,
			tss:            tss,
			blockCacheSize: blockCacheSize,
			fail:           false,
		},
		{
			name:           "should return error on invalid block cache size",
			chain:          chain,
			chainParams:    chainParams,
			appContext:     appContext,
			zetacoreClient: zetacoreClient,
			tss:            tss,
			blockCacheSize: 0,
			fail:           true,
			message:        "error creating block cache",
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob, err := base.NewObserver(
				tt.chain,
				tt.chainParams,
				tt.zetacoreClient,
				tt.tss,
				tt.blockCacheSize,
				60,
				nil,
				database,
				base.DefaultLogger(),
			)
			if tt.fail {
				require.ErrorContains(t, err, tt.message)
				require.Nil(t, ob)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, ob)
		})
	}
}

func TestStop(t *testing.T) {
	t.Run("should be able to stop observer", func(t *testing.T) {
		// create observer and initialize db
		ob := createObserver(t, chains.Ethereum, defaultAlertLatency)

		// stop observer
		ob.Stop()
	})
}

func TestObserverGetterAndSetter(t *testing.T) {
	chain := chains.Ethereum

	t.Run("should be able to update chain", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update chain
		newChain := chains.BscMainnet
		ob = ob.WithChain(chains.BscMainnet)
		require.Equal(t, newChain, ob.Chain())
	})

	t.Run("should be able to update zetacore client", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update zetacore client
		newZetacoreClient := mocks.NewZetacoreClient(t)
		ob = ob.WithZetacoreClient(newZetacoreClient)
		require.Equal(t, newZetacoreClient, ob.ZetacoreClient())
	})

	t.Run("should be able to update tss", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update tss
		newTSS := mocks.NewTSS(t)
		ob = ob.WithTSS(newTSS)
		require.Equal(t, newTSS, ob.TSS())
	})

	t.Run("should be able to update last block", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update last block
		newLastBlock := uint64(100)
		ob = ob.WithLastBlock(newLastBlock)
		require.Equal(t, newLastBlock, ob.LastBlock())
	})

	t.Run("should be able to update last block scanned", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update last block scanned
		newLastBlockScanned := uint64(100)
		ob = ob.WithLastBlockScanned(newLastBlockScanned)
		require.Equal(t, newLastBlockScanned, ob.LastBlockScanned())
	})

	t.Run("should be able to update last tx scanned", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update last tx scanned
		newLastTxScanned := sample.EthAddress().String()
		ob = ob.WithLastTxScanned(newLastTxScanned)
		require.Equal(t, newLastTxScanned, ob.LastTxScanned())
	})

	t.Run("should be able to replace block cache", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update block cache
		newBlockCache, err := lru.New(200)
		require.NoError(t, err)

		ob = ob.WithBlockCache(newBlockCache)
		require.Equal(t, newBlockCache, ob.BlockCache())
	})

	t.Run("should be able to update telemetry server", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)

		// update telemetry server
		newServer := metrics.NewTelemetryServer()
		ob = ob.WithTelemetryServer(newServer)
		require.Equal(t, newServer, ob.TelemetryServer())
	})

	t.Run("should be able to get logger", func(t *testing.T) {
		ob := createObserver(t, chain, defaultAlertLatency)
		logger := ob.Logger()

		// should be able to print log
		logger.Chain.Info().Msg("print chain log")
		logger.Inbound.Info().Msg("print inbound log")
		logger.Outbound.Info().Msg("print outbound log")
		logger.GasPrice.Info().Msg("print gasprice log")
		logger.Headers.Info().Msg("print headers log")
		logger.Compliance.Info().Msg("print compliance log")
	})
}

func TestTSSAddressString(t *testing.T) {
	tests := []struct {
		name         string
		chain        chains.Chain
		forceError   bool
		addrExpected string
	}{
		{
			name:         "should return TSS BTC address for Bitcoin chain",
			chain:        chains.BitcoinMainnet,
			addrExpected: "btc",
		},
		{
			name:         "should return TSS EVM address for EVM chain",
			chain:        chains.Ethereum,
			addrExpected: "eth",
		},
		{
			name:         "should return TSS EVM address for other non-BTC chain",
			chain:        chains.SolanaDevnet,
			addrExpected: "eth",
		},
		{
			name:         "should return empty address for unknown BTC chain",
			chain:        chains.BitcoinMainnet,
			forceError:   true,
			addrExpected: "",
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create observer
			ob := createObserver(t, tt.chain, defaultAlertLatency)

			// force error if needed
			if tt.forceError {
				// pause TSS to cause error
				tss := mocks.NewTSS(t)
				tss.Pause()
				ob = ob.WithTSS(tss)
				c := chains.BitcoinRegtest
				c.ChainId = 123123123
				ob.WithChain(c)
			}

			// get TSS address
			addr := ob.TSSAddressString()
			switch tt.addrExpected {
			case "":
				require.Equal(t, "", addr)
			case "btc":
				require.True(t, strings.HasPrefix(addr, "bc"))
			case "eth":
				require.True(t, strings.HasPrefix(addr, "0x"))
			default:
				t.Fail()
			}
		})
	}
}

func TestIsBlockConfirmed(t *testing.T) {
	tests := []struct {
		name      string
		chain     chains.Chain
		block     uint64
		lastBlock uint64
		confirmed bool
	}{
		{
			name:      "should confirm block 100 when confirmation arrives 2",
			chain:     chains.BitcoinMainnet,
			block:     100,
			lastBlock: 101, // got 2 confirmations
			confirmed: true,
		},
		{
			name:      "should not confirm block 100 when confirmation < 2",
			chain:     chains.BitcoinMainnet,
			block:     100,
			lastBlock: 100, // got 1 confirmation, need one more
			confirmed: false,
		},
		{
			name:      "should confirm block 100 when confirmation arrives 2",
			chain:     chains.Ethereum,
			block:     100,
			lastBlock: 99, // last block lagging behind, need to wait
			confirmed: false,
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create observer
			ob := createObserver(t, tt.chain, defaultAlertLatency)
			ob = ob.WithLastBlock(tt.lastBlock)

			// check if block is confirmed
			confirmed := ob.IsBlockConfirmed(tt.block)
			require.Equal(t, tt.confirmed, confirmed)
		})
	}
}

func TestOutboundID(t *testing.T) {
	tests := []struct {
		name  string
		chain chains.Chain
		tss   interfaces.TSSSigner
		nonce uint64
	}{
		{
			name:  "should get correct outbound id for Ethereum chain",
			chain: chains.Ethereum,
			tss:   mocks.NewTSS(t),
			nonce: 100,
		},
		{
			name:  "should get correct outbound id for Bitcoin chain",
			chain: chains.BitcoinMainnet,
			tss:   mocks.NewTSS(t),
			nonce: 200,
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create observer
			ob := createObserver(t, tt.chain, defaultAlertLatency)
			ob = ob.WithTSS(tt.tss)

			// get outbound id
			outboundID := ob.OutboundID(tt.nonce)

			// expected outbound id
			exepctedID := fmt.Sprintf("%d-%s-%d", tt.chain.ChainId, ob.TSSAddressString(), tt.nonce)
			require.Equal(t, exepctedID, outboundID)
		})
	}
}

func TestLoadLastBlockScanned(t *testing.T) {
	chain := chains.Ethereum
	envvar := base.EnvVarLatestBlockByChain(chain)

	t.Run("should be able to load last block scanned", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// create db and write 100 as last block scanned
		err := ob.WriteLastBlockScannedToDB(100)
		require.NoError(t, err)

		// read last block scanned
		err = ob.LoadLastBlockScanned(log.Logger)
		require.NoError(t, err)
		require.EqualValues(t, 100, ob.LastBlockScanned())
	})

	t.Run("latest block scanned should be 0 if not found in db", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// read last block scanned
		err := ob.LoadLastBlockScanned(log.Logger)
		require.NoError(t, err)
		require.EqualValues(t, 0, ob.LastBlockScanned())
	})

	t.Run("should overwrite last block scanned if env var is set", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// create db and write 100 as last block scanned
		ob.WriteLastBlockScannedToDB(100)

		// set env var
		os.Setenv(envvar, "101")

		// read last block scanned
		err := ob.LoadLastBlockScanned(log.Logger)
		require.NoError(t, err)
		require.EqualValues(t, 101, ob.LastBlockScanned())
	})

	t.Run("last block scanned should remain 0 if env var is set to latest", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// create db and write 100 as last block scanned
		ob.WriteLastBlockScannedToDB(100)

		// set env var to 'latest'
		os.Setenv(envvar, base.EnvVarLatestBlock)

		// last block scanned should remain 0
		err := ob.LoadLastBlockScanned(log.Logger)
		require.NoError(t, err)
		require.EqualValues(t, 0, ob.LastBlockScanned())
	})

	t.Run("should return error on invalid env var", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// set invalid env var
		os.Setenv(envvar, "invalid")

		// read last block scanned
		err := ob.LoadLastBlockScanned(log.Logger)
		require.Error(t, err)
	})
}

func TestSaveLastBlockScanned(t *testing.T) {
	t.Run("should be able to save last block scanned", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chains.Ethereum, defaultAlertLatency)

		// save 100 as last block scanned
		err := ob.SaveLastBlockScanned(100)
		require.NoError(t, err)

		// check last block scanned in memory
		require.EqualValues(t, 100, ob.LastBlockScanned())

		// read last block scanned from db
		lastBlockScanned, err := ob.ReadLastBlockScannedFromDB()
		require.NoError(t, err)
		require.EqualValues(t, 100, lastBlockScanned)
	})
}

func TestReadWriteDBLastBlockScanned(t *testing.T) {
	chain := chains.Ethereum
	t.Run("should be able to write and read last block scanned to db", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// write last block scanned
		err := ob.WriteLastBlockScannedToDB(100)
		require.NoError(t, err)

		lastBlockScanned, err := ob.ReadLastBlockScannedFromDB()
		require.NoError(t, err)
		require.EqualValues(t, 100, lastBlockScanned)
	})

	t.Run("should return error when last block scanned not found in db", func(t *testing.T) {
		// create empty db
		ob := createObserver(t, chain, defaultAlertLatency)

		lastScannedBlock, err := ob.ReadLastBlockScannedFromDB()
		require.Error(t, err)
		require.Zero(t, lastScannedBlock)
	})
}
func TestLoadLastTxScanned(t *testing.T) {
	chain := chains.SolanaDevnet
	envvar := base.EnvVarLatestTxByChain(chain)
	lastTx := "5LuQMorgd11p8GWEw6pmyHCDtA26NUyeNFhLWPNk2oBoM9pkag1LzhwGSRos3j4TJLhKjswFhZkGtvSGdLDkmqsk"

	t.Run("should be able to load last tx scanned", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// create db and write sample hash as last tx scanned
		ob.WriteLastTxScannedToDB(lastTx)

		// read last tx scanned
		ob.LoadLastTxScanned()
		require.EqualValues(t, lastTx, ob.LastTxScanned())
	})

	t.Run("latest tx scanned should be empty if not found in db", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// read last tx scanned
		ob.LoadLastTxScanned()
		require.Empty(t, ob.LastTxScanned())
	})

	t.Run("should overwrite last tx scanned if env var is set", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// create db and write sample hash as last tx scanned
		ob.WriteLastTxScannedToDB(lastTx)

		// set env var to other tx
		otherTx := "4Q27KQqJU1gJQavNtkvhH6cGR14fZoBdzqWdWiFd9KPeJxFpYsDRiKAwsQDpKMPtyRhppdncyURTPZyokrFiVHrx"
		os.Setenv(envvar, otherTx)

		// read last block scanned
		ob.LoadLastTxScanned()
		require.EqualValues(t, otherTx, ob.LastTxScanned())
	})
}

func TestSaveLastTxScanned(t *testing.T) {
	chain := chains.SolanaDevnet
	t.Run("should be able to save last tx scanned", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// save random tx hash
		lastSlot := uint64(100)
		lastTx := "5LuQMorgd11p8GWEw6pmyHCDtA26NUyeNFhLWPNk2oBoM9pkag1LzhwGSRos3j4TJLhKjswFhZkGtvSGdLDkmqsk"
		err := ob.SaveLastTxScanned(lastTx, lastSlot)
		require.NoError(t, err)

		// check last tx and slot scanned in memory
		require.EqualValues(t, lastTx, ob.LastTxScanned())
		require.EqualValues(t, lastSlot, ob.LastBlockScanned())

		// read last tx scanned from db
		lastTxScanned, err := ob.ReadLastTxScannedFromDB()
		require.NoError(t, err)
		require.EqualValues(t, lastTx, lastTxScanned)
	})
}

func TestReadWriteDBLastTxScanned(t *testing.T) {
	chain := chains.SolanaDevnet
	t.Run("should be able to write and read last tx scanned to db", func(t *testing.T) {
		// create observer and open db
		ob := createObserver(t, chain, defaultAlertLatency)

		// write last tx scanned
		lastTx := "5LuQMorgd11p8GWEw6pmyHCDtA26NUyeNFhLWPNk2oBoM9pkag1LzhwGSRos3j4TJLhKjswFhZkGtvSGdLDkmqsk"
		err := ob.WriteLastTxScannedToDB(lastTx)
		require.NoError(t, err)

		lastTxScanned, err := ob.ReadLastTxScannedFromDB()
		require.NoError(t, err)
		require.EqualValues(t, lastTx, lastTxScanned)
	})

	t.Run("should return error when last tx scanned not found in db", func(t *testing.T) {
		// create empty db
		ob := createObserver(t, chain, defaultAlertLatency)

		lastTxScanned, err := ob.ReadLastTxScannedFromDB()
		require.Error(t, err)
		require.Empty(t, lastTxScanned)
	})
}

func TestPostVoteInbound(t *testing.T) {
	t.Run("should be able to post vote inbound", func(t *testing.T) {
		// create observer
		ob := createObserver(t, chains.Ethereum, defaultAlertLatency)

		// create mock zetacore client
		zetacoreClient := mocks.NewZetacoreClient(t)
		zetacoreClient.WithPostVoteInbound("", "sampleBallotIndex")
		ob = ob.WithZetacoreClient(zetacoreClient)

		// post vote inbound
		msg := sample.InboundVote(coin.CoinType_Gas, chains.Ethereum.ChainId, chains.ZetaChainMainnet.ChainId)
		ballot, err := ob.PostVoteInbound(context.TODO(), &msg, 100000)
		require.NoError(t, err)
		require.Equal(t, "sampleBallotIndex", ballot)
	})

	t.Run("should not post vote if message basic validation fails", func(t *testing.T) {
		// create observer
		ob := createObserver(t, chains.Ethereum, defaultAlertLatency)

		// create mock zetacore client
		zetacoreClient := mocks.NewZetacoreClient(t)
		ob = ob.WithZetacoreClient(zetacoreClient)

		// create sample message with long Message
		msg := sample.InboundVote(coin.CoinType_Gas, chains.Ethereum.ChainId, chains.ZetaChainMainnet.ChainId)
		msg.Message = strings.Repeat("1", crosschaintypes.MaxMessageLength+1)

		// post vote inbound
		ballot, err := ob.PostVoteInbound(context.TODO(), &msg, 100000)
		require.NoError(t, err)
		require.Empty(t, ballot)
	})
}

func TestAlertOnRPCLatency(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name         string
		blockTime    time.Time
		alertLatency int64
		alerted      bool
	}{
		{
			name:         "should alert on high RPC latency",
			blockTime:    now.Add(-60 * time.Second),
			alertLatency: 55,
			alerted:      true,
		},
		{
			name:         "should not alert on normal RPC latency",
			blockTime:    now.Add(-60 * time.Second),
			alertLatency: 65,
			alerted:      false,
		},
		{
			name:         "should alert on higher RPC latency then default",
			blockTime:    now.Add(-65 * time.Second),
			alertLatency: 0, // 0 means not set
			alerted:      true,
		},
		{
			name:         "should not alert on normal RPC latency when compared to default",
			blockTime:    now.Add(-55 * time.Second),
			alertLatency: 0, // 0 means not set
			alerted:      false,
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create observer
			ob := createObserver(t, chains.Ethereum, tt.alertLatency)

			alerted := ob.AlertOnRPCLatency(tt.blockTime, time.Duration(defaultAlertLatency)*time.Second)
			require.Equal(t, tt.alerted, alerted)
		})
	}
}

func createDatabase(t *testing.T) *db.DB {
	sqlDatabase, err := db.NewFromSqliteInMemory(true)
	require.NoError(t, err)

	return sqlDatabase
}
