package services

import (
	"context"
	"errors"
	"event-planner/internal/entities"
	"event-planner/internal/entities/packet"
	"time"
)

// CreateEventWithSlots creates an event with proposed time slots
func (s *service) CreateEventWithSlots(ctx context.Context, req packet.CreateEventReq) (*packet.CreateEventResp, error) {
	event := entities.Event{
		Title:           req.Title,
		CreatedBy:       req.CreatedBy,
		DurationMinutes: req.DurationMinutes,
		CreatedAt:       time.Now(),
	}

	eventID, err := s.models.CreateEvent(ctx, &event)
	if err != nil {
		return nil, err
	}

	var slots []entities.EventSlot
	for _, s := range req.Slots {
		slots = append(slots, entities.EventSlot{
			EventID:   eventID,
			StartTime: s.StartTime,
			EndTime:   s.EndTime,
		})
	}

	err = s.models.CreateEventSlots(ctx, slots)
	if err != nil {
		return nil, err
	}

	// Evaluate availability for each slot
	slotStats := make([]packet.SlotAvailabilityStat, 0)

	for _, slot := range slots {
		count, userIDs, err := s.GetAvailableParticipantsForSlot(ctx, eventID, slot)
		if err != nil {
			continue // log internally if needed
		}

		slotStats = append(slotStats, packet.SlotAvailabilityStat{
			SlotID:                 slot.ID,
			StartTime:              slot.StartTime,
			EndTime:                slot.EndTime,
			PossibleParticipantIDs: userIDs,
			ParticipantCount:       count,
		})
	}

	return &packet.CreateEventResp{
		EventID:   eventID,
		SlotStats: slotStats,
		Message:   "Event created successfully. Choose the best slot based on availability.",
	}, nil
}

func (s *service) GetAvailableParticipantsForSlot(ctx context.Context, eventID int64, slot entities.EventSlot) (int, []int64, error) {
	availabilities, err := s.models.GetAvailabilitiesByEvent(ctx, eventID)
	if err != nil {
		return 0, nil, err
	}

	userSet := make(map[int64]bool)
	for _, a := range availabilities {
		if a.StartTime.Before(slot.EndTime) && a.EndTime.After(slot.StartTime) {
			userSet[a.UserID] = true
		}
	}

	var userIDs []int64
	for id := range userSet {
		userIDs = append(userIDs, id)
	}

	return len(userIDs), userIDs, nil
}

// GetEventByID returns a single event by ID
func (s *service) GetEventByID(ctx context.Context, eventID int64) (*entities.Event, error) {
	return s.models.GetEventByID(ctx, eventID)
}

// UpdateEvent updates an event's title or duration
func (s *service) UpdateEvent(ctx context.Context, event *entities.Event) error {
	return s.models.UpdateEvent(ctx, event)
}

// DeleteEvent removes an event and its associated slots
func (s *service) DeleteEvent(ctx context.Context, eventID int64) error {
	return s.models.DeleteEvent(ctx, eventID)
}

// GetEventsByUser fetches all events created by a specific user
func (s *service) GetEventsByUser(ctx context.Context, userID int64) ([]*entities.Event, error) {
	return s.models.GetAllEventsByUser(ctx, userID)
}

// ConfirmFinalSlot confirms the final slot for an event
func (s *service) ConfirmFinalSlot(ctx context.Context, req packet.ConfirmSlotReq) error {
	// Validate that the slot belongs to the event
	valid, err := s.models.IsSlotPartOfEvent(ctx, req.SlotID, req.EventID)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("slot does not belong to the given event")
	}

	err = s.models.ConfirmFinalSlot(ctx, req.EventID, req.SlotID)
	if err != nil {
		return err
	}

	return nil
}
