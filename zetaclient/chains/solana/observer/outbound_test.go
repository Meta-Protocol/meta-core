package observer_test

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/pkg/coin"
	contracts "github.com/zeta-chain/node/pkg/contracts/solana"
	"github.com/zeta-chain/node/testutil/sample"
	"github.com/zeta-chain/node/zetaclient/chains/base"
	"github.com/zeta-chain/node/zetaclient/chains/interfaces"
	"github.com/zeta-chain/node/zetaclient/chains/solana/observer"
	"github.com/zeta-chain/node/zetaclient/db"
	"github.com/zeta-chain/node/zetaclient/testutils"
	"github.com/zeta-chain/node/zetaclient/testutils/mocks"
)

const (
	// gatewayAddressDevnet is the gateway address on devnet for testing
	GatewayAddressTest = "2kJndCL9NBR36ySiQ4bmArs4YgWQu67LmCDfLzk5Gb7s"

	// withdrawTxTest is an archived withdraw tx result on devnet for testing
	// https://explorer.solana.com/tx/5iBYjBYCphzjHKfmPwddMWpV2RNssmzk9Z8NNmV9Rei71pZKBTEVdkmUeyXfn7eWbV8932uSsPfBxgA7UgERNTvq?cluster=devnet
	withdrawTxTest = "5iBYjBYCphzjHKfmPwddMWpV2RNssmzk9Z8NNmV9Rei71pZKBTEVdkmUeyXfn7eWbV8932uSsPfBxgA7UgERNTvq"

	// withdrawFailedTxTest is an archived failed withdraw tx result on devnet for testing
	// https://explorer.solana.com/tx/5nFUQgNSdqTd4aPS4a1xNcbehj19hDzuQLfBqFRj8g7BJdESVY6hFuTFPWFuV6aWAfzEMfVfCdNu9DfzVp5FsHg5?cluster=devnet
	withdrawFailedTxTest = "5nFUQgNSdqTd4aPS4a1xNcbehj19hDzuQLfBqFRj8g7BJdESVY6hFuTFPWFuV6aWAfzEMfVfCdNu9DfzVp5FsHg5"

	// tssAddressTest is the TSS address for testing
	tssAddressTest = "0x05C7dBdd1954D59c9afaB848dA7d8DD3F35e69Cd"
	tssPubKeyTest  = "0x0441707acf75468fd132dfe8a4d48a7726adca036199bbacac7be37e9b7104f2b3b69197bbffa6c7e25ba478ba10505c8929a632e4a84dd03e5e04c260e6c52a00"

	// whitelistTxTest is local devnet tx result for testing
	whitelistTxTest = "phM9bESbiqojmpkkUxgjed8EABkxvPGNau9q31B8Yk1sXUtsxJvd6G9VbZZQPsEyn6RiTH4YBtqJ89omqfbbNNY"

	// withdrawSPLTxTest is local devnet tx result for testing
	withdrawSPLTxTest = "3NgoR4K9FJq7UunorPRGW9wpqMV8oNvZERejutd7bKmqh3CKEV5DMZndhZn7hQ1i4RhTyHXRWxtR5ZNVHmmjAUSF"
)

// createTestObserver creates a test observer for testing
func createTestObserver(
	t *testing.T,
	chain chains.Chain,
	solClient interfaces.SolanaRPCClient,
	tss interfaces.TSSSigner,
) *observer.Observer {
	database, err := db.NewFromSqliteInMemory(true)
	require.NoError(t, err)

	testLogger := zerolog.New(zerolog.NewTestWriter(t))
	logger := base.Logger{Std: testLogger, Compliance: testLogger}

	// create observer
	chainParams := sample.ChainParams(chain.ChainId)
	chainParams.GatewayAddress = GatewayAddressTest
	ob, err := observer.NewObserver(chain, solClient, *chainParams, nil, tss, 60, database, logger, nil)
	require.NoError(t, err)

	return ob
}

func Test_CheckFinalizedTx(t *testing.T) {
	// ARRANGE
	// the test chain and transaction hash
	chain := chains.SolanaDevnet
	txHash := withdrawTxTest
	txHashFailed := withdrawFailedTxTest
	txSig := solana.MustSignatureFromBase58(txHash)
	coinType := coin.CoinType_Gas
	nonce := uint64(0)

	// load archived outbound tx result
	txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)

	// mock GetTransaction result
	solClient := mocks.NewSolanaRPCClient(t)
	solClient.On("GetTransaction", mock.Anything, txSig, mock.Anything).Return(txResult, nil)

	// mock TSS
	tss := mocks.NewTSS(t).FakePubKey(tssPubKeyTest)

	// create observer with and TSS
	ob := createTestObserver(t, chain, solClient, tss)
	ctx := context.Background()

	t.Run("should successfully check finalized tx", func(t *testing.T) {
		// ACT
		tx, finalized := ob.CheckFinalizedTx(ctx, txHash, nonce, coinType)

		// ASSERT
		require.True(t, finalized)
		require.NotNil(t, tx)
	})

	t.Run("should return error on invalid tx hash", func(t *testing.T) {
		// ACT
		tx, finalized := ob.CheckFinalizedTx(ctx, "invalid_hash_1234", nonce, coinType)

		// ASSERT
		require.False(t, finalized)
		require.Nil(t, tx)
	})

	t.Run("should return error on GetTransaction error", func(t *testing.T) {
		// ARRANGE
		// mock GetTransaction error
		client := mocks.NewSolanaRPCClient(t)
		client.On("GetTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		// create observer
		ob := createTestObserver(t, chain, client, tss)

		// ACT
		tx, finalized := ob.CheckFinalizedTx(ctx, txHash, nonce, coinType)

		// ASSERT
		require.False(t, finalized)
		require.Nil(t, tx)
	})

	t.Run("should return error on if transaction is failed", func(t *testing.T) {
		// ARRANGE
		// load archived outbound tx result which is failed due to nonce mismatch
		failedResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHashFailed)

		// mock GetTransaction result with failed status
		client := mocks.NewSolanaRPCClient(t)
		client.On("GetTransaction", mock.Anything, txSig, mock.Anything).Return(failedResult, nil)

		// create observer
		ob := createTestObserver(t, chain, client, tss)

		// ACT
		tx, finalized := ob.CheckFinalizedTx(ctx, txHash, nonce, coinType)

		// ASSERT
		require.False(t, finalized)
		require.Nil(t, tx)
	})

	t.Run("should return error on ParseGatewayInstruction error", func(t *testing.T) {
		// ACT
		// use CoinType_Zeta to cause ParseGatewayInstruction error
		tx, finalized := ob.CheckFinalizedTx(ctx, txHash, nonce, coin.CoinType_Zeta)

		// ASSERT
		require.False(t, finalized)
		require.Nil(t, tx)
	})

	t.Run("should return error on ECDSA signer mismatch", func(t *testing.T) {
		// ARRANGE
		// create observer with other TSS address
		tssOther := mocks.NewTSS(t)
		ob := createTestObserver(t, chain, solClient, tssOther)

		// ACT
		tx, finalized := ob.CheckFinalizedTx(ctx, txHash, nonce, coinType)

		// ASSERT
		require.False(t, finalized)
		require.Nil(t, tx)
	})

	t.Run("should return error on nonce mismatch", func(t *testing.T) {
		// ACT
		// use different nonce
		tx, finalized := ob.CheckFinalizedTx(ctx, txHash, nonce+1, coinType)

		// ASSERT
		require.False(t, finalized)
		require.Nil(t, tx)
	})
}

func Test_ParseGatewayInstruction(t *testing.T) {
	// the test chain and transaction hash
	chain := chains.SolanaDevnet
	txHash := withdrawTxTest
	txAmount := uint64(890880)

	// gateway address
	gatewayID, err := solana.PublicKeyFromBase58(GatewayAddressTest)
	require.NoError(t, err)

	t.Run("should parse gateway instruction", func(t *testing.T) {
		// ARRANGE
		// load archived outbound tx result
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)

		// ACT
		// parse gateway instruction
		inst, err := observer.ParseGatewayInstruction(txResult, gatewayID, coin.CoinType_Gas)
		require.NoError(t, err)

		// ASSERT
		// check sender, nonce and amount
		sender, err := inst.Signer()
		require.NoError(t, err)
		require.Equal(t, tssAddressTest, sender.String())
		require.EqualValues(t, inst.GatewayNonce(), 0)
		require.EqualValues(t, inst.TokenAmount(), txAmount)
	})

	t.Run("should return error on invalid number of instructions", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// remove all instructions
		tx.Message.Instructions = nil

		// ACT
		inst, err := observer.ParseGatewayInstruction(txResult, gatewayID, coin.CoinType_Gas)

		// ASSERT
		require.ErrorContains(t, err, "want 1 instruction, got 0")
		require.Nil(t, inst)
	})

	t.Run("should return error on invalid program id index", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// set invalid program id index (out of range)
		tx.Message.Instructions[0].ProgramIDIndex = 4

		// ACT
		inst, err := observer.ParseGatewayInstruction(txResult, gatewayID, coin.CoinType_Gas)

		// ASSERT
		require.ErrorContains(t, err, "error getting program ID")
		require.Nil(t, inst)
	})

	t.Run("should return error when invoked program is not gateway", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// set invalid program id index (pda account index)
		tx.Message.Instructions[0].ProgramIDIndex = 1

		// ACT
		inst, err := observer.ParseGatewayInstruction(txResult, gatewayID, coin.CoinType_Gas)

		// ASSERT
		require.ErrorContains(t, err, "not matching gatewayID")
		require.Nil(t, inst)
	})

	t.Run("should return error when instruction parsing fails", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// set invalid instruction data to cause parsing error
		tx.Message.Instructions[0].Data = []byte("invalid instruction data")

		// ACT
		inst, err := observer.ParseGatewayInstruction(txResult, gatewayID, coin.CoinType_Gas)

		// ASSERT
		require.Error(t, err)
		require.Nil(t, inst)
	})

	t.Run("should return error on unsupported coin type", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)

		// ACT
		inst, err := observer.ParseGatewayInstruction(txResult, gatewayID, coin.CoinType_Zeta)

		// ASSERT
		require.ErrorContains(t, err, "unsupported outbound coin type")
		require.Nil(t, inst)
	})
}

func Test_ParseInstructionWithdraw(t *testing.T) {
	// the test chain and transaction hash
	chain := chains.SolanaDevnet
	txHash := withdrawTxTest
	txAmount := uint64(890880)

	t.Run("should parse instruction withdraw", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		instruction := tx.Message.Instructions[0]

		// ACT
		inst, err := contracts.ParseInstructionWithdraw(instruction)
		require.NoError(t, err)

		// ASSERT
		// check sender, nonce and amount
		sender, err := inst.Signer()
		require.NoError(t, err)
		require.Equal(t, tssAddressTest, sender.String())
		require.EqualValues(t, inst.GatewayNonce(), 0)
		require.EqualValues(t, inst.TokenAmount(), txAmount)
	})

	t.Run("should return error on invalid instruction data", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		txFake, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// set invalid instruction data
		instruction := txFake.Message.Instructions[0]
		instruction.Data = []byte("invalid instruction data")

		// ACT
		inst, err := contracts.ParseInstructionWithdraw(instruction)

		// ASSERT
		require.ErrorContains(t, err, "error deserializing instruction")
		require.Nil(t, inst)
	})

	t.Run("should return error on discriminator mismatch", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		txFake, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// overwrite discriminator (first 8 bytes)
		instruction := txFake.Message.Instructions[0]
		fakeDiscriminator := "b712469c946da12100980d0000000000"
		fakeDiscriminatorBytes, err := hex.DecodeString(fakeDiscriminator)
		require.NoError(t, err)
		copy(instruction.Data, fakeDiscriminatorBytes)

		// ACT
		inst, err := contracts.ParseInstructionWithdraw(instruction)

		// ASSERT
		require.ErrorContains(t, err, "not a withdraw instruction")
		require.Nil(t, inst)
	})
}

func Test_ParseInstructionWhitelist(t *testing.T) {
	// the test chain and transaction hash
	chain := chains.SolanaDevnet
	txHash := whitelistTxTest
	txAmount := uint64(0)

	t.Run("should parse instruction whitelist", func(t *testing.T) {
		// ARRANGE
		// tss address used in local devnet
		tssAddress := "0x7E8c7bAcd3c6220DDC35A4EA1141BE14F2e1dFEB"
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		instruction := tx.Message.Instructions[0]

		// ACT
		inst, err := contracts.ParseInstructionWhitelist(instruction)
		require.NoError(t, err)

		// ASSERT
		// check sender, nonce and amount
		sender, err := inst.Signer()
		require.NoError(t, err)
		require.Equal(t, tssAddress, sender.String())
		require.EqualValues(t, inst.GatewayNonce(), 3)
		require.EqualValues(t, inst.TokenAmount(), txAmount)
	})

	t.Run("should return error on invalid instruction data", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		txFake, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// set invalid instruction data
		instruction := txFake.Message.Instructions[0]
		instruction.Data = []byte("invalid instruction data")

		// ACT
		inst, err := contracts.ParseInstructionWhitelist(instruction)

		// ASSERT
		require.ErrorContains(t, err, "error deserializing instruction")
		require.Nil(t, inst)
	})

	t.Run("should return error on discriminator mismatch", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		txFake, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// overwrite discriminator (first 8 bytes)
		instruction := txFake.Message.Instructions[0]
		fakeDiscriminator := "b712469c946da12100980d0000000000"
		fakeDiscriminatorBytes, err := hex.DecodeString(fakeDiscriminator)
		require.NoError(t, err)
		copy(instruction.Data, fakeDiscriminatorBytes)

		// ACT
		inst, err := contracts.ParseInstructionWhitelist(instruction)

		// ASSERT
		require.ErrorContains(t, err, "not a whitelist_spl_mint instruction")
		require.Nil(t, inst)
	})
}

func Test_ParseInstructionWithdrawSPL(t *testing.T) {
	// the test chain and transaction hash
	chain := chains.SolanaDevnet
	txHash := withdrawSPLTxTest
	txAmount := uint64(1000000)

	t.Run("should parse instruction withdraw spl", func(t *testing.T) {
		// ARRANGE
		// tss address used in local devnet
		tssAddress := "0x9c427Bc95cC11dE0D3Fb7603A99833e8f781Cfba"
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		tx, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		instruction := tx.Message.Instructions[0]

		// ACT
		inst, err := contracts.ParseInstructionWithdrawSPL(instruction)
		require.NoError(t, err)

		// ASSERT
		// check sender, nonce and amount
		sender, err := inst.Signer()
		require.NoError(t, err)
		require.Equal(t, tssAddress, sender.String())
		require.EqualValues(t, 3, inst.GatewayNonce())
		require.EqualValues(t, txAmount, inst.TokenAmount())
		require.EqualValues(t, 6, inst.Decimals)
		require.EqualValues(t, contracts.DiscriminatorWithdrawSPL, inst.Discriminator)
	})

	t.Run("should return error on invalid instruction data", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		txFake, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// set invalid instruction data
		instruction := txFake.Message.Instructions[0]
		instruction.Data = []byte("invalid instruction data")

		// ACT
		inst, err := contracts.ParseInstructionWithdrawSPL(instruction)

		// ASSERT
		require.ErrorContains(t, err, "error deserializing instruction")
		require.Nil(t, inst)
	})

	t.Run("should return error on discriminator mismatch", func(t *testing.T) {
		// ARRANGE
		// load and unmarshal archived transaction
		txResult := testutils.LoadSolanaOutboundTxResult(t, TestDataDir, chain.ChainId, txHash)
		txFake, err := txResult.Transaction.GetTransaction()
		require.NoError(t, err)

		// overwrite discriminator (first 8 bytes)
		instruction := txFake.Message.Instructions[0]
		fakeDiscriminator := "b712469c946da12100980d0000000000"
		fakeDiscriminatorBytes, err := hex.DecodeString(fakeDiscriminator)
		require.NoError(t, err)
		copy(instruction.Data, fakeDiscriminatorBytes)

		// ACT
		inst, err := contracts.ParseInstructionWithdrawSPL(instruction)

		// ASSERT
		require.ErrorContains(t, err, "not a withdraw_spl_token instruction")
		require.Nil(t, inst)
	})
}
