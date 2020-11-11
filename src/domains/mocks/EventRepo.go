// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domains "tiket.vip/src/domains"
)

// EventRepo is an autogenerated mock type for the EventRepo type
type EventRepo struct {
	mock.Mock
}

// CreateEvent provides a mock function with given fields: _a0
func (_m *EventRepo) CreateEvent(_a0 domains.Event) (*domains.Event, error) {
	ret := _m.Called(_a0)

	var r0 *domains.Event
	if rf, ok := ret.Get(0).(func(domains.Event) *domains.Event); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domains.Event) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEvent provides a mock function with given fields: _a0
func (_m *EventRepo) DeleteEvent(_a0 string) (*string, error) {
	ret := _m.Called(_a0)

	var r0 *string
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllEvents provides a mock function with given fields:
func (_m *EventRepo) GetAllEvents() (*[]domains.Event, error) {
	ret := _m.Called()

	var r0 *[]domains.Event
	if rf, ok := ret.Get(0).(func() *[]domains.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domains.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllEventsPaginate provides a mock function with given fields: _a0
func (_m *EventRepo) GetAllEventsPaginate(_a0 domains.EventPagi) (*domains.Events, error) {
	ret := _m.Called(_a0)

	var r0 *domains.Events
	if rf, ok := ret.Get(0).(func(domains.EventPagi) *domains.Events); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Events)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domains.EventPagi) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvent provides a mock function with given fields: _a0
func (_m *EventRepo) GetEvent(_a0 string) (*domains.Event, error) {
	ret := _m.Called(_a0)

	var r0 *domains.Event
	if rf, ok := ret.Get(0).(func(string) *domains.Event); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: _a0
func (_m *EventRepo) UpdateEvent(_a0 domains.Event) (*domains.Event, error) {
	ret := _m.Called(_a0)

	var r0 *domains.Event
	if rf, ok := ret.Get(0).(func(domains.Event) *domains.Event); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domains.Event) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
