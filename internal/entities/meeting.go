package entities

import "time"

type Meeting struct {
	ID           int       `json:"id"`
	OrganizerID  int       `json:"organizer_id"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Participants []int     `json:"participants"`
}

type MeetingParticipant struct {
	MeetingID int `json:"meeting_id"`
	UserID    int `json:"user_id"`
}
