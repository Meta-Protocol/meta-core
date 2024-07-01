// Package zetacore provides functionalities for interacting with ZetaChain
package zetacore

import (
	"fmt"
	"sync"
	"time"

	"cosmossdk.io/simapp/params"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/zeta-chain/zetacore/app"
	"github.com/zeta-chain/zetacore/pkg/authz"
	"github.com/zeta-chain/zetacore/pkg/chains"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"github.com/zeta-chain/zetacore/zetaclient/chains/interfaces"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	"github.com/zeta-chain/zetacore/zetaclient/context"
	keyinterfaces "github.com/zeta-chain/zetacore/zetaclient/keys/interfaces"
	"github.com/zeta-chain/zetacore/zetaclient/metrics"
)

var _ interfaces.ZetacoreClient = &Client{}

// Client is the client to send tx to zetacore
type Client struct {
	logger        zerolog.Logger
	blockHeight   int64
	accountNumber map[authz.KeyType]uint64
	seqNumber     map[authz.KeyType]uint64
	grpcConn      *grpc.ClientConn
	cfg           config.ClientConfiguration
	encodingCfg   params.EncodingConfig
	keys          keyinterfaces.ObserverKeys
	broadcastLock *sync.RWMutex
	chainID       string
	chain         chains.Chain
	stop          chan struct{}
	pause         chan struct{}
	Telemetry     *metrics.TelemetryServer

	// enableMockSDKClient is a flag that determines whether the mock cosmos sdk client should be used, primarily for
	// unit testing
	enableMockSDKClient bool
	mockSDKClient       rpcclient.Client
}

// NewClient create a new instance of Client
func NewClient(
	keys keyinterfaces.ObserverKeys,
	chainIP string,
	signerName string,
	chainID string,
	hsmMode bool,
	telemetry *metrics.TelemetryServer,
) (*Client, error) {
	// main module logger
	logger := log.With().Str("module", "ZetacoreClient").Logger()
	cfg := config.ClientConfiguration{
		ChainHost:    fmt.Sprintf("%s:1317", chainIP),
		SignerName:   signerName,
		SignerPasswd: "password",
		ChainRPC:     fmt.Sprintf("%s:26657", chainIP),
		HsmMode:      hsmMode,
	}

	grpcConn, err := grpc.Dial(
		fmt.Sprintf("%s:9090", chainIP),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Error().Err(err).Msg("grpc dial fail")
		return nil, err
	}
	accountsMap := make(map[authz.KeyType]uint64)
	seqMap := make(map[authz.KeyType]uint64)
	for _, keyType := range authz.GetAllKeyTypes() {
		accountsMap[keyType] = 0
		seqMap[keyType] = 0
	}

	zetaChain, err := chains.ZetaChainFromChainID(chainID)
	if err != nil {
		return nil, fmt.Errorf("invalid chain id %s, %w", chainID, err)
	}

	return &Client{
		logger:              logger,
		grpcConn:            grpcConn,
		accountNumber:       accountsMap,
		seqNumber:           seqMap,
		cfg:                 cfg,
		encodingCfg:         app.MakeEncodingConfig(),
		keys:                keys,
		broadcastLock:       &sync.RWMutex{},
		stop:                make(chan struct{}),
		chainID:             chainID,
		chain:               zetaChain,
		pause:               make(chan struct{}),
		Telemetry:           telemetry,
		enableMockSDKClient: false,
		mockSDKClient:       nil,
	}, nil
}

func (c *Client) UpdateChainID(chainID string) error {
	if c.chainID != chainID {
		c.chainID = chainID

		zetaChain, err := chains.ZetaChainFromChainID(chainID)
		if err != nil {
			return fmt.Errorf("invalid chain id %s, %w", chainID, err)
		}
		c.chain = zetaChain
	}

	return nil
}

// Chain returns the Chain chain object
func (c *Client) Chain() chains.Chain {
	return c.chain
}

func (c *Client) GetLogger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) GetKeys() keyinterfaces.ObserverKeys {
	return c.keys
}

func (c *Client) Stop() {
	c.logger.Info().Msgf("zetacore client is stopping")
	close(c.stop) // this notifies all configupdater to stop
}

// GetAccountNumberAndSequenceNumber We do not use multiple KeyType for now , but this can be optionally used in the future to seprate TSS signer from Zetaclient GRantee
func (c *Client) GetAccountNumberAndSequenceNumber(_ authz.KeyType) (uint64, uint64, error) {
	ctx, err := c.GetContext()
	if err != nil {
		return 0, 0, err
	}
	address, err := c.keys.GetAddress()
	if err != nil {
		return 0, 0, err
	}
	return ctx.AccountRetriever.GetAccountNumberSequence(ctx, address)
}

// SetAccountNumber sets the account number and sequence number for the given keyType
func (c *Client) SetAccountNumber(keyType authz.KeyType) error {
	ctx, err := c.GetContext()
	if err != nil {
		return errors.Wrap(err, "fail to get context")
	}
	address, err := c.keys.GetAddress()
	if err != nil {
		return errors.Wrap(err, "fail to get address")
	}
	accN, seq, err := ctx.AccountRetriever.GetAccountNumberSequence(ctx, address)
	if err != nil {
		return errors.Wrap(err, "fail to get account number and sequence number")
	}
	c.accountNumber[keyType] = accN
	c.seqNumber[keyType] = seq

	return nil
}

// WaitForZetacoreToCreateBlocks waits for zetacore to create blocks
func (c *Client) WaitForZetacoreToCreateBlocks() error {
	retryCount := 0
	for {
		block, err := c.GetLatestZetaBlock()
		if err == nil && block.Header.Height > 1 {
			c.logger.Info().Msgf("Zetacore height: %d", block.Header.Height)
			break
		}
		retryCount++
		c.logger.Debug().Msgf("Failed to get latest Block , Retry : %d/%d", retryCount, DefaultRetryCount)
		if retryCount > ExtendedRetryCount {
			return fmt.Errorf("zetacore is not ready, waited for %d seconds", DefaultRetryCount*DefaultRetryInterval)
		}
		time.Sleep(DefaultRetryInterval * time.Second)
	}
	return nil
}

// UpdateZetacoreContext updates zetacore context
// zetacore stores zetacore context for all clients
func (c *Client) UpdateZetacoreContext(coreContext *context.AppContext, init bool, sampledLogger zerolog.Logger) error {
	bn, err := c.GetBlockHeight()
	if err != nil {
		return fmt.Errorf("failed to get zetablock height: %w", err)
	}
	plan, err := c.GetUpgradePlan()
	if err != nil {
		// if there is no active upgrade plan, plan will be nil, err will be nil as well.
		return fmt.Errorf("failed to get upgrade plan: %w", err)
	}
	if plan != nil && bn == plan.Height-1 { // stop zetaclients; notify operator to upgrade and restart
		c.logger.Warn().
			Msgf("Active upgrade plan detected and upgrade height reached: %s at height %d; ZetaClient is stopped;"+
				"please kill this process, replace zetaclientd binary with upgraded version, and restart zetaclientd", plan.Name, plan.Height)
		c.pause <- struct{}{} // notify Orchestrator to stop Observers, Signers, and Orchestrator itself
	}

	chainParams, err := c.GetChainParams()
	if err != nil {
		return fmt.Errorf("failed to get chain params: %w", err)
	}

	newEVMParams := make(map[int64]*observertypes.ChainParams)
	var newBTCParams *observertypes.ChainParams

	// check and update chain params for each chain
	for _, chainParam := range chainParams {
		err := observertypes.ValidateChainParams(chainParam)
		if err != nil {
			sampledLogger.Warn().Err(err).Msgf("Invalid chain params for chain %d", chainParam.ChainId)
			continue
		}
		if chains.IsBitcoinChain(chainParam.ChainId) {
			newBTCParams = chainParam
		} else if chains.IsEVMChain(chainParam.ChainId) {
			newEVMParams[chainParam.ChainId] = chainParam
		}
	}

	supportedChains, err := c.GetSupportedChains()
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}
	newChains := make([]chains.Chain, len(supportedChains))
	for i, chain := range supportedChains {
		newChains[i] = *chain
	}
	keyGen, err := c.GetKeyGen()
	if err != nil {
		c.logger.Info().Msg("Unable to fetch keygen from zetacore")
		return fmt.Errorf("failed to get keygen: %w", err)
	}

	tss, err := c.GetCurrentTss()
	if err != nil {
		c.logger.Info().Err(err).Msg("Unable to fetch TSS from zetacore")
		return fmt.Errorf("failed to get current tss: %w", err)
	}
	tssPubKey := tss.GetTssPubkey()

	crosschainFlags, err := c.GetCrosschainFlags()
	if err != nil {
		c.logger.Info().Msg("Unable to fetch cross-chain flags from zetacore")
		return fmt.Errorf("failed to get crosschain flags: %w", err)
	}

	blockHeaderEnabledChains, err := c.GetBlockHeaderEnabledChains()
	if err != nil {
		c.logger.Info().Msg("Unable to fetch block header enabled chains from zetacore")
		return err
	}

	coreContext.Update(
		keyGen,
		newChains,
		newEVMParams,
		newBTCParams,
		tssPubKey,
		crosschainFlags,
		blockHeaderEnabledChains,
		init,
	)

	return nil
}

// Pause pauses the client
func (c *Client) Pause() {
	<-c.pause
}

// Unpause unpauses the client
func (c *Client) Unpause() {
	c.pause <- struct{}{}
}

// EnableMockSDKClient enables the mock cosmos sdk client
// TODO(revamp): move this to a test package
func (c *Client) EnableMockSDKClient(client rpcclient.Client) {
	c.mockSDKClient = client
	c.enableMockSDKClient = true
}
