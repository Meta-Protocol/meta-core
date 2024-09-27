package sample

import (
	"crypto/rand"
	"testing"
	"time"

	"cosmossdk.io/math"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/tonkeeper/tongo/boc"
	"github.com/tonkeeper/tongo/tlb"
	"github.com/tonkeeper/tongo/ton"

	toncontracts "github.com/zeta-chain/node/pkg/contracts/ton"
)

const (
	tonWorkchainID    = 0
	tonShardID        = 123
	tonDepositFee     = 10_000_000 // 0.01 TON
	tonSampleGasUsage = 50_000_000 // 0.05 TON
)

type TONTransactionProps struct {
	Account ton.AccountID
	GasUsed uint64
	BlockID ton.BlockIDExt

	// For simplicity let's have only one input
	// and one output (both optional)
	Input  *tlb.Message
	Output *tlb.Message
}

type intMsgInfo struct {
	IhrDisabled bool
	Bounce      bool
	Bounced     bool
	Src         tlb.MsgAddress
	Dest        tlb.MsgAddress
	Value       tlb.CurrencyCollection
	IhrFee      tlb.Grams
	FwdFee      tlb.Grams
	CreatedLt   uint64
	CreatedAt   uint32
}

func TONDonateProps(t *testing.T, acc ton.AccountID, d toncontracts.Donation) TONTransactionProps {
	body, err := d.AsBody()
	require.NoError(t, err)

	deposited := tonSampleGasUsage + d.Amount.Uint64()

	return TONTransactionProps{
		Account: acc,
		Input: &tlb.Message{
			Info: internalMessageInfo(&intMsgInfo{
				Bounce: true,
				Src:    d.Sender.ToMsgAddress(),
				Dest:   acc.ToMsgAddress(),
				Value:  tlb.CurrencyCollection{Grams: tlb.Grams(deposited)},
			}),
			Body: tlb.EitherRef[tlb.Any]{Value: tlb.Any(*body)},
		},
	}
}

func TONDepositProps(t *testing.T, acc ton.AccountID, d toncontracts.Deposit) TONTransactionProps {
	body, err := d.AsBody()
	require.NoError(t, err)

	logBody := depositLogMock(t, d.Sender, d.Amount.Uint64(), d.Recipient, nil)

	return TONTransactionProps{
		Account: acc,
		Input: &tlb.Message{
			Info: internalMessageInfo(&intMsgInfo{
				Bounce: true,
				Src:    d.Sender.ToMsgAddress(),
				Dest:   acc.ToMsgAddress(),
				Value:  tlb.CurrencyCollection{Grams: fakeDepositAmount(d.Amount)},
			}),
			Body: tlb.EitherRef[tlb.Any]{Value: tlb.Any(*body)},
		},
		Output: &tlb.Message{
			Body: tlb.EitherRef[tlb.Any]{IsRight: true, Value: tlb.Any(*logBody)},
		},
	}
}

func TONDepositAndCallProps(t *testing.T, acc ton.AccountID, d toncontracts.DepositAndCall) TONTransactionProps {
	body, err := d.AsBody()
	require.NoError(t, err)

	logBody := depositLogMock(t, d.Sender, d.Amount.Uint64(), d.Recipient, d.CallData)

	return TONTransactionProps{
		Account: acc,
		Input: &tlb.Message{
			Info: internalMessageInfo(&intMsgInfo{
				Bounce: true,
				Src:    d.Sender.ToMsgAddress(),
				Dest:   acc.ToMsgAddress(),
				Value:  tlb.CurrencyCollection{Grams: fakeDepositAmount(d.Amount)},
			}),
			Body: tlb.EitherRef[tlb.Any]{Value: tlb.Any(*body)},
		},
		Output: &tlb.Message{
			Body: tlb.EitherRef[tlb.Any]{IsRight: true, Value: tlb.Any(*logBody)},
		},
	}
}

// TONTransaction creates a sample TON transaction.
func TONTransaction(t *testing.T, p TONTransactionProps) ton.Transaction {
	require.False(t, p.Account.IsZero(), "account address is empty")
	require.False(t, p.Input == nil && p.Output == nil, "both input and output are empty")

	now := time.Now().UTC()

	if p.GasUsed == 0 {
		p.GasUsed = tonSampleGasUsage
	}

	if p.BlockID.BlockID.Seqno == 0 {
		p.BlockID = tonBlockID(now)
	}

	// Simulate logical time as `2 * now()`
	lt := uint64(2 * now.Unix())

	input := tlb.Maybe[tlb.Ref[tlb.Message]]{}
	if p.Input != nil {
		input.Exists = true
		input.Value.Value = *p.Input
	}

	var outputs tlb.HashmapE[tlb.Uint15, tlb.Ref[tlb.Message]]
	if p.Output != nil {
		outputs = tlb.NewHashmapE(
			[]tlb.Uint15{0},
			[]tlb.Ref[tlb.Message]{{*p.Output}},
		)
	}

	type messages struct {
		InMsg   tlb.Maybe[tlb.Ref[tlb.Message]]
		OutMsgs tlb.HashmapE[tlb.Uint15, tlb.Ref[tlb.Message]]
	}

	return ton.Transaction{
		BlockID: p.BlockID,
		Transaction: tlb.Transaction{
			AccountAddr: p.Account.Address,
			Lt:          lt,
			Now:         uint32(now.Unix()),
			OutMsgCnt:   tlb.Uint15(len(outputs.Keys())),
			TotalFees:   tlb.CurrencyCollection{Grams: tlb.Grams(p.GasUsed)},
			Msgs:        messages{InMsg: input, OutMsgs: outputs},
		},
	}
}

func GenerateTONAccountID() ton.AccountID {
	var addr [32]byte

	//nolint:errcheck // test code
	rand.Read(addr[:])

	return *ton.NewAccountID(0, addr)
}

func internalMessageInfo(info *intMsgInfo) tlb.CommonMsgInfo {
	return tlb.CommonMsgInfo{
		SumType: "IntMsgInfo",
		IntMsgInfo: (*struct {
			IhrDisabled bool
			Bounce      bool
			Bounced     bool
			Src         tlb.MsgAddress
			Dest        tlb.MsgAddress
			Value       tlb.CurrencyCollection
			IhrFee      tlb.Grams
			FwdFee      tlb.Grams
			CreatedLt   uint64
			CreatedAt   uint32
		})(info),
	}
}

func tonBlockID(now time.Time) ton.BlockIDExt {
	// simulate shard seqno as unix timestamp
	seqno := uint32(now.Unix())

	return ton.BlockIDExt{
		BlockID: ton.BlockID{
			Workchain: tonWorkchainID,
			Shard:     tonShardID,
			Seqno:     seqno,
		},
	}
}

func fakeDepositAmount(v math.Uint) tlb.Grams {
	return tlb.Grams(v.Uint64() + tonDepositFee)
}

func depositLogMock(
	t *testing.T,
	sender ton.AccountID,
	amount uint64,
	recipient eth.Address,
	callData []byte,
) *boc.Cell {
	//     cell log = begin_cell()
	//        .store_uint(op::internal::deposit_and_call, size::op_code_size)
	//        .store_uint(0, size::query_id_size)
	//        .store_slice(sender)
	//        .store_coins(deposit_amount)
	//        .store_uint(evm_recipient, size::evm_address)
	//        .store_ref(call_data) // only for DepositAndCall
	//        .end_cell();

	b := boc.NewCell()
	require.NoError(t, b.WriteUint(0, 32+64))

	// skip
	msgAddr := sender.ToMsgAddress()
	require.NoError(t, tlb.Marshal(b, msgAddr))

	coins := tlb.Grams(amount)
	require.NoError(t, coins.MarshalTLB(b, nil))

	require.NoError(t, b.WriteBytes(recipient.Bytes()))

	if callData != nil {
		callDataCell, err := toncontracts.MarshalSnakeCell(callData)
		require.NoError(t, err)
		require.NoError(t, b.AddRef(callDataCell))
	}

	return b
}
