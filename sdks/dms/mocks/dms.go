// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dms "bitbucket.org/junglee_games/getsetgo/sdks/dms"
	mock "github.com/stretchr/testify/mock"
)

// DMS is an autogenerated mock type for the DMS type
type DMS struct {
	mock.Mock
}

// Initiate provides a mock function with given fields: req
func (_m *DMS) Initiate(req dms.IntiateRequest) (dms.IntiateResponse, error) {
	ret := _m.Called(req)

	var r0 dms.IntiateResponse
	if rf, ok := ret.Get(0).(func(dms.IntiateRequest) dms.IntiateResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(dms.IntiateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dms.IntiateRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDMS interface {
	mock.TestingT
	Cleanup(func())
}

// NewDMS creates a new instance of DMS. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDMS(t mockConstructorTestingTNewDMS) *DMS {
	mock := &DMS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
