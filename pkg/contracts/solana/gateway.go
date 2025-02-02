// Package solana provides structures and constants that are used when interacting with the gateway program on Solana chain.
package solana

import (
	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"
	idlgateway "github.com/zeta-chain/protocol-contracts-solana/go-idl/generated"
)

const (
	// PDASeed is the seed for the Solana gateway program derived address
	PDASeed = "meta"

	// AccountsNumberOfDeposit is the number of accounts required for Solana gateway deposit instruction
	// [signer, pda, system_program]
	accountsNumDeposit = 3

	// AccountsNumberOfDeposit is the number of accounts required for Solana gateway deposit spl instruction
	// [signer, pda, whitelist_entry, mint_account, token_program, from, to, system_program]
	accountsNumberDepositSPL = 8
)

var (
	// DiscriminatorInitialize returns the discriminator for Solana gateway 'initialize' instruction
	DiscriminatorInitialize = idlgateway.IDLGateway.GetDiscriminator("initialize")

	// DiscriminatorDeposit returns the discriminator for Solana gateway 'deposit' instruction
	DiscriminatorDeposit = idlgateway.IDLGateway.GetDiscriminator("deposit")

	// DiscriminatorDeposit returns the discriminator for Solana gateway 'deposit_and_call' instruction
	DiscriminatorDepositAndCall = idlgateway.IDLGateway.GetDiscriminator("deposit_and_call")

	// DiscriminatorDepositSPL returns the discriminator for Solana gateway 'deposit_spl_token' instruction
	DiscriminatorDepositSPL = idlgateway.IDLGateway.GetDiscriminator("deposit_spl_token")

	// DiscriminatorDepositSPLAndCall returns the discriminator for Solana gateway 'deposit_spl_token_and_call' instruction
	DiscriminatorDepositSPLAndCall = idlgateway.IDLGateway.GetDiscriminator("deposit_spl_token_and_call")

	// DiscriminatorWithdraw returns the discriminator for Solana gateway 'withdraw' instruction
	DiscriminatorWithdraw = idlgateway.IDLGateway.GetDiscriminator("withdraw")

	// DiscriminatorExecute returns the discriminator for Solana gateway 'execute' instruction
	// TODO: merge current IDL generation PR and update this
	DiscriminatorExecute = [8]byte{
		130,
		221,
		242,
		154,
		13,
		193,
		189,
		29,
	}

	// DiscriminatorWithdrawSPL returns the discriminator for Solana gateway 'withdraw_spl_token' instruction
	DiscriminatorWithdrawSPL = idlgateway.IDLGateway.GetDiscriminator("withdraw_spl_token")

	// DiscriminatorWhitelist returns the discriminator for Solana gateway 'whitelist_spl_mint' instruction
	DiscriminatorWhitelistSplMint = idlgateway.IDLGateway.GetDiscriminator("whitelist_spl_mint")
)

// ParseGatewayWithPDA parses the gateway id and program derived address from the given string
func ParseGatewayWithPDA(gatewayAddress string) (solana.PublicKey, solana.PublicKey, error) {
	var gatewayID, pda solana.PublicKey

	// decode gateway address
	gatewayID, err := solana.PublicKeyFromBase58(gatewayAddress)
	if err != nil {
		return gatewayID, pda, errors.Wrap(err, "unable to decode address")
	}

	// compute gateway PDA
	seed := []byte(PDASeed)
	pda, _, err = solana.FindProgramAddress([][]byte{seed}, gatewayID)

	return gatewayID, pda, err
}

func ComputeConnectedPdaAddress(connected solana.PublicKey) (solana.PublicKey, error) {
	seed := []byte("connected")
	pdaComputed, _, err := solana.FindProgramAddress([][]byte{seed}, connected)
	if err != nil {
		return solana.PublicKey{}, err
	}

	return pdaComputed, nil
}
