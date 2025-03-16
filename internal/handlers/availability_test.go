package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"event-planner/internal/entities"
	svcMock "event-planner/internal/services/mocks"
	"event-planner/pkg/middlewares"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func generateRequest(method string, url string, body interface{}) *http.Request {
	currUserClaims := entities.UserClaims{
		UserID: 1,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalln(err)
	}
	// return req

	ctx := context.WithValue(req.Context(), middlewares.JWTContextKey, currUserClaims)
	return req.WithContext(ctx)
}

func Test_handler_CreateAvailability(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2021-01-01T01:00:00Z")

	type testCase struct {
		name             string
		requestBody      interface{}
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
		rawBody          []byte // For invalid JSON test
	}

	tests := []testCase{
		{
			name: "successful availability creation",
			requestBody: map[string]interface{}{
				"eventID":   1,
				"userID":    1,
				"startTime": "2021-01-01T00:00:00Z",
				"endTime":   "2021-01-01T01:00:00Z",
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("CreateAvailability", mock.Anything, &entities.ParticipantAvailability{
					EventID:   1,
					UserID:    1,
					StartTime: startTime,
					EndTime:   endTime,
				}).Return(int64(1), nil).Once()
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `"availability_id":1`,
		},
		{
			name:             "invalid request body",
			rawBody:          []byte("invalid-json"),
			mockSetup:        func(s *svcMock.Service) {}, // no service call expected
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid request",
		},
		{
			name: "service layer error",
			requestBody: map[string]interface{}{
				"eventID":   2,
				"userID":    1,
				"startTime": "2021-01-01T00:00:00Z",
				"endTime":   "2021-01-01T01:00:00Z",
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("CreateAvailability", mock.Anything, mock.Anything).Return(int64(0), errors.New("event not found")).Once()
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "event not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := new(svcMock.Service)
			if tt.mockSetup != nil {
				tt.mockSetup(mockSvc)
			}

			h := &handler{service: mockSvc}
			w := httptest.NewRecorder()

			var req *http.Request
			if tt.rawBody != nil {
				req, _ = http.NewRequest("POST", "/availability", bytes.NewBuffer(tt.rawBody))
				ctx := context.WithValue(req.Context(), middlewares.JWTContextKey, entities.UserClaims{UserID: 1})
				req = req.WithContext(ctx)
			} else {
				req = generateRequest("POST", "/availability", tt.requestBody)
			}

			h.CreateAvailability(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)

			mockSvc.AssertExpectations(t)
		})
	}
}

func Test_handler_GetAvailability(t *testing.T) {
	type testCase struct {
		name             string
		queryParam       string
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
	}

	tests := []testCase{
		{
			name:       "successfully fetched availability",
			queryParam: "?id=1",
			mockSetup: func(s *svcMock.Service) {
				s.On("GetAvailabilityByID", mock.Anything, int64(1)).Return(&entities.ParticipantAvailability{
					ID:        1,
					EventID:   2,
					UserID:    3,
					StartTime: time.Date(2025, 3, 16, 23, 46, 25, 301344000, time.FixedZone("IST", 5*60*60+30*60)),
					EndTime:   time.Date(2025, 3, 17, 0, 46, 25, 301344000, time.FixedZone("IST", 5*60*60+30*60)),
				}, nil).Once()
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":1,"eventID":2,"userID":3,"startTime":"2025-03-16T23:46:25.301344+05:30","endTime":"2025-03-17T00:46:25.301344+05:30"}`,
		},
		{
			name:             "invalid id in query param",
			queryParam:       "?id=abc",
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid availability ID",
		},
		{
			name:       "availability not found",
			queryParam: "?id=5",
			mockSetup: func(s *svcMock.Service) {
				s.On("GetAvailabilityByID", mock.Anything, int64(5)).Return(nil, errors.New("not found")).Once()
			},
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "Availability not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := new(svcMock.Service)
			tt.mockSetup(svc)

			h := &handler{service: svc}
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/availability"+tt.queryParam, nil)

			h.GetAvailability(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
			svc.AssertExpectations(t)
		})
	}
}

func Test_handler_UpdateAvailability(t *testing.T) {
	startTime := time.Now()
	endTime := startTime.Add(1 * time.Hour)

	type testCase struct {
		name             string
		queryParam       string
		requestBody      interface{}
		rawBody          []byte
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
	}

	tests := []testCase{
		{
			name:       "successfully updated availability",
			queryParam: "?id=1",
			requestBody: map[string]interface{}{
				"startTime": startTime.Format(time.RFC3339),
				"endTime":   endTime.Format(time.RFC3339),
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("GetAvailabilityByID", mock.Anything, int64(1)).Return(&entities.ParticipantAvailability{
					ID: 1,
				}, nil).Once()
				s.On("UpdateAvailability", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `"status":"updated"`,
		},
		{
			name:             "invalid id in query param",
			queryParam:       "?id=abc",
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid availability ID",
		},
		{
			name:             "invalid request body",
			queryParam:       "?id=1",
			rawBody:          []byte("not-json"),
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid request body",
		},
		{
			name:       "availability not found",
			queryParam: "?id=2",
			requestBody: map[string]interface{}{
				"startTime": startTime.Format(time.RFC3339),
				"endTime":   endTime.Format(time.RFC3339),
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("GetAvailabilityByID", mock.Anything, int64(2)).Return(nil, errors.New("not found")).Once()
			},
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "Availability not found",
		},
		{
			name:       "update failed",
			queryParam: "?id=3",
			requestBody: map[string]interface{}{
				"startTime": startTime.Format(time.RFC3339),
				"endTime":   endTime.Format(time.RFC3339),
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("GetAvailabilityByID", mock.Anything, int64(3)).Return(&entities.ParticipantAvailability{ID: 3}, nil).Once()
				s.On("UpdateAvailability", mock.Anything, mock.Anything).Return(errors.New("update error")).Once()
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "update error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := new(svcMock.Service)
			tt.mockSetup(svc)

			h := &handler{service: svc}
			w := httptest.NewRecorder()

			var req *http.Request
			if tt.rawBody != nil {
				req = httptest.NewRequest("PUT", "/availability"+tt.queryParam, bytes.NewBuffer(tt.rawBody))
			} else {
				bodyBytes, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("PUT", "/availability"+tt.queryParam, bytes.NewBuffer(bodyBytes))
			}

			h.UpdateAvailability(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
			svc.AssertExpectations(t)
		})
	}
}

func Test_handler_DeleteAvailability(t *testing.T) {
	type testCase struct {
		name             string
		queryParam       string
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
	}

	tests := []testCase{
		{
			name:       "successfully deleted availability",
			queryParam: "?id=1",
			mockSetup: func(s *svcMock.Service) {
				s.On("DeleteAvailability", mock.Anything, int64(1)).Return(nil).Once()
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `"status":"deleted"`,
		},
		{
			name:             "invalid id in query param",
			queryParam:       "?id=abc",
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid availability ID",
		},
		{
			name:       "delete failed",
			queryParam: "?id=2",
			mockSetup: func(s *svcMock.Service) {
				s.On("DeleteAvailability", mock.Anything, int64(2)).Return(errors.New("delete error")).Once()
			},
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: "Failed to delete availability",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := new(svcMock.Service)
			tt.mockSetup(svc)

			h := &handler{service: svc}
			w := httptest.NewRecorder()

			req := httptest.NewRequest("DELETE", "/availability"+tt.queryParam, nil)
			h.DeleteAvailability(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
			svc.AssertExpectations(t)
		})
	}
}
