package services

import (
	"context"
	"event-planner/internal/entities"
	"fmt"
	"time"
)

func (s *service) GetAvailableSlots(ctx context.Context, duration int) (map[time.Time][]int, error) {
	availabilityMap, err := s.models.GetAvailabilities(ctx, duration)
	if err != nil {
		return nil, err
	}

	return availabilityMap, nil
}

func (s *service) ScheduleMeeting(ctx context.Context, organizerID int, duration int) (*entities.Meeting, error) {
	availabilityMap, err := s.GetAvailableSlots(ctx, duration)
	if err != nil {
		return nil, err
	}

	var bestSlot time.Time
	maxParticipants := 0
	selectedParticipants := []int{}

	for slot, participants := range availabilityMap {
		if len(participants) > maxParticipants {
			bestSlot = slot
			maxParticipants = len(participants)
			selectedParticipants = participants
		}
	}

	if maxParticipants == 0 {
		return nil, fmt.Errorf("no suitable meeting slot found")
	}

	meeting := &entities.Meeting{
		OrganizerID:  organizerID,
		StartTime:    bestSlot,
		EndTime:      bestSlot.Add(time.Minute * time.Duration(duration)),
		Participants: selectedParticipants,
	}

	err = s.models.CreateMeeting(ctx, meeting)
	if err != nil {
		return nil, err
	}

	return meeting, nil
}
