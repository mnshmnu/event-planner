// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Handlers is an autogenerated mock type for the Handlers type
type Handlers struct {
	mock.Mock
}

// CreateAvailability provides a mock function with given fields: w, r
func (_m *Handlers) CreateAvailability(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// CreateEvent provides a mock function with given fields: w, r
func (_m *Handlers) CreateEvent(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// DeleteAvailability provides a mock function with given fields: w, r
func (_m *Handlers) DeleteAvailability(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// DeleteEvent provides a mock function with given fields: w, r
func (_m *Handlers) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetAllEventsByUser provides a mock function with given fields: w, r
func (_m *Handlers) GetAllEventsByUser(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetAvailability provides a mock function with given fields: w, r
func (_m *Handlers) GetAvailability(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetEventByID provides a mock function with given fields: w, r
func (_m *Handlers) GetEventByID(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Login provides a mock function with given fields: w, r
func (_m *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Register provides a mock function with given fields: w, r
func (_m *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UpdateAvailability provides a mock function with given fields: w, r
func (_m *Handlers) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UpdateEvent provides a mock function with given fields: w, r
func (_m *Handlers) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewHandlers creates a new instance of Handlers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlers(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handlers {
	mock := &Handlers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
