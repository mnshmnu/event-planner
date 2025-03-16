package services

import (
	"context"
	"errors"
	"event-planner/internal/entities"
)

// CreateAvailability creates a new availability for a participant
func (s *service) CreateAvailability(ctx context.Context, a *entities.ParticipantAvailability) (int64, error) {
	// Validation: End time should be after start time
	if !a.EndTime.After(a.StartTime) {
		return 0, errors.New("end time must be after start time")
	}

	return s.models.CreateAvailability(ctx, a)
}

// GetAvailabilityByID fetches a specific availability by ID
func (s *service) GetAvailabilityByID(ctx context.Context, id int64) (*entities.ParticipantAvailability, error) {
	return s.models.GetAvailabilityByID(ctx, id)
}

// UpdateAvailability updates an existing availability entry
func (s *service) UpdateAvailability(ctx context.Context, a *entities.ParticipantAvailability) error {
	// Validation: End time should be after start time
	if !a.EndTime.After(a.StartTime) {
		return errors.New("end time must be after start time")
	}

	return s.models.UpdateAvailability(ctx, a)
}

// DeleteAvailability deletes an availability entry by ID
func (s *service) DeleteAvailability(ctx context.Context, id int64) error {
	return s.models.DeleteAvailability(ctx, id)
}

// GetAvailabilitiesByEvent fetches all participant availabilities for an event
func (s *service) GetAvailabilitiesByEvent(ctx context.Context, eventID int64) ([]entities.ParticipantAvailability, error) {
	return s.models.GetAvailabilitiesByEvent(ctx, eventID)
}
