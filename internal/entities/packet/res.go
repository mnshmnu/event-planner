package packet

import "time"

type EventSuggestedSlot struct {
	EventID            int64
	SlotStartTime      time.Time
	SlotEndTime        time.Time
	AvailableUserCount int
}

type CreateEventResp struct {
	EventID   int64                  `json:"eventID"`
	SlotStats []SlotAvailabilityStat `json:"slotStats"`
	Message   string                 `json:"message"`
}

type SlotAvailabilityStat struct {
	SlotID                 int64     `json:"slotID"`
	StartTime              time.Time `json:"startTime"`
	EndTime                time.Time `json:"endTime"`
	ParticipantCount       int       `json:"participantCount"`
	PossibleParticipantIDs []int64   `json:"possibleParticipantIDs"`
}
