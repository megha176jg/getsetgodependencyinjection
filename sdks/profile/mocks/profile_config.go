// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ProfileConfig is an autogenerated mock type for the ProfileConfig type
type ProfileConfig struct {
	mock.Mock
}

// GetJWRAuthToken provides a mock function with given fields:
func (_m *ProfileConfig) GetJWRAuthToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetProfileEndpoint provides a mock function with given fields:
func (_m *ProfileConfig) GetProfileEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewProfileConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewProfileConfig creates a new instance of ProfileConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProfileConfig(t mockConstructorTestingTNewProfileConfig) *ProfileConfig {
	mock := &ProfileConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
