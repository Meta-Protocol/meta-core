// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	observertypes "github.com/zeta-chain/node/x/observer/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// EmissionObserverKeeper is an autogenerated mock type for the EmissionObserverKeeper type
type EmissionObserverKeeper struct {
	mock.Mock
}

// ClearMaturedBallotsAndBallotList provides a mock function with given fields: ctx, maturityBlocks
func (_m *EmissionObserverKeeper) ClearMaturedBallotsAndBallotList(ctx types.Context, maturityBlocks int64) {
	_m.Called(ctx, maturityBlocks)
}

// GetBallot provides a mock function with given fields: ctx, index
func (_m *EmissionObserverKeeper) GetBallot(ctx types.Context, index string) (observertypes.Ballot, bool) {
	ret := _m.Called(ctx, index)

	if len(ret) == 0 {
		panic("no return value specified for GetBallot")
	}

	var r0 observertypes.Ballot
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string) (observertypes.Ballot, bool)); ok {
		return rf(ctx, index)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) observertypes.Ballot); ok {
		r0 = rf(ctx, index)
	} else {
		r0 = ret.Get(0).(observertypes.Ballot)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) bool); ok {
		r1 = rf(ctx, index)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetMaturedBallots provides a mock function with given fields: ctx, maturityBlocks
func (_m *EmissionObserverKeeper) GetMaturedBallots(ctx types.Context, maturityBlocks int64) (observertypes.BallotListForHeight, bool) {
	ret := _m.Called(ctx, maturityBlocks)

	if len(ret) == 0 {
		panic("no return value specified for GetMaturedBallots")
	}

	var r0 observertypes.BallotListForHeight
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, int64) (observertypes.BallotListForHeight, bool)); ok {
		return rf(ctx, maturityBlocks)
	}
	if rf, ok := ret.Get(0).(func(types.Context, int64) observertypes.BallotListForHeight); ok {
		r0 = rf(ctx, maturityBlocks)
	} else {
		r0 = ret.Get(0).(observertypes.BallotListForHeight)
	}

	if rf, ok := ret.Get(1).(func(types.Context, int64) bool); ok {
		r1 = rf(ctx, maturityBlocks)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewEmissionObserverKeeper creates a new instance of EmissionObserverKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmissionObserverKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmissionObserverKeeper {
	mock := &EmissionObserverKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
