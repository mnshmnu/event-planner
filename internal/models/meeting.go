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

func (m *model) GetMeetings(ctx context.Context) ([]*entities.Meeting, error) {
	var meetings []*entities.Meeting
	rows, err := m.db.Query(ctx, `
		SELECT m.id, m.organizer_id, m.start_time, m.end_time, array_agg(mp.user_id) as participants
		FROM meetings m
		LEFT JOIN meeting_participants mp ON m.id = mp.meeting_id
		GROUP BY m.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var meeting entities.Meeting
		var participants []int
		if err := rows.Scan(&meeting.ID, &meeting.OrganizerID, &meeting.StartTime, &meeting.EndTime, &participants); err != nil {
			return nil, err
		}
		meeting.Participants = participants
		meetings = append(meetings, &meeting)
	}

	return meetings, nil
}
