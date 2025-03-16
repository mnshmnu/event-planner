// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "event-planner/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// Model is an autogenerated mock type for the Model type
type Model struct {
	mock.Mock
}

// CreateAvailability provides a mock function with given fields: ctx, availability
func (_m *Model) CreateAvailability(ctx context.Context, availability *entities.ParticipantAvailability) (int64, error) {
	ret := _m.Called(ctx, availability)

	if len(ret) == 0 {
		panic("no return value specified for CreateAvailability")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParticipantAvailability) (int64, error)); ok {
		return rf(ctx, availability)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParticipantAvailability) int64); ok {
		r0 = rf(ctx, availability)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.ParticipantAvailability) error); ok {
		r1 = rf(ctx, availability)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEvent provides a mock function with given fields: ctx, event
func (_m *Model) CreateEvent(ctx context.Context, event *entities.Event) (int64, error) {
	ret := _m.Called(ctx, event)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvent")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Event) (int64, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Event) int64); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.Event) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEventSlot provides a mock function with given fields: ctx, slot
func (_m *Model) CreateEventSlot(ctx context.Context, slot *entities.EventSlot) (int64, error) {
	ret := _m.Called(ctx, slot)

	if len(ret) == 0 {
		panic("no return value specified for CreateEventSlot")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.EventSlot) (int64, error)); ok {
		return rf(ctx, slot)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.EventSlot) int64); ok {
		r0 = rf(ctx, slot)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.EventSlot) error); ok {
		r1 = rf(ctx, slot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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

// DeleteAvailability provides a mock function with given fields: ctx, id
func (_m *Model) DeleteAvailability(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAvailability")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEvent provides a mock function with given fields: ctx, id
func (_m *Model) DeleteEvent(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllEventsByUser provides a mock function with given fields: ctx, userID
func (_m *Model) GetAllEventsByUser(ctx context.Context, userID int64) ([]*entities.Event, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllEventsByUser")
	}

	var r0 []*entities.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*entities.Event, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*entities.Event); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvailabilitiesByEvent provides a mock function with given fields: ctx, eventID
func (_m *Model) GetAvailabilitiesByEvent(ctx context.Context, eventID int64) ([]entities.ParticipantAvailability, error) {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for GetAvailabilitiesByEvent")
	}

	var r0 []entities.ParticipantAvailability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entities.ParticipantAvailability, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entities.ParticipantAvailability); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ParticipantAvailability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvailabilityByID provides a mock function with given fields: ctx, id
func (_m *Model) GetAvailabilityByID(ctx context.Context, id int64) (*entities.ParticipantAvailability, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAvailabilityByID")
	}

	var r0 *entities.ParticipantAvailability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entities.ParticipantAvailability, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entities.ParticipantAvailability); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ParticipantAvailability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventByID provides a mock function with given fields: ctx, id
func (_m *Model) GetEventByID(ctx context.Context, id int64) (*entities.Event, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetEventByID")
	}

	var r0 *entities.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entities.Event, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entities.Event); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
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

// UpdateAvailability provides a mock function with given fields: ctx, availability
func (_m *Model) UpdateAvailability(ctx context.Context, availability *entities.ParticipantAvailability) error {
	ret := _m.Called(ctx, availability)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAvailability")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParticipantAvailability) error); ok {
		r0 = rf(ctx, availability)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEvent provides a mock function with given fields: ctx, event
func (_m *Model) UpdateEvent(ctx context.Context, event *entities.Event) error {
	ret := _m.Called(ctx, event)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Event) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
