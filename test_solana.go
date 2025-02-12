package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/near/borsh-go"
)

// AccountInfo represents the Solana account information structure
type AccountInfo struct {
	Pubkey  string `json:"pubkey"`
	Account struct {
		Lamports   uint64   `json:"lamports"`
		Data       []string `json:"data"`
		Owner      string   `json:"owner"`
		Executable bool     `json:"executable"`
		RentEpoch  uint64   `json:"rentEpoch"`
		Space      uint64   `json:"space"`
	} `json:"account"`
}

// PdaInfo represents the PDA account data structure
type PdaInfo struct {
	// Discriminator is the unique identifier for the PDA
	Discriminator [8]byte

	// Nonce is the current nonce for the PDA
	Nonce uint64

	// TssAddress is the TSS address for the PDA
	TssAddress [20]byte

	// Authority is the authority for the PDA
	Authority [32]byte

	// ChainId is the Solana chain id
	ChainID uint64

	DepositPaused bool

	Upgraded bool
}

func main() {
	// Example account info JSON
	accountInfoJSON := `{
  "pubkey": "9dcAyYG4bawApZocwZSyJBi9Mynf5EuKAJfifXdfkqik",
  "account": {
    "lamports": 1447680,
    "data": [
      "qfUAzeEkK14AAAAAAAAAALmVe5VlrKJ8e8vUkioHAoFe0zZEH4B8AhfPMF2NcVsdxF8YiarCWgR8cQzeptF3E04UYwWGAwAAAAAAAAAAAAA=",
      "base64"
    ],
    "owner": "94U5AHQMKkV5txNJ17QPXWoh474PheGou6cNP2FEuL1d",
    "executable": false,
    "rentEpoch": 18446744073709551615,
    "space": 80
  }
}`

	// Parse the account info JSON
	var accountInfo AccountInfo
	err := json.Unmarshal([]byte(accountInfoJSON), &accountInfo)
	if err != nil {
		fmt.Printf("Error parsing account info JSON: %v\n", err)
		return
	}

	// Decode the base64 data
	data, err := base64.StdEncoding.DecodeString(accountInfo.Account.Data[0])
	if err != nil {
		fmt.Printf("Error decoding base64 data: %v\n", err)
		return
	}

	// Create a new PdaInfo instance
	var pdaInfo PdaInfo

	// Deserialize the data using borsh
	err = borsh.Deserialize(&pdaInfo, data)
	if err != nil {
		fmt.Printf("Error deserializing data: %v\n", err)
		return
	}

	// Print the parsed PDA info
	fmt.Printf("Discriminator: %x\n", pdaInfo.Discriminator)
	fmt.Printf("Nonce: %d\n", pdaInfo.Nonce)
	fmt.Printf("TSS Address: %x\n", pdaInfo.TssAddress)
	fmt.Printf("Authority: %x\n", pdaInfo.Authority)
	fmt.Printf("Chain ID: %d\n", pdaInfo.ChainID)
	fmt.Printf("Deposit Paused: %v\n", pdaInfo.DepositPaused)
	fmt.Printf("Upgraded: %v\n", pdaInfo.Upgraded)
}
