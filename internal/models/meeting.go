package models

import (
	"context"
	"event-planner/internal/entities"
)

func (m *model) CreateMeeting(ctx context.Context, meeting *entities.Meeting) error {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var meetingID int
	err = tx.
		QueryRow(ctx, "INSERT INTO meetings (organizer_id, start_time, end_time) VALUES ($1, $2, $3) RETURNING id",
			meeting.OrganizerID, meeting.StartTime, meeting.EndTime,
		).
		Scan(&meetingID)
	if err != nil {
		return err
	}

	for _, participantID := range meeting.Participants {
		_, err := tx.
			Exec(ctx, "INSERT INTO meeting_participants (meeting_id, user_id) VALUES ($1, $2)",
				meetingID, participantID,
			)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
