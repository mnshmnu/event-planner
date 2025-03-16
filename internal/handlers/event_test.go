package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"event-planner/internal/entities"
	"event-planner/internal/entities/packet"
	svcMock "event-planner/internal/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// helper for creating a new handler with mocked service
func getTestHandlerWithMockService(mockSetup func(s *svcMock.Service)) (*handler, *svcMock.Service) {
	svc := new(svcMock.Service)
	mockSetup(svc)
	return &handler{service: svc}, svc
}

// ------------------- CreateEvent -------------------

func Test_handler_CreateEvent(t *testing.T) {
	type testCase struct {
		name             string
		requestBody      *packet.CreateEventReq
		mockSetup        func(s *svcMock.Service)
		expectedStatus   int
		expectedResponse string
		rawBody          []byte // raw JSON body for invalid json test
	}

	tests := []testCase{
		{
			name: "successful event creation",
			requestBody: &packet.CreateEventReq{
				Title:           "Team Sync",
				DurationMinutes: 30,
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("CreateEventWithSlots", mock.Anything, mock.Anything).
					Return(&packet.CreateEventResp{EventID: 10}, nil).
					Once()
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"eventID":10,"slotStats":null,"message":""}`,
		},
		{
			name:             "invalid request body",
			requestBody:      &packet.CreateEventReq{}, // Simulate missing fields or malformed JSON manually
			mockSetup:        func(s *svcMock.Service) {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid request",
			rawBody:          []byte("invalid"),
		},
		{
			name: "internal error from service",
			requestBody: &packet.CreateEventReq{
				Title: "Team Sync", DurationMinutes: 30,
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("CreateEventWithSlots", mock.Anything, mock.Anything).
					Return(&packet.CreateEventResp{}, errors.New("db error")).Once()
			},
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: "Failed to create event",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			body, _ := json.Marshal(tt.requestBody)

			var req *http.Request

			if tt.rawBody != nil {
				req, _ = http.NewRequest("POST", "/event", bytes.NewBuffer(tt.rawBody))
			} else {
				req = httptest.NewRequest("POST", "/event", bytes.NewReader(body))

			}

			w := httptest.NewRecorder()

			h.CreateEvent(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
			svc.AssertExpectations(t)
		})
	}
}

// ------------------- GetEventByID -------------------

func Test_handler_GetEventByID(t *testing.T) {
	type testCase struct {
		name           string
		queryParam     string
		mockSetup      func(s *svcMock.Service)
		expectedStatus int
		expectedResp   string
	}

	tests := []testCase{
		{
			name:       "event found",
			queryParam: "?id=1",
			mockSetup: func(s *svcMock.Service) {
				s.On("GetEventByID", mock.Anything, int64(1)).
					Return(&entities.Event{ID: 1, Title: "Test Event"}, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedResp:   `"id":1`,
		},
		{
			name:           "invalid id",
			queryParam:     "?id=abc",
			mockSetup:      func(s *svcMock.Service) {},
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "Invalid event id",
		},
		{
			name:       "event not found",
			queryParam: "?id=5",
			mockSetup: func(s *svcMock.Service) {
				s.On("GetEventByID", mock.Anything, int64(5)).Return(nil, errors.New("not found")).Once()
			},
			expectedStatus: http.StatusNotFound,
			expectedResp:   "Event not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			req := httptest.NewRequest("GET", "/event"+tt.queryParam, nil)
			w := httptest.NewRecorder()

			h.GetEventByID(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResp)
			svc.AssertExpectations(t)
		})
	}
}

// ------------------- UpdateEvent -------------------

func Test_handler_UpdateEvent(t *testing.T) {
	type testCase struct {
		name           string
		input          packet.UpdateEventReq
		mockSetup      func(s *svcMock.Service)
		expectedStatus int
		expectedResp   string
		rawBody        []byte // raw JSON body for invalid json test
	}

	tests := []testCase{
		{
			name: "successful update",
			input: packet.UpdateEventReq{
				ID:              1,
				Title:           "Updated Event",
				DurationMinutes: 60,
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("UpdateEvent", mock.Anything, &entities.Event{
					ID:              1,
					Title:           "Updated Event",
					DurationMinutes: 60,
				}).Return(nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedResp:   "success",
		},
		{
			name:           "decode failure",
			input:          packet.UpdateEventReq{},
			mockSetup:      func(s *svcMock.Service) {},
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "Invalid request",
			rawBody:        []byte("invalid"),
		},
		{
			name: "service failure",
			input: packet.UpdateEventReq{
				ID: 2, Title: "Fail", DurationMinutes: 45,
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("UpdateEvent", mock.Anything, mock.Anything).
					Return(errors.New("fail")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResp:   "Failed to update event",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			body, _ := json.Marshal(tt.input)

			var req *http.Request
			if tt.rawBody != nil {
				req, _ = http.NewRequest("PUT", "/event", bytes.NewBuffer(tt.rawBody))
			} else {
				req = httptest.NewRequest("PUT", "/event", bytes.NewReader(body))
			}

			w := httptest.NewRecorder()

			h.UpdateEvent(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResp)
			svc.AssertExpectations(t)
		})
	}
}

// ------------------- DeleteEvent -------------------

func Test_handler_DeleteEvent(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		mockSetup      func(s *svcMock.Service)
		expectedStatus int
		expectedResp   string
		rawBody        []byte // raw JSON body for invalid json test
	}{
		{
			name:       "successful delete",
			queryParam: "?id=1",
			mockSetup: func(s *svcMock.Service) {
				s.On("DeleteEvent", mock.Anything, int64(1)).Return(nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedResp:   "deleted",
		},
		{
			name:           "invalid id",
			queryParam:     "?id=x",
			mockSetup:      func(s *svcMock.Service) {},
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "Invalid event id",
		},
		{
			name:       "service error",
			queryParam: "?id=10",
			mockSetup: func(s *svcMock.Service) {
				s.On("DeleteEvent", mock.Anything, int64(10)).Return(errors.New("fail")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResp:   "Failed to delete event",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			req := httptest.NewRequest("DELETE", "/event"+tt.queryParam, nil)
			w := httptest.NewRecorder()

			h.DeleteEvent(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResp)
			svc.AssertExpectations(t)
		})
	}
}

// ------------------- GetAllEventsByUser -------------------

func Test_handler_GetAllEventsByUser(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		mockSetup      func(s *svcMock.Service)
		expectedStatus int
		expectedResp   string
	}{
		{
			name:       "success",
			queryParam: "?user_id=2",
			mockSetup: func(s *svcMock.Service) {
				s.On("GetEventsByUser", mock.Anything, int64(2)).
					Return([]*entities.Event{{ID: 1, Title: "Event 1"}}, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedResp:   `"id":1`,
		},
		{
			name:           "invalid user id",
			queryParam:     "?user_id=abc",
			mockSetup:      func(s *svcMock.Service) {},
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "Invalid user id",
		},
		{
			name:       "service error",
			queryParam: "?user_id=5",
			mockSetup: func(s *svcMock.Service) {
				s.On("GetEventsByUser", mock.Anything, int64(5)).Return(nil, errors.New("fail")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResp:   "Failed to fetch events",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			req := httptest.NewRequest("GET", "/events"+tt.queryParam, nil)
			w := httptest.NewRecorder()

			h.GetAllEventsByUser(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResp)
			svc.AssertExpectations(t)
		})
	}
}

// ------------------- ConfirmFinalSlot -------------------

func Test_handler_ConfirmFinalSlot(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    packet.ConfirmSlotReq
		mockSetup      func(s *svcMock.Service)
		expectedStatus int
		expectedResp   string
		rawBody        []byte // raw JSON body for invalid json test
	}{
		{
			name: "success",
			requestBody: packet.ConfirmSlotReq{
				EventID: 1,
				SlotID:  5,
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("ConfirmFinalSlot", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedResp:   "confirmed",
		},
		{
			name:           "decode error",
			requestBody:    packet.ConfirmSlotReq{},
			mockSetup:      func(s *svcMock.Service) {},
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "Invalid request body",
			rawBody:        []byte("invalid"),
		},
		{
			name: "service failure",
			requestBody: packet.ConfirmSlotReq{
				EventID: 1,
				SlotID:  10,
			},
			mockSetup: func(s *svcMock.Service) {
				s.On("ConfirmFinalSlot", mock.Anything, mock.Anything).
					Return(errors.New("something went wrong")).Once()
			},
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "something went wrong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, svc := getTestHandlerWithMockService(tt.mockSetup)
			body, _ := json.Marshal(tt.requestBody)

			var req *http.Request
			w := httptest.NewRecorder()

			if tt.rawBody != nil {
				req, _ = http.NewRequest("POST", "/confirm", bytes.NewBuffer(tt.rawBody))
			} else {
				req = httptest.NewRequest("POST", "/confirm", bytes.NewReader(body))
			}

			h.ConfirmFinalSlot(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedResp)
			svc.AssertExpectations(t)
		})
	}
}
