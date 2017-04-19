package types

import (
	"time"

	"../models"
)

// CreateFreeEventRequest ...
type CreateFreeEventRequest struct {
	StartTime time.Time `json:"start_time"`
}

// GetEventsResponse ...
type GetEventsResponse struct {
	OwnEvents      []models.Event `json:"own_events"`
	ReservedEvents []models.Event `json:"reserved_events"`
}
