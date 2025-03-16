package entities

import "time"

type ParticipantStatus string

var (
	Invited  ParticipantStatus = "invited"
	Accepted ParticipantStatus = "accepted"
	Declined ParticipantStatus = "declined"
)

type Event struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	CreatedBy       int64     `json:"created_by"`
	DurationMinutes int       `json:"duration_minutes"`
	CreatedAt       time.Time `json:"created_at"`
}
type EventSlot struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"event_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type EventParticipant struct {
	ID      int64 `json:"id"`
	EventID int64 `json:"event_id"`
	UserID  int64 `json:"user_id"`
}
