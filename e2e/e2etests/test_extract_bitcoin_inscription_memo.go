package e2etests

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	zetabitcoin "github.com/zeta-chain/zetacore/zetaclient/chains/bitcoin"
	btcobserver "github.com/zeta-chain/zetacore/zetaclient/chains/bitcoin/observer"

	"github.com/zeta-chain/zetacore/e2e/runner"
)

func TestExtractBitcoinInscriptionMemo(r *runner.E2ERunner, args []string) {

	r.SetBtcAddress(r.Name, false)

	// obtain some initial fund
	stop := r.MineBlocksIfLocalBitcoin()
	defer stop()
	r.Logger.Info("Mined blocks")

	// list deployer utxos
	utxos, err := r.ListDeployerUTXOs()
	require.NoError(r, err)

	amount := parseFloat(r, args[0])
	memo, err := hex.DecodeString(
		"72f080c854647755d0d9e6f6821f6931f855b9acffd53d87433395672756d58822fd143360762109ab898626556b1c3b8d3096d2361f1297df4a41c1b429471a9aa2fc9be5f27c13b3863d6ac269e4b587d8389f8fd9649859935b0d48dea88cdb40f20c",
	)
	require.NoError(r, err)

	txid, err := r.InscribeToTSSFromDeployerWithMemo(amount, utxos, memo)
	require.NoError(r, err)

	_, err = r.GenerateToAddressIfLocalBitcoin(6, r.BTCDeployerAddress)
	require.NoError(r, err)

	rawtx, err := r.BtcRPCClient.GetRawTransactionVerbose(txid)
	require.NoError(r, err)
	r.Logger.Info("obtained reveal txn id %s", txid)

	dummyCoinbaseTxn := rawtx
	depositorFee := zetabitcoin.DefaultDepositorFee
	events, err := btcobserver.FilterAndParseIncomingTx(
		r.BtcRPCClient,
		[]btcjson.TxRawResult{*dummyCoinbaseTxn, *rawtx},
		0,
		r.BTCTSSAddress.String(),
		log.Logger,
		r.BitcoinParams,
		depositorFee,
	)
	require.NoError(r, err)

	require.Equal(r, 1, len(events))
	event := events[0]

	require.Equal(r, event.MemoBytes, memo)
}
