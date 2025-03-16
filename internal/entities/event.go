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
	CreatedBy       int64     `json:"createdBy"`
	DurationMinutes int       `json:"durationMinutes"`
	CreatedAt       time.Time `json:"createdAt"`
}
type EventSlot struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"eventID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type EventParticipant struct {
	ID      int64 `json:"id"`
	EventID int64 `json:"eventID"`
	UserID  int64 `json:"userID"`
}
