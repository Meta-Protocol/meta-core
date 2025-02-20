package sui

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/hex"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"
)

// Digest calculates tx digest (hash) for further signing by TSS.
func Digest(tx models.TxnMetaData) ([32]byte, error) {
	txBytes, err := base64.StdEncoding.DecodeString(tx.TxBytes)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "failed to decode tx bytes")
	}

	message := messageWithIntentPrefix(txBytes)

	// "When invoking the signing API, you must first hash the intent message of the tx
	// data to 32 bytes using Blake2b ... For ECDSA Secp256k1 and Secp256r1,
	// you must use SHA256 as the internal hash function"
	// https://docs.sui.io/concepts/cryptography/transaction-auth/signatures#signature-requirements
	return blake2b.Sum256(message), nil
}

// https://github.com/MystenLabs/sui/blob/0dc1a38f800fc2d8fabe11477fdef702058cf00d/crates/sui-types/src/intent.rs
// #1 = IntentScope(transactionData=0)
// #2 = Version(0)
// #3 = AppId(Sui=0)
var defaultIntent = []byte{0, 0, 0}

// Constructs binary message with intent prefix.
// https://docs.sui.io/concepts/cryptography/transaction-auth/intent-signing#structs
func messageWithIntentPrefix(message []byte) []byte {
	glued := make([]byte, len(defaultIntent)+len(message))
	copy(glued, defaultIntent)
	copy(glued[len(defaultIntent):], message)

	return glued
}

// AddressFromPubKeyECDSA converts ECDSA public key to Sui address.
// https://docs.sui.io/concepts/cryptography/transaction-auth/keys-addresses
// https://docs.sui.io/concepts/cryptography/transaction-auth/signatures
func AddressFromPubKeyECDSA(pk *ecdsa.PublicKey) string {
	const flagSecp256k1 = 0x01

	pubBytes := elliptic.MarshalCompressed(pk.Curve, pk.X, pk.Y)

	raw := make([]byte, 1+len(pubBytes))
	raw[0] = flagSecp256k1
	copy(raw[1:], pubBytes)

	addrBytes := blake2b.Sum256(raw)

	return "0x" + hex.EncodeToString(addrBytes[:])
}
