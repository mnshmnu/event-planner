package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"event-planner/internal/entities"
	svcMock "event-planner/internal/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_handler_Register(t *testing.T) {
	mockUser := entities.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securepassword",
	}

	tests := []struct {
		name             string
		requestBody      interface{}
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:        "successful registration",
			requestBody: mockUser,
			mockSetup: func(s *svcMock.Service) {
				s.On("RegisterUser", mock.Anything, mock.AnythingOfType("*entities.User")).Return(nil).Once()
			},
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"message":"User registered successfully"}`,
		},
		{
			name:             "invalid request body",
			requestBody:      "{invalid_json}",
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid request",
		},
		{
			name:        "internal server error",
			requestBody: mockUser,
			mockSetup: func(s *svcMock.Service) {
				s.On("RegisterUser", mock.Anything, mock.AnythingOfType("*entities.User")).Return(errors.New("db error")).Once()
			},
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: "Failed to register user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			var body []byte
			if strBody, ok := tt.requestBody.(string); ok {
				body = []byte(strBody)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
			w := httptest.NewRecorder()
			h.Register(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
			svc.AssertExpectations(t)
		})
	}
}

func Test_handler_Login(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      interface{}
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "successful login",
			requestBody: map[string]string{
				"email":    "john@example.com",
				"password": "securepassword",
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("AuthenticateUser", mock.Anything, "john@example.com", "securepassword").Return("valid.jwt.token", nil).Once()
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"token":"valid.jwt.token"}`,
		},
		{
			name:             "invalid request body",
			requestBody:      "{invalid_json}",
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid request",
		},
		{
			name: "invalid credentials",
			requestBody: map[string]string{
				"email":    "john@example.com",
				"password": "wrongpassword",
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("AuthenticateUser", mock.Anything, "john@example.com", "wrongpassword").Return("", errors.New("invalid credentials")).Once()
			},
			expectedStatus:   http.StatusUnauthorized,
			expectedResponse: "Invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			var body []byte
			if strBody, ok := tt.requestBody.(string); ok {
				body = []byte(strBody)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
			w := httptest.NewRecorder()
			h.Login(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
			svc.AssertExpectations(t)
		})
	}
}
