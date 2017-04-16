package types

import "time"

// CreateFreeEventRequest ...
type CreateFreeEventRequest struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
