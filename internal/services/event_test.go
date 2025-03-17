package services

import (
	"context"
	"errors"
	"event-planner/internal/entities"
	"event-planner/internal/entities/packet"
	"event-planner/internal/models/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateEventWithSlots(t *testing.T) {
	type testCase struct {
		name         string
		req          packet.CreateEventReq
		mockSetup    func(m *mocks.Model, s *service)
		expectedResp *packet.CreateEventResp
		expectedErr  error
	}

	now := time.Now()
	slot1 := entities.EventSlot{StartTime: now.Add(1 * time.Hour), EndTime: now.Add(2 * time.Hour)}
	slot2 := entities.EventSlot{StartTime: now.Add(3 * time.Hour), EndTime: now.Add(4 * time.Hour)}

	tests := []testCase{
		{
			name: "success with slot availability stats",
			req: packet.CreateEventReq{
				Title:           "Team Sync",
				CreatedBy:       101,
				DurationMinutes: 60,
				Slots:           []entities.EventSlot{slot1, slot2},
			},
			mockSetup: func(m *mocks.Model, s *service) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*entities.Event")).
					Return(int64(1001), nil)

				m.On("CreateEventSlots", mock.Anything, mock.AnythingOfType("[]entities.EventSlot")).
					Return(nil)
				m.On("GetAvailabilitiesByEvent", mock.Anything, int64(1001)).
					Return([]entities.ParticipantAvailability{
						{UserID: 201, StartTime: slot1.StartTime, EndTime: slot1.EndTime},
						{UserID: 202, StartTime: slot1.StartTime, EndTime: slot1.EndTime},
						{UserID: 203, StartTime: slot2.StartTime, EndTime: slot2.EndTime},
					}, nil)
			},
			expectedResp: &packet.CreateEventResp{
				EventID: 1001,
				SlotStats: []packet.SlotAvailabilityStat{
					{
						StartTime:              slot1.StartTime,
						EndTime:                slot1.EndTime,
						ParticipantCount:       2,
						PossibleParticipantIDs: []int64{201, 202},
					},
					{
						StartTime:              slot2.StartTime,
						EndTime:                slot2.EndTime,
						ParticipantCount:       1,
						PossibleParticipantIDs: []int64{203},
					},
				},
				Message: "Event created successfully. Choose the best slot based on availability.",
			},
			expectedErr: nil,
		},
		{
			name: "fail to create event",
			req: packet.CreateEventReq{
				Title:           "Bug Bash",
				CreatedBy:       202,
				DurationMinutes: 45,
				Slots:           []entities.EventSlot{slot1},
			},
			mockSetup: func(m *mocks.Model, _ *service) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*entities.Event")).
					Return(int64(0), errors.New("db error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("db error"),
		},
		{
			name: "fail to create slots",
			req: packet.CreateEventReq{
				Title:           "Design Review",
				CreatedBy:       303,
				DurationMinutes: 30,
				Slots:           []entities.EventSlot{slot1},
			},
			mockSetup: func(m *mocks.Model, _ *service) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*entities.Event")).
					Return(int64(1002), nil)

				m.On("CreateEventSlots", mock.Anything, mock.AnythingOfType("[]entities.EventSlot")).
					Return(errors.New("slot insert error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("slot insert error"),
		},
		// {
		// 	name: "error in availability evaluation - should still return response",
		// 	req: packet.CreateEventReq{

		// 		Title: "Test Event With Eval Error",
		// 		Slots: []entities.EventSlot{
		// 			slot1,
		// 			slot2,
		// 		},
		// 	},
		// 	expectedResp: &packet.CreateEventResp{
		// 		EventID: 1003,
		// 		SlotStats: []packet.SlotAvailabilityStat{
		// 			{
		// 				StartTime:              slot1.StartTime,
		// 				EndTime:                slot1.EndTime,
		// 				ParticipantCount:       0,
		// 				PossibleParticipantIDs: nil,
		// 			},
		// 			{
		// 				StartTime:              slot2.StartTime,
		// 				EndTime:                slot2.EndTime,
		// 				ParticipantCount:       2,
		// 				PossibleParticipantIDs: []int64{501, 502},
		// 			},
		// 		},
		// 		Message: "Event created successfully. Choose the best slot based on availability.",
		// 	},
		// 	mockSetup: func(m *mocks.Model, _ *service) {
		// 		m.On("CreateEvent", mock.Anything, mock.Anything).Return(int64(1003), nil)
		// 		m.On("CreateSlots", mock.Anything, int64(1003), mock.Anything).Return(nil)
		// 		m.On("GetAvailabilitiesByEvent", mock.Anything, int64(1003)).Return([]entities.ParticipantAvailability{
		// 			{
		// 				ID:        1,
		// 				EventID:   1003,
		// 				UserID:    501,
		// 				StartTime: slot2.StartTime,
		// 				EndTime:   slot2.EndTime,
		// 			},
		// 			{
		// 				ID:        2,
		// 				EventID:   1003,
		// 				UserID:    502,
		// 				StartTime: slot2.StartTime,
		// 				EndTime:   slot2.EndTime,
		// 			},
		// 		}, errors.New("simulated error in availability fetch")).Once()
		// 		m.On("CreateEventSlots", mock.Anything, mock.Anything).Return(nil)

		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}

			// Allow overriding GetAvailableParticipantsForSlot (func field)
			if tt.mockSetup != nil {
				tt.mockSetup(model, &s)
			}

			resp, err := s.CreateEventWithSlots(context.Background(), tt.req)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp.EventID, resp.EventID)
				assert.Equal(t, tt.expectedResp.Message, resp.Message)
				assert.ElementsMatch(t, tt.expectedResp.SlotStats, resp.SlotStats)
			}
		})
	}
}

func TestService_GetAvailableParticipantsForSlot(t *testing.T) {
	type testCase struct {
		name            string
		mockSetup       func(m *mocks.Model)
		expectedCount   int
		expectedUserIDs []int64
		expectedErr     error
	}

	slot := entities.EventSlot{
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	tests := []testCase{
		{
			name: "participants available",
			mockSetup: func(m *mocks.Model) {
				m.On("GetAvailabilitiesByEvent", mock.Anything, int64(1)).
					Return([]entities.ParticipantAvailability{
						{UserID: 1, StartTime: slot.StartTime.Add(-10 * time.Minute), EndTime: slot.EndTime},
						{UserID: 2, StartTime: slot.StartTime, EndTime: slot.EndTime.Add(30 * time.Minute)},
					}, nil)
			},
			expectedCount:   2,
			expectedUserIDs: []int64{1, 2},
			expectedErr:     nil,
		},
		{
			name: "model error",
			mockSetup: func(m *mocks.Model) {
				m.On("GetAvailabilitiesByEvent", mock.Anything, int64(1)).
					Return(nil, errors.New("db error"))
			},
			expectedCount:   0,
			expectedUserIDs: nil,
			expectedErr:     errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}
			tt.mockSetup(model)

			count, ids, err := s.GetAvailableParticipantsForSlot(context.Background(), 1, slot)
			assert.Equal(t, tt.expectedCount, count)
			assert.ElementsMatch(t, tt.expectedUserIDs, ids)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestService_GetEventByID(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(m *mocks.Model)
		expected    *entities.Event
		expectedErr error
	}{
		{
			name: "success",
			mockSetup: func(m *mocks.Model) {
				m.On("GetEventByID", mock.Anything, int64(1)).
					Return(&entities.Event{ID: 1, Title: "Test"}, nil)
			},
			expected:    &entities.Event{ID: 1, Title: "Test"},
			expectedErr: nil,
		},
		{
			name: "error",
			mockSetup: func(m *mocks.Model) {
				m.On("GetEventByID", mock.Anything, int64(1)).
					Return(nil, errors.New("not found"))
			},
			expected:    nil,
			expectedErr: errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}
			tt.mockSetup(model)

			e, err := s.GetEventByID(context.Background(), 1)
			assert.Equal(t, tt.expected, e)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestService_UpdateEvent(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *mocks.Model)
		expected  error
	}{
		{
			name: "success",
			mockSetup: func(m *mocks.Model) {
				m.On("UpdateEvent", mock.Anything, mock.Anything).
					Return(nil)
			},
			expected: nil,
		},
		{
			name: "error",
			mockSetup: func(m *mocks.Model) {
				m.On("UpdateEvent", mock.Anything, mock.Anything).
					Return(errors.New("update failed"))
			},
			expected: errors.New("update failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}
			tt.mockSetup(model)

			err := s.UpdateEvent(context.Background(), &entities.Event{})
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestService_DeleteEvent(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *mocks.Model)
		expected  error
	}{
		{
			name: "success",
			mockSetup: func(m *mocks.Model) {
				m.On("DeleteEvent", mock.Anything, int64(1)).
					Return(nil)
			},
			expected: nil,
		},
		{
			name: "error",
			mockSetup: func(m *mocks.Model) {
				m.On("DeleteEvent", mock.Anything, int64(1)).
					Return(errors.New("deletion failed"))
			},
			expected: errors.New("deletion failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}
			tt.mockSetup(model)

			err := s.DeleteEvent(context.Background(), 1)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestService_GetEventsByUser(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(m *mocks.Model)
		expected    []*entities.Event
		expectedErr error
	}{
		{
			name: "success",
			mockSetup: func(m *mocks.Model) {
				m.On("GetAllEventsByUser", mock.Anything, int64(1)).
					Return([]*entities.Event{{ID: 1, CreatedBy: 1}}, nil)
			},
			expected:    []*entities.Event{{ID: 1, CreatedBy: 1}},
			expectedErr: nil,
		},
		{
			name: "error",
			mockSetup: func(m *mocks.Model) {
				m.On("GetAllEventsByUser", mock.Anything, int64(1)).
					Return(nil, errors.New("db error"))
			},
			expected:    nil,
			expectedErr: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}
			tt.mockSetup(model)

			evts, err := s.GetEventsByUser(context.Background(), 1)
			assert.Equal(t, tt.expected, evts)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestService_ConfirmFinalSlot(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(m *mocks.Model)
		input       packet.ConfirmSlotReq
		expectedErr error
	}{
		{
			name:  "success",
			input: packet.ConfirmSlotReq{EventID: 1, SlotID: 2},
			mockSetup: func(m *mocks.Model) {
				m.On("IsSlotPartOfEvent", mock.Anything, int64(2), int64(1)).Return(true, nil)
				m.On("ConfirmFinalSlot", mock.Anything, int64(1), int64(2)).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:  "slot not part of event",
			input: packet.ConfirmSlotReq{EventID: 1, SlotID: 2},
			mockSetup: func(m *mocks.Model) {
				m.On("IsSlotPartOfEvent", mock.Anything, int64(2), int64(1)).Return(false, nil)
			},
			expectedErr: errors.New("slot does not belong to the given event"),
		},
		{
			name:  "db error in validation",
			input: packet.ConfirmSlotReq{EventID: 1, SlotID: 2},
			mockSetup: func(m *mocks.Model) {
				m.On("IsSlotPartOfEvent", mock.Anything, int64(2), int64(1)).Return(false, errors.New("db error"))
			},
			expectedErr: errors.New("db error"),
		},
		{
			name:  "db error on confirm",
			input: packet.ConfirmSlotReq{EventID: 1, SlotID: 2},
			mockSetup: func(m *mocks.Model) {
				m.On("IsSlotPartOfEvent", mock.Anything, int64(2), int64(1)).Return(true, nil)
				m.On("ConfirmFinalSlot", mock.Anything, int64(1), int64(2)).Return(errors.New("db error"))
			},
			expectedErr: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := new(mocks.Model)
			s := service{
				models: model,
			}
			tt.mockSetup(model)

			err := s.ConfirmFinalSlot(context.Background(), tt.input)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
