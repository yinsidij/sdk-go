// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	history "go.temporal.io/api/history/v1"

	mock "github.com/stretchr/testify/mock"
)

// HistoryEventIterator is an autogenerated mock type for the HistoryEventIterator type
type HistoryEventIterator struct {
	mock.Mock
}

// HasNext provides a mock function with given fields:
func (_m *HistoryEventIterator) HasNext() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HasNext")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *HistoryEventIterator) Next() (*history.HistoryEvent, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Next")
	}

	var r0 *history.HistoryEvent
	var r1 error
	if rf, ok := ret.Get(0).(func() (*history.HistoryEvent, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *history.HistoryEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*history.HistoryEvent)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHistoryEventIterator creates a new instance of HistoryEventIterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHistoryEventIterator(t interface {
	mock.TestingT
	Cleanup(func())
}) *HistoryEventIterator {
	mock := &HistoryEventIterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
