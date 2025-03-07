package models

import (
	"context"
	"event-planner/internal/entities"
	"time"
)

func (m *model) AddAvailability(ctx context.Context, availability *entities.Availability) error {
	_, err := m.db.Exec(ctx,
		"INSERT INTO availabilities (user_id, start_time, end_time) VALUES ($1, $2, $3)",
		availability.UserID, availability.StartTime, availability.EndTime,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *model) GetAvailabilities(ctx context.Context, duration int) (map[time.Time][]int, error) {
	rows, err := m.db.Query(ctx, "SELECT user_id, start_time, end_time FROM availabilities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	availabilityMap := make(map[time.Time][]int)

	for rows.Next() {
		var avail entities.Availability
		if err := rows.Scan(&avail.UserID, &avail.StartTime, &avail.EndTime); err != nil {
			return nil, err
		}
		t := avail.StartTime
		for t.Add(time.Minute*time.Duration(duration)).Before(avail.EndTime) || t.Add(time.Minute*time.Duration(duration)).Equal(avail.EndTime) {
			availabilityMap[t] = append(availabilityMap[t], avail.UserID)
			t = t.Add(time.Minute * time.Duration(30))
		}
	}
	return availabilityMap, nil
}
