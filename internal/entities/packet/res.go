package packet

import "time"

type EventSuggestedSlot struct {
	EventID            int64
	SlotStartTime      time.Time
	SlotEndTime        time.Time
	AvailableUserCount int
}

type CreateEventResp struct {
	EventID   int64                  `json:"event_id"`
	SlotStats []SlotAvailabilityStat `json:"slot_stats"`
	Message   string                 `json:"message"`
}

type SlotAvailabilityStat struct {
	SlotID                 int64     `json:"slot_id"`
	StartTime              time.Time `json:"start_time"`
	EndTime                time.Time `json:"end_time"`
	ParticipantCount       int       `json:"participant_count"`
	PossibleParticipantIDs []int64   `json:"possible_participant_ids"`
}
