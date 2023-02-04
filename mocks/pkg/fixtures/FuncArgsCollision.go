// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FuncArgsCollision is an autogenerated mock type for the FuncArgsCollision type
type FuncArgsCollision struct {
	mock.Mock
}

// Foo provides a mock function with given fields: ret
func (_m *FuncArgsCollision) Foo(ret interface{}) error {
	ret_1 := _m.Called(ret)

	var r0 error

	if rf, ok := ret_1.Get(0).(func(interface{}) error); ok {
		r0 = rf(ret)
	} else {
		r0 = ret_1.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFuncArgsCollision interface {
	mock.TestingT
	Cleanup(func())
}

// NewFuncArgsCollision creates a new instance of FuncArgsCollision. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFuncArgsCollision(t mockConstructorTestingTNewFuncArgsCollision) *FuncArgsCollision {
	mock := &FuncArgsCollision{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
