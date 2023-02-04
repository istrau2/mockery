// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	fixtures "github.com/vektra/mockery/v2/pkg/fixtures"

	test "github.com/vektra/mockery/v2/pkg/fixtures/test"
)

// ImportsSameAsPackage is an autogenerated mock type for the ImportsSameAsPackage type
type ImportsSameAsPackage struct {
	mock.Mock
}

// A provides a mock function with given fields:
func (_m *ImportsSameAsPackage) A() test.B {
	ret := _m.Called()

	var r0 test.B

	if rf, ok := ret.Get(0).(func() test.B); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(test.B)
	}

	return r0
}

// B provides a mock function with given fields:
func (_m *ImportsSameAsPackage) B() fixtures.KeyManager {
	ret := _m.Called()

	var r0 fixtures.KeyManager

	if rf, ok := ret.Get(0).(func() fixtures.KeyManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fixtures.KeyManager)
		}
	}

	return r0
}

// C provides a mock function with given fields: _a0
func (_m *ImportsSameAsPackage) C(_a0 fixtures.C) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewImportsSameAsPackage interface {
	mock.TestingT
	Cleanup(func())
}

// NewImportsSameAsPackage creates a new instance of ImportsSameAsPackage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImportsSameAsPackage(t mockConstructorTestingTNewImportsSameAsPackage) *ImportsSameAsPackage {
	mock := &ImportsSameAsPackage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
