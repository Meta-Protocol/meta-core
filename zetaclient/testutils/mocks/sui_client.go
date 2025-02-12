// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/block-vision/sui-go-sdk/models"

	time "time"
)

// SuiClient is an autogenerated mock type for the suiClient type
type SuiClient struct {
	mock.Mock
}

// GetLatestCheckpoint provides a mock function with given fields: ctx
func (_m *SuiClient) GetLatestCheckpoint(ctx context.Context) (models.CheckpointResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestCheckpoint")
	}

	var r0 models.CheckpointResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (models.CheckpointResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) models.CheckpointResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(models.CheckpointResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *SuiClient) HealthCheck(ctx context.Context) (time.Time, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for HealthCheck")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (time.Time, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) time.Time); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuiXGetReferenceGasPrice provides a mock function with given fields: ctx
func (_m *SuiClient) SuiXGetReferenceGasPrice(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuiXGetReferenceGasPrice")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuiXQueryEvents provides a mock function with given fields: ctx, req
func (_m *SuiClient) SuiXQueryEvents(ctx context.Context, req models.SuiXQueryEventsRequest) (models.PaginatedEventsResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for SuiXQueryEvents")
	}

	var r0 models.PaginatedEventsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SuiXQueryEventsRequest) (models.PaginatedEventsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.SuiXQueryEventsRequest) models.PaginatedEventsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(models.PaginatedEventsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.SuiXQueryEventsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSuiClient creates a new instance of SuiClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSuiClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SuiClient {
	mock := &SuiClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
