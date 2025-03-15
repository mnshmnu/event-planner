// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "event-planner/internal/entities"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Model is an autogenerated mock type for the Model type
type Model struct {
	mock.Mock
}

// AddAvailability provides a mock function with given fields: ctx, availability
func (_m *Model) AddAvailability(ctx context.Context, availability *entities.Availability) error {
	ret := _m.Called(ctx, availability)

	if len(ret) == 0 {
		panic("no return value specified for AddAvailability")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Availability) error); ok {
		r0 = rf(ctx, availability)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMeeting provides a mock function with given fields: ctx, meeting
func (_m *Model) CreateMeeting(ctx context.Context, meeting *entities.Meeting) error {
	ret := _m.Called(ctx, meeting)

	if len(ret) == 0 {
		panic("no return value specified for CreateMeeting")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Meeting) error); ok {
		r0 = rf(ctx, meeting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, user, hPass
func (_m *Model) CreateUser(ctx context.Context, user *entities.User, hPass string) error {
	ret := _m.Called(ctx, user, hPass)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User, string) error); ok {
		r0 = rf(ctx, user, hPass)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAvailabilities provides a mock function with given fields: ctx, duration
func (_m *Model) GetAvailabilities(ctx context.Context, duration int) (map[time.Time][]int, error) {
	ret := _m.Called(ctx, duration)

	if len(ret) == 0 {
		panic("no return value specified for GetAvailabilities")
	}

	var r0 map[time.Time][]int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (map[time.Time][]int, error)); ok {
		return rf(ctx, duration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) map[time.Time][]int); ok {
		r0 = rf(ctx, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[time.Time][]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Model) GetUserByEmail(ctx context.Context, email string) (*entities.User, string, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *entities.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.User, string, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewModel creates a new instance of Model. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *Model {
	mock := &Model{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
