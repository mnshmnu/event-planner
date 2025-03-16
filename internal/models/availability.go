package models

import (
	"context"
	"event-planner/internal/entities"
)

func (m *model) CreateAvailability(ctx context.Context, availability *entities.ParticipantAvailability) (int64, error) {
	query := `
        INSERT INTO participant_availabilities (event_id, user_id, start_time, end_time)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	err := m.db.QueryRow(ctx, query, availability.EventID, availability.UserID, availability.StartTime, availability.EndTime).
		Scan(&availability.ID)
	return availability.ID, err
}

// GetAvailabilityByID fetches a single availability by ID
func (m *model) GetAvailabilityByID(ctx context.Context, id int64) (*entities.ParticipantAvailability, error) {
	query := `
        SELECT id, event_id, user_id, start_time, end_time
        FROM participant_availabilities
        WHERE id = $1;
    `
	var a entities.ParticipantAvailability
	err := m.db.QueryRow(ctx, query, id).Scan(&a.ID, &a.EventID, &a.UserID, &a.StartTime, &a.EndTime)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// UpdateAvailability updates start and end time of an availability
func (m *model) UpdateAvailability(ctx context.Context, availability *entities.ParticipantAvailability) error {
	query := `
        UPDATE participant_availabilities
        SET start_time = $1, end_time = $2
        WHERE id = $3;
    `
	_, err := m.db.Exec(ctx, query, availability.StartTime, availability.EndTime, availability.ID)
	return err
}

// DeleteAvailability removes an availability entry
func (m *model) DeleteAvailability(ctx context.Context, id int64) error {
	query := `DELETE FROM participant_availabilities WHERE id = $1;`
	_, err := m.db.Exec(ctx, query, id)
	return err
}

// GetAvailabilitiesByEvent fetches all availabilities for a given event
func (m *model) GetAvailabilitiesByEvent(ctx context.Context, eventID int64) ([]entities.ParticipantAvailability, error) {
	query := `
        SELECT id, event_id, user_id, start_time, end_time
        FROM participant_availabilities
        WHERE event_id = $1;
    `
	rows, err := m.db.Query(ctx, query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var availabilities []entities.ParticipantAvailability
	for rows.Next() {
		var avail entities.ParticipantAvailability
		if err := rows.Scan(&avail.ID, &avail.EventID, &avail.UserID, &avail.StartTime, &avail.EndTime); err != nil {
			return nil, err
		}
		availabilities = append(availabilities, avail)
	}

	return availabilities, nil
}
