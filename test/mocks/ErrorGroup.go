// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ErrorGroup is an autogenerated mock type for the ErrorGroup type
type ErrorGroup struct {
	mock.Mock
}

// Append provides a mock function with given fields: _a0
func (_m *ErrorGroup) Append(_a0 error) {
	_m.Called(_a0)
}

// Empty provides a mock function with given fields:
func (_m *ErrorGroup) Empty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *ErrorGroup) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Errors provides a mock function with given fields:
func (_m *ErrorGroup) Errors() []error {
	ret := _m.Called()

	var r0 []error
	if rf, ok := ret.Get(0).(func() []error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}
