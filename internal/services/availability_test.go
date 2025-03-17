package services

import (
	"context"
	"errors"
	"event-planner/internal/entities"
	"event-planner/internal/models/mocks"

	// "event-planner/internal/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupServiceWithMockModel() (*service, *mocks.Model) {
	mockModel := new(mocks.Model)
	svc := &service{models: mockModel}
	return svc, mockModel
}

func TestService_CreateAvailability(t *testing.T) {
	svc, mockModel := setupServiceWithMockModel()
	now := time.Now()

	testCases := []struct {
		name        string
		input       *entities.ParticipantAvailability
		mockSetup   func()
		expectedID  int64
		expectError bool
	}{
		{
			name: "valid availability - success",
			input: &entities.ParticipantAvailability{
				EventID:   1,
				UserID:    2,
				StartTime: now,
				EndTime:   now.Add(1 * time.Hour),
			},
			mockSetup: func() {
				mockModel.On("CreateAvailability", mock.Anything, mock.Anything).Return(int64(101), nil).Once()
			},
			expectedID:  101,
			expectError: false,
		},
		{
			name: "invalid time range - end time before start",
			input: &entities.ParticipantAvailability{
				EventID:   1,
				UserID:    2,
				StartTime: now,
				EndTime:   now.Add(-1 * time.Hour),
			},
			mockSetup:   func() {},
			expectedID:  0,
			expectError: true,
		},
		{
			name: "model error",
			input: &entities.ParticipantAvailability{
				EventID:   1,
				UserID:    2,
				StartTime: now,
				EndTime:   now.Add(1 * time.Hour),
			},
			mockSetup: func() {
				mockModel.On("CreateAvailability", mock.Anything, mock.Anything).Return(int64(0), errors.New("db error")).Once()
			},
			expectedID:  0,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			id, err := svc.CreateAvailability(context.Background(), tc.input)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedID, id)
			}
			mockModel.AssertExpectations(t)
		})
	}
}

func TestService_GetAvailabilityByID(t *testing.T) {
	svc, mockModel := setupServiceWithMockModel()

	testCases := []struct {
		name         string
		id           int64
		mockSetup    func()
		expectedResp *entities.ParticipantAvailability
		expectError  bool
	}{
		{
			name: "successful fetch",
			id:   123,
			mockSetup: func() {
				mockModel.On("GetAvailabilityByID", mock.Anything, int64(123)).
					Return(&entities.ParticipantAvailability{ID: 123}, nil).Once()
			},
			expectedResp: &entities.ParticipantAvailability{ID: 123},
			expectError:  false,
		},
		{
			name: "model error",
			id:   456,
			mockSetup: func() {
				mockModel.On("GetAvailabilityByID", mock.Anything, int64(456)).
					Return(nil, errors.New("not found")).Once()
			},
			expectedResp: nil,
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			resp, err := svc.GetAvailabilityByID(context.Background(), tc.id)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResp, resp)
			}
			mockModel.AssertExpectations(t)
		})
	}
}

func TestService_UpdateAvailability(t *testing.T) {
	svc, mockModel := setupServiceWithMockModel()
	now := time.Now()

	testCases := []struct {
		name        string
		input       *entities.ParticipantAvailability
		mockSetup   func()
		expectError bool
	}{
		{
			name: "successful update",
			input: &entities.ParticipantAvailability{
				ID:        10,
				EventID:   1,
				UserID:    2,
				StartTime: now,
				EndTime:   now.Add(1 * time.Hour),
			},
			mockSetup: func() {
				mockModel.On("UpdateAvailability", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectError: false,
		},
		{
			name: "validation error - end before start",
			input: &entities.ParticipantAvailability{
				ID:        10,
				StartTime: now,
				EndTime:   now.Add(-1 * time.Hour),
			},
			mockSetup:   func() {},
			expectError: true,
		},
		{
			name: "model error",
			input: &entities.ParticipantAvailability{
				ID:        10,
				EventID:   1,
				UserID:    2,
				StartTime: now,
				EndTime:   now.Add(1 * time.Hour),
			},
			mockSetup: func() {
				mockModel.On("UpdateAvailability", mock.Anything, mock.Anything).Return(errors.New("db error")).Once()
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			err := svc.UpdateAvailability(context.Background(), tc.input)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockModel.AssertExpectations(t)
		})
	}
}

func TestService_DeleteAvailability(t *testing.T) {
	svc, mockModel := setupServiceWithMockModel()

	testCases := []struct {
		name        string
		id          int64
		mockSetup   func()
		expectError bool
	}{
		{
			name: "successful delete",
			id:   111,
			mockSetup: func() {
				mockModel.On("DeleteAvailability", mock.Anything, int64(111)).Return(nil).Once()
			},
			expectError: false,
		},
		{
			name: "model error",
			id:   222,
			mockSetup: func() {
				mockModel.On("DeleteAvailability", mock.Anything, int64(222)).Return(errors.New("db error")).Once()
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			err := svc.DeleteAvailability(context.Background(), tc.id)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockModel.AssertExpectations(t)
		})
	}
}

func TestService_GetAvailabilitiesByEvent(t *testing.T) {
	svc, mockModel := setupServiceWithMockModel()

	testCases := []struct {
		name         string
		eventID      int64
		mockSetup    func()
		expectedResp []entities.ParticipantAvailability
		expectError  bool
	}{
		{
			name:    "successful fetch",
			eventID: 1001,
			mockSetup: func() {
				mockModel.On("GetAvailabilitiesByEvent", mock.Anything, int64(1001)).
					Return([]entities.ParticipantAvailability{
						{ID: 1, EventID: 1001}, {ID: 2, EventID: 1001},
					}, nil).Once()
			},
			expectedResp: []entities.ParticipantAvailability{
				{ID: 1, EventID: 1001}, {ID: 2, EventID: 1001},
			},
			expectError: false,
		},
		{
			name:    "model error",
			eventID: 2002,
			mockSetup: func() {
				mockModel.On("GetAvailabilitiesByEvent", mock.Anything, int64(2002)).
					Return(nil, errors.New("db error")).Once()
			},
			expectedResp: nil,
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			resp, err := svc.GetAvailabilitiesByEvent(context.Background(), tc.eventID)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResp, resp)
			}
			mockModel.AssertExpectations(t)
		})
	}
}
