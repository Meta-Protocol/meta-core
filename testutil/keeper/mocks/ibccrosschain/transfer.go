// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LightclientTransferKeeper is an autogenerated mock type for the LightclientTransferKeeper type
type LightclientTransferKeeper struct {
	mock.Mock
}

// NewLightclientTransferKeeper creates a new instance of LightclientTransferKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLightclientTransferKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *LightclientTransferKeeper {
	mock := &LightclientTransferKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
