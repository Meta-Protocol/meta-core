package observer

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tonkeeper/tongo/tlb"
	"github.com/tonkeeper/tongo/ton"
	toncontracts "github.com/zeta-chain/node/pkg/contracts/ton"
	"github.com/zeta-chain/node/testutil/sample"
	"github.com/zeta-chain/node/zetaclient/chains/ton/liteapi"
)

func TestInbound(t *testing.T) {
	gw := toncontracts.NewGateway(
		ton.MustParseAccountID("0:997d889c815aeac21c47f86ae0e38383efc3c3463067582f6263ad48c5a1485b"),
	)

	t.Run("Ensure last scanned tx", func(t *testing.T) {
		t.Run("Unable to get first tx", func(t *testing.T) {
			// ARRANGE
			ts := newTestSuite(t)

			// Given observer
			ob, err := New(ts.baseObserver, ts.liteClient, gw)
			require.NoError(t, err)

			// Given mocked lite client call
			ts.OnGetFirstTransaction(gw.AccountID(), nil, 0, errors.New("oops")).Once()

			// ACT
			// Observe inbounds once
			err = ob.observeInbound(ts.ctx)

			// ASSERT
			assert.ErrorContains(t, err, "unable to ensure last scanned tx")
			assert.Empty(t, ob.LastTxScanned())
		})

		t.Run("All good", func(t *testing.T) {
			// ARRANGE
			ts := newTestSuite(t)

			// Given mocked lite client calls
			firstTX := sample.TONDonation(t, gw.AccountID(), toncontracts.Donation{
				Sender: sample.GenerateTONAccountID(),
				Amount: tonCoins(t, "1"),
			})

			ts.OnGetFirstTransaction(gw.AccountID(), &firstTX, 0, nil).Once()
			ts.OnGetTransactionsUntil(gw.AccountID(), firstTX.Lt, txHash(firstTX), nil, nil).Once()

			// Given observer
			ob, err := New(ts.baseObserver, ts.liteClient, gw)
			require.NoError(t, err)

			// ACT
			// Observe inbounds once
			err = ob.observeInbound(ts.ctx)

			// ASSERT
			assert.NoError(t, err)

			// Check that last scanned tx is set and is valid
			lastScanned, err := ob.ReadLastTxScannedFromDB()
			assert.NoError(t, err)
			assert.Equal(t, ob.LastTxScanned(), lastScanned)

			lt, hash, err := liteapi.TransactionHashFromString(lastScanned)
			assert.NoError(t, err)
			assert.Equal(t, firstTX.Lt, lt)
			assert.Equal(t, firstTX.Hash().Hex(), hash.Hex())
		})
	})

	t.Run("Donation", func(t *testing.T) {
		// ARRANGE
		ts := newTestSuite(t)

		// Given observer
		ob, err := New(ts.baseObserver, ts.liteClient, gw)
		require.NoError(t, err)

		lastScanned := ts.SetupLastScannedTX(gw.AccountID())

		// Given mocked lite client calls
		donation := sample.TONDonation(t, gw.AccountID(), toncontracts.Donation{
			Sender: sample.GenerateTONAccountID(),
			Amount: tonCoins(t, "12"),
		})

		txs := []ton.Transaction{donation}

		ts.
			OnGetTransactionsUntil(gw.AccountID(), lastScanned.Lt, txHash(lastScanned), txs, nil).
			Once()

		// ACT
		// Observe inbounds once
		err = ob.observeInbound(ts.ctx)

		// ASSERT
		assert.NoError(t, err)

		// nothing happened, but tx scanned
		lt, hash, err := liteapi.TransactionHashFromString(ob.LastTxScanned())
		assert.NoError(t, err)
		assert.Equal(t, donation.Lt, lt)
		assert.Equal(t, donation.Hash().Hex(), hash.Hex())
	})

	t.Run("Deposit", func(t *testing.T) {
		// ARRANGE
		ts := newTestSuite(t)

		// Given observer
		ob, err := New(ts.baseObserver, ts.liteClient, gw)
		require.NoError(t, err)

		lastScanned := ts.SetupLastScannedTX(gw.AccountID())

		// Given mocked lite client calls
		deposit := toncontracts.Deposit{
			Sender:    sample.GenerateTONAccountID(),
			Amount:    tonCoins(t, "12"),
			Recipient: sample.EthAddress(),
		}

		depositTX := sample.TONDeposit(t, gw.AccountID(), deposit)
		txs := []ton.Transaction{depositTX}

		ts.
			OnGetTransactionsUntil(gw.AccountID(), lastScanned.Lt, txHash(lastScanned), txs, nil).
			Once()

		ts.MockGetBlockHeader(depositTX.BlockID)

		// ACT
		// Observe inbounds once
		err = ob.observeInbound(ts.ctx)

		// ASSERT
		assert.NoError(t, err)

		// Check that cctx was sent to zetacore
		require.Len(t, ts.votesBag, 1)

		// Check CCTX
		cctx := ts.votesBag[0]

		assert.NotNil(t, cctx)

		assert.Equal(t, deposit.Sender.ToRaw(), cctx.Sender)
		assert.Equal(t, ts.chain.ChainId, cctx.SenderChainId)

		assert.Equal(t, "", cctx.Asset)
		assert.Equal(t, deposit.Amount.Uint64(), cctx.Amount.Uint64())
		assert.Equal(t, string(deposit.Recipient.Bytes()), cctx.Message)

		// Check hash & block height
		expectedHash := liteapi.TransactionHashToString(depositTX.Lt, txHash(depositTX))
		assert.Equal(t, expectedHash, cctx.InboundHash)

		blockInfo, err := ts.liteClient.GetBlockHeader(ts.ctx, depositTX.BlockID, 0)
		require.NoError(t, err)

		assert.Equal(t, uint64(blockInfo.MinRefMcSeqno), cctx.InboundBlockHeight)
	})

	t.Run("Deposit and call", func(t *testing.T) {
		// ARRANGE
		ts := newTestSuite(t)

		// Given observer
		ob, err := New(ts.baseObserver, ts.liteClient, gw)
		require.NoError(t, err)

		lastScanned := ts.SetupLastScannedTX(gw.AccountID())

		// Given mocked lite client calls
		const callData = "hey there"
		depositAndCall := toncontracts.DepositAndCall{
			Deposit: toncontracts.Deposit{
				Sender:    sample.GenerateTONAccountID(),
				Amount:    tonCoins(t, "4"),
				Recipient: sample.EthAddress(),
			},
			CallData: []byte(callData),
		}

		depositAndCallTX := sample.TONDepositAndCall(t, gw.AccountID(), depositAndCall)
		txs := []ton.Transaction{depositAndCallTX}

		ts.
			OnGetTransactionsUntil(gw.AccountID(), lastScanned.Lt, txHash(lastScanned), txs, nil).
			Once()

		ts.MockGetBlockHeader(depositAndCallTX.BlockID)

		// ACT
		// Observe inbounds once
		err = ob.observeInbound(ts.ctx)

		// ASSERT
		assert.NoError(t, err)

		// Check that cctx was sent to zetacore
		require.Len(t, ts.votesBag, 1)

		// Check CCTX
		cctx := ts.votesBag[0]

		assert.NotNil(t, cctx)

		assert.Equal(t, depositAndCall.Sender.ToRaw(), cctx.Sender)
		assert.Equal(t, ts.chain.ChainId, cctx.SenderChainId)

		assert.Equal(t, "", cctx.Asset)
		assert.Equal(t, depositAndCall.Amount.Uint64(), cctx.Amount.Uint64())
		assert.Equal(t, string(depositAndCall.Recipient.Bytes())+callData, cctx.Message)

		// Check hash & block height
		expectedHash := liteapi.TransactionHashToString(depositAndCallTX.Lt, txHash(depositAndCallTX))
		assert.Equal(t, expectedHash, cctx.InboundHash)

		blockInfo, err := ts.liteClient.GetBlockHeader(ts.ctx, depositAndCallTX.BlockID, 0)
		require.NoError(t, err)

		assert.Equal(t, uint64(blockInfo.MinRefMcSeqno), cctx.InboundBlockHeight)
	})

	t.Run("Multiple transactions", func(t *testing.T) {
		// ARRANGE
		ts := newTestSuite(t)

		// Given observer
		ob, err := New(ts.baseObserver, ts.liteClient, gw)
		require.NoError(t, err)

		lastScanned := ts.SetupLastScannedTX(gw.AccountID())

		// Given several transactions
		txs := []ton.Transaction{
			// should be skipped
			sample.TONDonation(t, gw.AccountID(), toncontracts.Donation{
				Sender: sample.GenerateTONAccountID(),
				Amount: tonCoins(t, "1"),
			}),
			// should be voted
			sample.TONDeposit(t, gw.AccountID(), toncontracts.Deposit{
				Sender:    sample.GenerateTONAccountID(),
				Amount:    tonCoins(t, "3"),
				Recipient: sample.EthAddress(),
			}),
			// should be skipped (invalid inbound message)
			sample.TONTransaction(t, sample.TONTransactionProps{
				Account: gw.AccountID(),
				Input:   &tlb.Message{},
			}),
			// should be voted
			sample.TONDeposit(t, gw.AccountID(), toncontracts.Deposit{
				Sender:    sample.GenerateTONAccountID(),
				Amount:    tonCoins(t, "3"),
				Recipient: sample.EthAddress(),
			}),
			// should be skipped (invalid inbound/outbound messages)
			sample.TONTransaction(t, sample.TONTransactionProps{
				Account: gw.AccountID(),
				Input:   &tlb.Message{},
				Output:  &tlb.Message{},
			}),
		}

		ts.
			OnGetTransactionsUntil(gw.AccountID(), lastScanned.Lt, txHash(lastScanned), txs, nil).
			Once()

		for _, tx := range txs {
			ts.MockGetBlockHeader(tx.BlockID)
		}

		// ACT
		// Observe inbounds once
		err = ob.observeInbound(ts.ctx)

		// ASSERT
		assert.NoError(t, err)

		// Check that cctx was sent to zetacore
		assert.Equal(t, 2, len(ts.votesBag))

		var (
			hash1 = liteapi.TransactionHashToString(txs[1].Lt, txHash(txs[1]))
			hash2 = liteapi.TransactionHashToString(txs[3].Lt, txHash(txs[3]))
		)

		assert.Equal(t, hash1, ts.votesBag[0].InboundHash)
		assert.Equal(t, hash2, ts.votesBag[1].InboundHash)

		// Check that last scanned tx points to the last tx in a list (even if it was skipped)
		var (
			lastTX          = txs[len(txs)-1]
			lastScannedHash = ob.LastTxScanned()
		)

		lastLT, lastHash, err := liteapi.TransactionHashFromString(lastScannedHash)
		assert.NoError(t, err)
		assert.Equal(t, lastTX.Lt, lastLT)
		assert.Equal(t, lastTX.Hash().Hex(), lastHash.Hex())
	})
}

func txHash(tx ton.Transaction) ton.Bits256 {
	return ton.Bits256(tx.Hash())
}
