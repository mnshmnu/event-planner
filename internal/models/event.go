package models

import (
	"context"
	"event-planner/internal/entities"
)

// CreateEvent inserts a new event
func (m *model) CreateEvent(ctx context.Context, event *entities.Event) (int64, error) {
	query := `
        INSERT INTO events (title, created_by, duration_minutes, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	err := m.db.QueryRow(ctx, query, event.Title, event.CreatedBy, event.DurationMinutes, event.CreatedAt).
		Scan(&event.ID)
	return event.ID, err
}

// GetEventByID fetches a single event by ID
func (m *model) GetEventByID(ctx context.Context, id int64) (*entities.Event, error) {
	query := `
        SELECT id, title, created_by, duration_minutes, created_at
        FROM events
        WHERE id = $1;
    `
	var e entities.Event
	err := m.db.QueryRow(ctx, query, id).Scan(&e.ID, &e.Title, &e.CreatedBy, &e.DurationMinutes, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// UpdateEvent updates an event's title or duration
func (m *model) UpdateEvent(ctx context.Context, event *entities.Event) error {
	query := `
        UPDATE events
        SET title = $1, duration_minutes = $2
        WHERE id = $3;
    `
	_, err := m.db.Exec(ctx, query, event.Title, event.DurationMinutes, event.ID)
	return err
}

// DeleteEvent removes an event by ID
func (m *model) DeleteEvent(ctx context.Context, id int64) error {
	query := `DELETE FROM events WHERE id = $1;`
	_, err := m.db.Exec(ctx, query, id)
	return err
}

// GetAllEventsByUser fetches all events created by a specific user
func (m *model) GetAllEventsByUser(ctx context.Context, userID int64) ([]*entities.Event, error) {
	query := `
        SELECT id, title, created_by, duration_minutes, created_at
        FROM events
        WHERE created_by = $1
        ORDER BY created_at DESC;
    `
	rows, err := m.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*entities.Event
	for rows.Next() {
		var e entities.Event
		if err := rows.Scan(&e.ID, &e.Title, &e.CreatedBy, &e.DurationMinutes, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, &e)
	}

	return events, nil
}

// CreateEventSlot inserts a new event slot
func (m *model) CreateEventSlot(ctx context.Context, slot *entities.EventSlot) (int64, error) {
	query := `
        INSERT INTO event_slots (event_id, start_time, end_time)
        VALUES ($1, $2, $3)
        RETURNING id;
    `
	err := m.db.QueryRow(ctx, query, slot.EventID, slot.StartTime, slot.EndTime).
		Scan(&slot.ID)
	return slot.ID, err
}

// CreateEventSlots inserts multiple event slots
func (m *model) CreateEventSlots(ctx context.Context, slots []entities.EventSlot) error {
	if len(slots) == 0 {
		return nil
	}

	query := `
		INSERT INTO event_slots (event_id, start_time, end_time)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	for i := range slots {
		err := m.db.QueryRow(
			ctx,
			query,
			slots[i].EventID,
			slots[i].StartTime,
			slots[i].EndTime,
		).Scan(&slots[i].ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsSlotPartOfEvent checks if a slot is part of an event
func (m *model) IsSlotPartOfEvent(ctx context.Context, slotID, eventID int64) (bool, error) {
	query := `SELECT COUNT(1) FROM event_slots WHERE id = $1 AND event_id = $2`
	var count int
	err := m.db.QueryRow(ctx, query, slotID, eventID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ConfirmFinalSlot updates the event's confirmed slot
func (m *model) ConfirmFinalSlot(ctx context.Context, eventID, slotID int64) error {
	query := `UPDATE events SET confirmed_slot_id = $1 WHERE id = $2`
	_, err := m.db.Exec(ctx, query, slotID, eventID)
	return err
}
