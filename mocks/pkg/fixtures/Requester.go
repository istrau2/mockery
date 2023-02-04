// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Requester is an autogenerated mock type for the Requester type
type Requester struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *Requester) Get(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	var r1 error

	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		r0, r1 = rf(path)
	} else {
		if rf, ok := ret.Get(0).(func(string) string); ok {
			r0 = rf(path)
		} else {
			r0 = ret.Get(0).(string)
		}

		if rf, ok := ret.Get(1).(func(string) error); ok {
			r1 = rf(path)
		} else {
			r1 = ret.Error(1)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewRequester interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequester creates a new instance of Requester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequester(t mockConstructorTestingTNewRequester) *Requester {
	mock := &Requester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
