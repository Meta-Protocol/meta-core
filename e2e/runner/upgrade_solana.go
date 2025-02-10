package runner

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

// Helper function to convert uint64 to big-endian bytes
func uint64ToBytes(val uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	return buf
}

func (r *E2ERunner) UpgradeGatewayContract(deployerPrivateKey string) {

	solFile := "gateway.so"
	root := "/root"
	const MAX_CHUNK_SIZE = 900
	filePath := filepath.Join(root, solFile)
	programID := r.GatewayProgram
	// compute the gateway PDA address
	pdaComputed := r.ComputePdaAddress()

	// get deployer account balance
	upgradeAuthorityPrivkey, err := solana.PrivateKeyFromBase58(deployerPrivateKey)
	require.NoError(r, err)
	bal, err := r.SolanaClient.GetBalance(r.Ctx, upgradeAuthorityPrivkey.PublicKey(), rpc.CommitmentConfirmed)
	require.NoError(r, err)
	r.Logger.Print("deployer address: %s, balance: %f SOL", upgradeAuthorityPrivkey.PublicKey().String(), float64(bal.Value)/1e9)
	r.Logger.Print("⚙️ upgrading gateway program on Solana %s , %s ,%s", pdaComputed, programID, upgradeAuthorityPrivkey.String())

	programData, err := os.ReadFile(filePath)
	require.NoError(r, err)

	bufferSize := uint64(len(programData))

	rentExemption, err := r.SolanaClient.GetMinimumBalanceForRentExemption(
		r.Ctx,
		bufferSize,
		rpc.CommitmentConfirmed,
	)
	require.NoError(r, err)
	bp := solana.NewWallet()
	// Create system program instruction to allocate buffer
	allocateIx := system.NewCreateAccountInstruction(
		rentExemption,
		bufferSize,
		solana.BPFLoaderUpgradeableProgramID,
		upgradeAuthorityPrivkey.PublicKey(),
		bp.PublicKey(),
	).Build()

	//initBufferIx := solana.NewInstruction(
	//	solana.BPFLoaderUpgradeableProgramID,
	//	[]*solana.AccountMeta{
	//		{PublicKey: upgradeAuthorityPrivkey.PublicKey(), IsWritable: true, IsSigner: true},
	//	},
	//	[]byte{0}, // 0 is initialize buffer command
	//)

	r.Logger.Print("⚙️ CreateSignedTransaction initBufferIx buffer %s , %s", pdaComputed, programID)
	// create and sign the transaction
	signedTx := r.CreateSignedTransaction([]solana.Instruction{allocateIx}, upgradeAuthorityPrivkey, []solana.PrivateKey{})

	r.Logger.Print("⚙️ BroadcastTxSync initBufferIx buffer %s , %s", pdaComputed, programID)
	// broadcast the transaction and wait for finalization
	_, out := r.BroadcastTxSync(signedTx)
	r.Logger.Print("create initBufferIx logs: %v", out.Meta.LogMessages)

	for offset := 0; offset < len(programData); offset += MAX_CHUNK_SIZE {
		end := offset + MAX_CHUNK_SIZE
		if end > len(programData) {
			end = len(programData)
		}
		chunk := programData[offset:end]
		writeData := make([]byte, 4+len(chunk)) // 4 bytes for offset + chunk data
		binary.LittleEndian.PutUint32(writeData[0:4], uint32(offset))
		copy(writeData[4:], chunk)

		// Write program data to buffer instruction
		writeIx := solana.NewInstruction(
			solana.BPFLoaderUpgradeableProgramID,
			[]*solana.AccountMeta{
				{PublicKey: pdaComputed, IsWritable: true, IsSigner: false},
				{PublicKey: upgradeAuthorityPrivkey.PublicKey(), IsWritable: false, IsSigner: true},
			},
			append([]byte{2}, // 2 is write command
				writeData...,
			),
		)
		r.Logger.Print("⚙️ CreateSignedTransaction writeIx buffer %s , %s", pdaComputed, programID)
		// create and sign the transaction
		signedTx := r.CreateSignedTransaction([]solana.Instruction{writeIx}, upgradeAuthorityPrivkey, []solana.PrivateKey{})

		r.Logger.Print("⚙️ BroadcastTxSync writeIx buffer %s , %s", pdaComputed, programID)
		// broadcast the transaction and wait for finalization
		_, out := r.BroadcastTxSync(signedTx)
		r.Logger.Print("create writeIx logs: %v", out.Meta.LogMessages)

		time.Sleep(time.Millisecond * 100)
	}

	upgradeIx := solana.NewInstruction(
		solana.BPFLoaderUpgradeableProgramID,
		[]*solana.AccountMeta{
			{PublicKey: pdaComputed, IsWritable: true, IsSigner: false},
			{PublicKey: programID, IsWritable: true, IsSigner: false},
			{PublicKey: upgradeAuthorityPrivkey.PublicKey(), IsWritable: false, IsSigner: true},
			{PublicKey: solana.SysVarRentPubkey, IsWritable: false, IsSigner: false},
			{PublicKey: solana.SysVarClockPubkey, IsWritable: false, IsSigner: false},
		},
		[]byte{3}, // 3 is upgrade command
	)

	r.Logger.Print("⚙️ CreateSignedTransaction upgradeIx %s , %s", pdaComputed, programID)
	// create and sign the transaction
	signedTx = r.CreateSignedTransaction([]solana.Instruction{upgradeIx}, upgradeAuthorityPrivkey, []solana.PrivateKey{})
	r.Logger.Print("⚙️ BroadcastTxSync upgradeIx %s , %s", pdaComputed, programID)
	// broadcast the transaction and wait for finalization
	_, out = r.BroadcastTxSync(signedTx)
	r.Logger.Print("initialize upgradeIx: %v", out.Meta.LogMessages)

}
