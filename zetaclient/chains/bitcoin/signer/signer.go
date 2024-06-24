package signer

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/zeta-chain/zetacore/pkg/chains"
	"github.com/zeta-chain/zetacore/pkg/coin"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"github.com/zeta-chain/zetacore/zetaclient/chains/base"
	"github.com/zeta-chain/zetacore/zetaclient/chains/bitcoin"
	"github.com/zeta-chain/zetacore/zetaclient/chains/bitcoin/observer"
	"github.com/zeta-chain/zetacore/zetaclient/chains/interfaces"
	"github.com/zeta-chain/zetacore/zetaclient/compliance"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	"github.com/zeta-chain/zetacore/zetaclient/context"
	"github.com/zeta-chain/zetacore/zetaclient/metrics"
	"github.com/zeta-chain/zetacore/zetaclient/outboundprocessor"
)

const (
	// the maximum number of inputs per outbound
	maxNoOfInputsPerTx = 20

	// the rank below (or equal to) which we consolidate UTXOs
	consolidationRank = 10
)

var _ interfaces.ChainSigner = &Signer{}

// Signer deals with signing BTC transactions and implements the ChainSigner interface
type Signer struct {
	// base.Signer implements the base chain signer
	base.Signer

	// client is the RPC client to interact with the Bitcoin chain
	client interfaces.BTCRPCClient
}

// NewSigner creates a new Bitcoin signer
func NewSigner(
	chain chains.Chain,
	zetacoreContext *context.ZetacoreContext,
	tss interfaces.TSSSigner,
	ts *metrics.TelemetryServer,
	logger base.Logger,
	cfg config.BTCConfig) (*Signer, error) {
	// create base signer
	baseSigner := base.NewSigner(chain, zetacoreContext, tss, ts, logger)

	// create the bitcoin rpc client using the provided config
	connCfg := &rpcclient.ConnConfig{
		Host:         cfg.RPCHost,
		User:         cfg.RPCUsername,
		Pass:         cfg.RPCPassword,
		HTTPPostMode: true,
		DisableTLS:   true,
		Params:       cfg.RPCParams,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating bitcoin rpc client: %s", err)
	}

	return &Signer{
		Signer: *baseSigner,
		client: client,
	}, nil
}

// SetZetaConnectorAddress does nothing for BTC
func (signer *Signer) SetZetaConnectorAddress(_ ethcommon.Address) {
}

// SetERC20CustodyAddress does nothing for BTC
func (signer *Signer) SetERC20CustodyAddress(_ ethcommon.Address) {
}

// GetZetaConnectorAddress returns dummy address
func (signer *Signer) GetZetaConnectorAddress() ethcommon.Address {
	return ethcommon.Address{}
}

// GetERC20CustodyAddress returns dummy address
func (signer *Signer) GetERC20CustodyAddress() ethcommon.Address {
	return ethcommon.Address{}
}

// AddWithdrawTxOutputs adds the 3 outputs to the withdraw tx
// 1st output: the nonce-mark btc to TSS itself
// 2nd output: the payment to the recipient
// 3rd output: the remaining btc to TSS itself
func (signer *Signer) AddWithdrawTxOutputs(
	tx *wire.MsgTx,
	to btcutil.Address,
	total float64,
	amount float64,
	nonceMark int64,
	fees *big.Int,
	cancelTx bool,
) error {
	// convert withdraw amount to satoshis
	amountSatoshis, err := bitcoin.GetSatoshis(amount)
	if err != nil {
		return err
	}

	// calculate remaining btc (the change) to TSS self
	remaining := total - amount
	remainingSats, err := bitcoin.GetSatoshis(remaining)
	if err != nil {
		return err
	}
	remainingSats -= fees.Int64()
	remainingSats -= nonceMark
	if remainingSats < 0 {
		return fmt.Errorf("remainder value is negative: %d", remainingSats)
	} else if remainingSats == nonceMark {
		signer.Logger().Std.Info().Msgf("adjust remainder value to avoid duplicate nonce-mark: %d", remainingSats)
		remainingSats--
	}

	// 1st output: the nonce-mark btc to TSS self
	tssAddrP2WPKH := signer.TSS().BTCAddressWitnessPubkeyHash()
	payToSelfScript, err := bitcoin.PayToAddrScript(tssAddrP2WPKH)
	if err != nil {
		return err
	}
	txOut1 := wire.NewTxOut(nonceMark, payToSelfScript)
	tx.AddTxOut(txOut1)

	// 2nd output: the payment to the recipient
	if !cancelTx {
		pkScript, err := bitcoin.PayToAddrScript(to)
		if err != nil {
			return err
		}
		txOut2 := wire.NewTxOut(amountSatoshis, pkScript)
		tx.AddTxOut(txOut2)
	} else {
		// send the amount to TSS self if tx is cancelled
		remainingSats += amountSatoshis
	}

	// 3rd output: the remaining btc to TSS self
	if remainingSats > 0 {
		txOut3 := wire.NewTxOut(remainingSats, payToSelfScript)
		tx.AddTxOut(txOut3)
	}
	return nil
}

// SignWithdrawTx receives utxos sorted by value, amount in BTC, feeRate in BTC per Kb
func (signer *Signer) SignWithdrawTx(
	to btcutil.Address,
	amount float64,
	gasPrice *big.Int,
	sizeLimit uint64,
	observer *observer.Observer,
	height uint64,
	nonce uint64,
	chain chains.Chain,
	cancelTx bool,
) (*wire.MsgTx, error) {
	estimateFee := float64(gasPrice.Uint64()*bitcoin.OutboundBytesMax) / 1e8
	nonceMark := chains.NonceMarkAmount(nonce)

	// refresh unspent UTXOs and continue with keysign regardless of error
	err := observer.FetchUTXOs()
	if err != nil {
		signer.Logger().
			Std.Error().
			Err(err).
			Msgf("SignWithdrawTx: FetchUTXOs error: nonce %d chain %d", nonce, chain.ChainId)
	}

	// select N UTXOs to cover the total expense
	prevOuts, total, consolidatedUtxo, consolidatedValue, err := observer.SelectUTXOs(
		amount+estimateFee+float64(nonceMark)*1e-8,
		maxNoOfInputsPerTx,
		nonce,
		consolidationRank,
		false,
	)
	if err != nil {
		return nil, err
	}

	// build tx with selected unspents
	tx := wire.NewMsgTx(wire.TxVersion)
	for _, prevOut := range prevOuts {
		hash, err := chainhash.NewHashFromStr(prevOut.TxID)
		if err != nil {
			return nil, err
		}
		outpoint := wire.NewOutPoint(hash, prevOut.Vout)
		txIn := wire.NewTxIn(outpoint, nil, nil)
		tx.AddTxIn(txIn)
	}

	// size checking
	// #nosec G701 always positive
	txSize, err := bitcoin.EstimateOutboundSize(uint64(len(prevOuts)), []btcutil.Address{to})
	if err != nil {
		return nil, err
	}
	if sizeLimit < bitcoin.BtcOutboundBytesWithdrawer { // ZRC20 'withdraw' charged less fee from end user
		signer.Logger().Std.Info().
			Msgf("sizeLimit %d is less than BtcOutboundBytesWithdrawer %d for nonce %d", sizeLimit, txSize, nonce)
	}
	if txSize < bitcoin.OutboundBytesMin { // outbound shouldn't be blocked a low sizeLimit
		signer.Logger().Std.Warn().
			Msgf("txSize %d is less than outboundBytesMin %d; use outboundBytesMin", txSize, bitcoin.OutboundBytesMin)
		txSize = bitcoin.OutboundBytesMin
	}
	if txSize > bitcoin.OutboundBytesMax { // in case of accident
		signer.Logger().Std.Warn().
			Msgf("txSize %d is greater than outboundBytesMax %d; use outboundBytesMax", txSize, bitcoin.OutboundBytesMax)
		txSize = bitcoin.OutboundBytesMax
	}

	// fee calculation
	// #nosec G701 always in range (checked above)
	fees := new(big.Int).Mul(big.NewInt(int64(txSize)), gasPrice)
	signer.Logger().
		Std.Info().
		Msgf("bitcoin outbound nonce %d gasPrice %s size %d fees %s consolidated %d utxos of value %v",
			nonce, gasPrice.String(), txSize, fees.String(), consolidatedUtxo, consolidatedValue)

	// add tx outputs
	err = signer.AddWithdrawTxOutputs(tx, to, total, amount, nonceMark, fees, cancelTx)
	if err != nil {
		return nil, err
	}

	// sign the tx
	sigHashes := txscript.NewTxSigHashes(tx)
	witnessHashes := make([][]byte, len(tx.TxIn))
	for ix := range tx.TxIn {
		amt, err := bitcoin.GetSatoshis(prevOuts[ix].Amount)
		if err != nil {
			return nil, err
		}
		pkScript, err := hex.DecodeString(prevOuts[ix].ScriptPubKey)
		if err != nil {
			return nil, err
		}
		witnessHashes[ix], err = txscript.CalcWitnessSigHash(pkScript, sigHashes, txscript.SigHashAll, tx, ix, amt)
		if err != nil {
			return nil, err
		}
	}

	sig65Bs, err := signer.TSS().SignBatch(witnessHashes, height, nonce, chain.ChainId)
	if err != nil {
		return nil, fmt.Errorf("SignBatch error: %v", err)
	}

	for ix := range tx.TxIn {
		sig65B := sig65Bs[ix]
		R := big.NewInt(0).SetBytes(sig65B[:32])
		S := big.NewInt(0).SetBytes(sig65B[32:64])
		sig := btcec.Signature{
			R: R,
			S: S,
		}

		pkCompressed := signer.TSS().PubKeyCompressedBytes()
		hashType := txscript.SigHashAll
		txWitness := wire.TxWitness{append(sig.Serialize(), byte(hashType)), pkCompressed}
		tx.TxIn[ix].Witness = txWitness
	}

	return tx, nil
}

func (signer *Signer) Broadcast(signedTx *wire.MsgTx) error {
	fmt.Printf("BTCSigner: Broadcasting: %s\n", signedTx.TxHash().String())

	var outBuff bytes.Buffer
	err := signedTx.Serialize(&outBuff)
	if err != nil {
		return err
	}
	str := hex.EncodeToString(outBuff.Bytes())
	fmt.Printf("BTCSigner: Transaction Data: %s\n", str)

	hash, err := signer.client.SendRawTransaction(signedTx, true)
	if err != nil {
		return err
	}

	signer.Logger().Std.Info().Msgf("Broadcasting BTC tx , hash %s ", hash)
	return nil
}

func (signer *Signer) TryProcessOutbound(
	cctx *types.CrossChainTx,
	outboundProcessor *outboundprocessor.Processor,
	outboundID string,
	chainObserver interfaces.ChainObserver,
	zetacoreClient interfaces.ZetacoreClient,
	height uint64,
) {
	defer func() {
		outboundProcessor.EndTryProcess(outboundID)
		if err := recover(); err != nil {
			signer.Logger().Std.Error().Msgf("BTC TryProcessOutbound: %s, caught panic error: %v", cctx.Index, err)
		}
	}()

	logger := signer.Logger().Std.With().
		Str("OutboundID", outboundID).
		Str("SendHash", cctx.Index).
		Logger()

	params := cctx.GetCurrentOutboundParam()
	coinType := cctx.InboundParams.CoinType
	if coinType == coin.CoinType_Zeta || coinType == coin.CoinType_ERC20 {
		logger.Error().Msgf("BTC TryProcessOutbound: can only send BTC to a BTC network")
		return
	}

	logger.Info().
		Msgf("BTC TryProcessOutbound: %s, value %d to %s", cctx.Index, params.Amount.BigInt(), params.Receiver)
	btcObserver, ok := chainObserver.(*observer.Observer)
	if !ok {
		logger.Error().Msgf("chain observer is not a bitcoin observer")
		return
	}
	flags := signer.ZetacoreContext().GetCrossChainFlags()
	if !flags.IsOutboundEnabled {
		logger.Info().Msgf("outbound is disabled")
		return
	}
	chain := btcObserver.Chain()
	outboundTssNonce := params.TssNonce
	signerAddress, err := zetacoreClient.GetKeys().GetAddress()
	if err != nil {
		logger.Error().Err(err).Msgf("cannot get signer address")
		return
	}

	// get size limit and gas price
	sizelimit := params.GasLimit
	gasprice, ok := new(big.Int).SetString(params.GasPrice, 10)
	if !ok || gasprice.Cmp(big.NewInt(0)) < 0 {
		logger.Error().Msgf("cannot convert gas price  %s ", params.GasPrice)
		return
	}

	// Check receiver P2WPKH address
	to, err := chains.DecodeBtcAddress(params.Receiver, params.ReceiverChainId)
	if err != nil {
		logger.Error().Err(err).Msgf("cannot decode address %s ", params.Receiver)
		return
	}
	if !chains.IsBtcAddressSupported(to) {
		logger.Error().Msgf("unsupported address %s", params.Receiver)
		return
	}
	amount := float64(params.Amount.Uint64()) / 1e8

	// Add 1 satoshi/byte to gasPrice to avoid minRelayTxFee issue
	networkInfo, err := signer.client.GetNetworkInfo()
	if err != nil {
		logger.Error().Err(err).Msgf("cannot get bitcoin network info")
		return
	}
	satPerByte := bitcoin.FeeRateToSatPerByte(networkInfo.RelayFee)
	gasprice.Add(gasprice, satPerByte)

	// compliance check
	cancelTx := compliance.IsCctxRestricted(cctx)
	if cancelTx {
		compliance.PrintComplianceLog(logger, signer.Logger().Compliance,
			true, chain.ChainId, cctx.Index, cctx.InboundParams.Sender, params.Receiver, "BTC")
		amount = 0.0 // zero out the amount to cancel the tx
	}
	logger.Info().Msgf("SignWithdrawTx: to %s, value %d sats", to.EncodeAddress(), params.Amount.Uint64())

	// sign withdraw tx
	tx, err := signer.SignWithdrawTx(
		to,
		amount,
		gasprice,
		sizelimit,
		btcObserver,
		height,
		outboundTssNonce,
		chain,
		cancelTx,
	)
	if err != nil {
		logger.Warn().Err(err).Msgf("SignOutbound error: nonce %d chain %d", outboundTssNonce, params.ReceiverChainId)
		return
	}
	logger.Info().
		Msgf("Key-sign success: %d => %s, nonce %d", cctx.InboundParams.SenderChainId, chain.ChainName, outboundTssNonce)

	// FIXME: add prometheus metrics
	_, err = zetacoreClient.GetObserverList()
	if err != nil {
		logger.Warn().
			Err(err).
			Msgf("unable to get observer list: chain %d observation %s", outboundTssNonce, observertypes.ObservationType_OutboundTx.String())
	}
	if tx != nil {
		outboundHash := tx.TxHash().String()
		logger.Info().
			Msgf("on chain %s nonce %d, outboundHash %s signer %s", chain.ChainName, outboundTssNonce, outboundHash, signerAddress)
		// TODO: pick a few broadcasters.
		//if len(signers) == 0 || myid == signers[send.OutboundParams.Broadcaster] || myid == signers[int(send.OutboundParams.Broadcaster+1)%len(signers)] {
		// retry loop: 1s, 2s, 4s, 8s, 16s in case of RPC error
		for i := 0; i < 5; i++ {
			// #nosec G404 randomness is not a security issue here
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) //random delay to avoid sychronized broadcast
			err := signer.Broadcast(tx)
			if err != nil {
				logger.Warn().
					Err(err).
					Msgf("broadcasting tx %s to chain %s: nonce %d, retry %d", outboundHash, chain.ChainName, outboundTssNonce, i)
				continue
			}
			logger.Info().
				Msgf("Broadcast success: nonce %d to chain %s outboundHash %s", outboundTssNonce, chain.String(), outboundHash)
			zetaHash, err := zetacoreClient.AddOutboundTracker(
				chain.ChainId,
				outboundTssNonce,
				outboundHash,
				nil,
				"",
				-1,
			)
			if err != nil {
				logger.Err(err).
					Msgf("Unable to add to tracker on zetacore: nonce %d chain %s outboundHash %s", outboundTssNonce, chain.ChainName, outboundHash)
			}
			logger.Info().Msgf("Broadcast to core successful %s", zetaHash)

			// Save successfully broadcasted transaction to btc chain observer
			btcObserver.SaveBroadcastedTx(outboundHash, outboundTssNonce)

			break // successful broadcast; no need to retry
		}
	}
}
