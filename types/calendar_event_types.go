package types

import "time"

// CreateEventRequest ...
type CreateFreeEventRequest struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
