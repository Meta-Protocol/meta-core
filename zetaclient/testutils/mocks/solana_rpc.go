// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/gagliardetto/solana-go/rpc"

	solana "github.com/gagliardetto/solana-go"
)

// SolanaRPCClient is an autogenerated mock type for the SolanaRPCClient type
type SolanaRPCClient struct {
	mock.Mock
}

// GetAccountInfo provides a mock function with given fields: ctx, account
func (_m *SolanaRPCClient) GetAccountInfo(ctx context.Context, account solana.PublicKey) (*rpc.GetAccountInfoResult, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountInfo")
	}

	var r0 *rpc.GetAccountInfoResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey) (*rpc.GetAccountInfoResult, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey) *rpc.GetAccountInfoResult); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetAccountInfoResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, account, commitment
func (_m *SolanaRPCClient) GetBalance(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType) (*rpc.GetBalanceResult, error) {
	ret := _m.Called(ctx, account, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 *rpc.GetBalanceResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) (*rpc.GetBalanceResult, error)); ok {
		return rf(ctx, account, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) *rpc.GetBalanceResult); ok {
		r0 = rf(ctx, account, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetBalanceResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, account, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfirmedTransactionWithOpts provides a mock function with given fields: ctx, signature, opts
func (_m *SolanaRPCClient) GetConfirmedTransactionWithOpts(ctx context.Context, signature solana.Signature, opts *rpc.GetTransactionOpts) (*rpc.TransactionWithMeta, error) {
	ret := _m.Called(ctx, signature, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetConfirmedTransactionWithOpts")
	}

	var r0 *rpc.TransactionWithMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) (*rpc.TransactionWithMeta, error)); ok {
		return rf(ctx, signature, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) *rpc.TransactionWithMeta); ok {
		r0 = rf(ctx, signature, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.TransactionWithMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) error); ok {
		r1 = rf(ctx, signature, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHealth provides a mock function with given fields: ctx
func (_m *SolanaRPCClient) GetHealth(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetHealth")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlockhash provides a mock function with given fields: ctx, commitment
func (_m *SolanaRPCClient) GetLatestBlockhash(ctx context.Context, commitment rpc.CommitmentType) (*rpc.GetLatestBlockhashResult, error) {
	ret := _m.Called(ctx, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBlockhash")
	}

	var r0 *rpc.GetLatestBlockhashResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) (*rpc.GetLatestBlockhashResult, error)); ok {
		return rf(ctx, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) *rpc.GetLatestBlockhashResult); ok {
		r0 = rf(ctx, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetLatestBlockhashResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, commitment)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetRecentPrioritizationFees provides a mock function with given fields: ctx, accounts
func (_m *SolanaRPCClient) GetRecentPrioritizationFees(ctx context.Context, accounts solana.PublicKeySlice) ([]rpc.PriorizationFeeResult, error) {
	ret := _m.Called(ctx, accounts)

	if len(ret) == 0 {
		panic("no return value specified for GetRecentPrioritizationFees")
	}

	var r0 []rpc.PriorizationFeeResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKeySlice) ([]rpc.PriorizationFeeResult, error)); ok {
		return rf(ctx, accounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKeySlice) []rpc.PriorizationFeeResult); ok {
		r0 = rf(ctx, accounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rpc.PriorizationFeeResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKeySlice) error); ok {
		r1 = rf(ctx, accounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignaturesForAddressWithOpts provides a mock function with given fields: ctx, account, opts
func (_m *SolanaRPCClient) GetSignaturesForAddressWithOpts(ctx context.Context, account solana.PublicKey, opts *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error) {
	ret := _m.Called(ctx, account, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetSignaturesForAddressWithOpts")
	}

	var r0 []*rpc.TransactionSignature
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error)); ok {
		return rf(ctx, account, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) []*rpc.TransactionSignature); ok {
		r0 = rf(ctx, account, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rpc.TransactionSignature)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) error); ok {
		r1 = rf(ctx, account, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlot provides a mock function with given fields: ctx, commitment
func (_m *SolanaRPCClient) GetSlot(ctx context.Context, commitment rpc.CommitmentType) (uint64, error) {
	ret := _m.Called(ctx, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) (uint64, error)); ok {
		return rf(ctx, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) uint64); ok {
		r0 = rf(ctx, commitment)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: ctx, txSig, opts
func (_m *SolanaRPCClient) GetTransaction(ctx context.Context, txSig solana.Signature, opts *rpc.GetTransactionOpts) (*rpc.GetTransactionResult, error) {
	ret := _m.Called(ctx, txSig, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 *rpc.GetTransactionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) (*rpc.GetTransactionResult, error)); ok {
		return rf(ctx, txSig, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) *rpc.GetTransactionResult); ok {
		r0 = rf(ctx, txSig, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetTransactionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) error); ok {
		r1 = rf(ctx, txSig, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVersion provides a mock function with given fields: ctx
func (_m *SolanaRPCClient) GetVersion(ctx context.Context) (*rpc.GetVersionResult, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetVersion")
	}

	var r0 *rpc.GetVersionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpc.GetVersionResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpc.GetVersionResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetVersionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransactionWithOpts provides a mock function with given fields: ctx, transaction, opts
func (_m *SolanaRPCClient) SendTransactionWithOpts(ctx context.Context, transaction *solana.Transaction, opts rpc.TransactionOpts) (solana.Signature, error) {
	ret := _m.Called(ctx, transaction, opts)

	if len(ret) == 0 {
		panic("no return value specified for SendTransactionWithOpts")
	}

	var r0 solana.Signature
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *solana.Transaction, rpc.TransactionOpts) (solana.Signature, error)); ok {
		return rf(ctx, transaction, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *solana.Transaction, rpc.TransactionOpts) solana.Signature); ok {
		r0 = rf(ctx, transaction, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(solana.Signature)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *solana.Transaction, rpc.TransactionOpts) error); ok {
		r1 = rf(ctx, transaction, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSolanaRPCClient creates a new instance of SolanaRPCClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSolanaRPCClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SolanaRPCClient {
	mock := &SolanaRPCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
