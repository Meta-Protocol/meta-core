package signer

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/gagliardetto/solana-go"
	"github.com/near/borsh-go"

	"github.com/zeta-chain/node/pkg/chains"
	contracts "github.com/zeta-chain/node/pkg/contracts/solana"
	"github.com/zeta-chain/node/x/crosschain/types"
)

// createAndSignMsgExecuteSPL creates and signs a execute spl message for gateway execute_spl_token instruction with TSS.
func (signer *Signer) createAndSignMsgExecuteSPL(
	ctx context.Context,
	params *types.OutboundParams,
	height uint64,
	asset string,
	decimals uint8,
	sender [20]byte,
	data []byte,
	remainingAccounts []*solana.AccountMeta,
	cancelTx bool,
) (*contracts.MsgExecuteSPL, error) {
	chain := signer.Chain()
	// #nosec G115 always positive
	chainID := uint64(signer.Chain().ChainId)
	nonce := params.TssNonce
	amount := params.Amount.Uint64()

	// zero out the amount if cancelTx is set. It's legal to withdraw 0 spl through the gateway.
	if cancelTx {
		amount = 0
	}

	// check receiver address
	to, err := chains.DecodeSolanaWalletAddress(params.Receiver)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot decode receiver address %s", params.Receiver)
	}

	// parse mint account
	mintAccount, err := solana.PublicKeyFromBase58(asset)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse asset public key %s", asset)
	}

	// get recipient ata
	destinationProgramPda, err := contracts.ComputeConnectedSPLPdaAddress(to)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode connected spl pda address")
	}

	destinationProgramPdaAta, _, err := solana.FindAssociatedTokenAddress(destinationProgramPda, mintAccount)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot find ATA for %s and mint account %s", destinationProgramPda, mintAccount)
	}

	// prepare execute spl msg and compute hash
	msg := contracts.NewMsgExecuteSPL(
		chainID,
		nonce,
		amount,
		decimals,
		mintAccount,
		to,
		destinationProgramPdaAta,
		sender,
		data,
		remainingAccounts,
	)
	msgHash := msg.Hash()

	// sign the message with TSS to get an ECDSA signature.
	// the produced signature is in the [R || S || V] format where V is 0 or 1.
	signature, err := signer.TSS().Sign(ctx, msgHash[:], height, nonce, chain.ChainId)
	if err != nil {
		return nil, errors.Wrap(err, "key-sign failed")
	}

	// attach the signature and return
	return msg.SetSignature(signature), nil
}

// createExecuteSPLInstruction wraps the execute spl 'msg' into a Solana instruction.
func (signer *Signer) createExecuteSPLInstruction(msg contracts.MsgExecuteSPL) (*solana.GenericInstruction, error) {
	// create execute spl instruction with program call data
	dataBytes, err := borsh.Serialize(contracts.ExecuteSPLInstructionParams{
		Discriminator: contracts.DiscriminatorExecuteSPL,
		Decimals:      msg.Decimals(),
		Amount:        msg.Amount(),
		Sender:        msg.Sender(),
		Data:          msg.Data(),
		Signature:     msg.SigRS(),
		RecoveryID:    msg.SigV(),
		MessageHash:   msg.Hash(),
		Nonce:         msg.Nonce(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot serialize execute spl instruction")
	}

	pdaAta, _, err := solana.FindAssociatedTokenAddress(signer.pda, msg.MintAccount())
	if err != nil {
		return nil, errors.Wrapf(err, "cannot find ATA for %s and mint account %s", signer.pda, msg.MintAccount())
	}

	destinationProgramPda, err := contracts.ComputeConnectedSPLPdaAddress(msg.To())
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode connected spl pda address")
	}

	predefinedAccounts := []*solana.AccountMeta{
		solana.Meta(signer.relayerKey.PublicKey()).WRITE().SIGNER(),
		solana.Meta(signer.pda).WRITE(),
		solana.Meta(pdaAta).WRITE(),
		solana.Meta(msg.MintAccount()),
		solana.Meta(msg.To()),
		solana.Meta(destinationProgramPda),
		solana.Meta(msg.RecipientAta()).WRITE(),
		solana.Meta(solana.TokenProgramID),
		solana.Meta(solana.SPLAssociatedTokenAccountProgramID),
		solana.Meta(solana.SystemProgramID),
	}
	allAccounts := append(predefinedAccounts, msg.RemainingAccounts()...)

	inst := &solana.GenericInstruction{
		ProgID:        signer.gatewayID,
		DataBytes:     dataBytes,
		AccountValues: allAccounts,
	}

	return inst, nil
}
