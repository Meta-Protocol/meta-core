package rpc_test

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	solanarpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/node/zetaclient/chains/solana/rpc"
	"github.com/zeta-chain/node/zetaclient/common"
)

// Test_SolanaRPCLive is a phony test to run all live tests
func Test_SolanaRPCLive(t *testing.T) {
	if !common.LiveTestEnabled() {
		return
	}

	LiveTest_GetTransactionWithVersion(t)
	LiveTest_GetFirstSignatureForAddress(t)
	LiveTest_GetSignaturesForAddressUntil(t)
	LiveTest_CheckRPCStatus(t)
}

func LiveTest_GetTransactionWithVersion(t *testing.T) {
	// create a Solana devnet RPC client
	client := solanarpc.New(solanarpc.DevNet_RPC)

	// example transaction of version "0"
	// https://explorer.solana.com/tx/Wqgj7hAaUUSfLzieN912G7GxyGHijzBZgY135NtuFtPRjevK8DnYjWwQZy7LAKFQZu582wsjuab2QP27VMUJzAi?cluster=devnet
	txSig := solana.MustSignatureFromBase58(
		"Wqgj7hAaUUSfLzieN912G7GxyGHijzBZgY135NtuFtPRjevK8DnYjWwQZy7LAKFQZu582wsjuab2QP27VMUJzAi",
	)

	t.Run("should get the transaction if the version is supported", func(t *testing.T) {
		ctx := context.Background()
		txResult, err := rpc.GetTransaction(ctx, client, txSig)
		require.NoError(t, err)
		require.NotNil(t, txResult)
	})

	t.Run("should skip the transaction if the version is not supported", func(t *testing.T) {
		ctx := context.Background()
		txResult, err := rpc.GetTransaction(ctx, client, txSig)
		require.ErrorIs(t, err, rpc.ErrUnsupportedTxVersion)
		require.Nil(t, txResult)
	})
}

func LiveTest_GetFirstSignatureForAddress(t *testing.T) {
	// create a Solana devnet RPC client
	client := solanarpc.New(solanarpc.DevNet_RPC)

	// program address
	address := solana.MustPublicKeyFromBase58("2kJndCL9NBR36ySiQ4bmArs4YgWQu67LmCDfLzk5Gb7s")

	// get the first signature for the address (one by one)
	sig, err := rpc.GetFirstSignatureForAddress(context.Background(), client, address, 1)
	require.NoError(t, err)

	// assert
	actualSig := "2tUQtcrXxtNFtV9kZ4kQsmY7snnEoEEArmu9pUptr4UCy8UdbtjPD6UtfEtPJ2qk5CTzZTmLwsbmZdLymcwSUcHu"
	require.Equal(t, actualSig, sig.String())
}

func LiveTest_GetSignaturesForAddressUntil(t *testing.T) {
	// create a Solana devnet RPC client
	client := solanarpc.New(solanarpc.DevNet_RPC)

	// program address
	address := solana.MustPublicKeyFromBase58("2kJndCL9NBR36ySiQ4bmArs4YgWQu67LmCDfLzk5Gb7s")
	untilSig := solana.MustSignatureFromBase58(
		"2tUQtcrXxtNFtV9kZ4kQsmY7snnEoEEArmu9pUptr4UCy8UdbtjPD6UtfEtPJ2qk5CTzZTmLwsbmZdLymcwSUcHu",
	)

	// get all signatures for the address until the first signature
	sigs, err := rpc.GetSignaturesForAddressUntil(context.Background(), client, address, untilSig, 100)
	require.NoError(t, err)

	// assert
	require.Greater(t, len(sigs), 0)

	// untilSig should not be in the list
	for _, sig := range sigs {
		require.NotEqual(t, untilSig, sig.Signature)
	}
}

func LiveTest_CheckRPCStatus(t *testing.T) {
	// create a Solana devnet RPC client
	client := solanarpc.New(solanarpc.DevNet_RPC)

	// check the RPC status
	ctx := context.Background()
	_, err := rpc.CheckRPCStatus(ctx, client, false)
	require.NoError(t, err)
}
