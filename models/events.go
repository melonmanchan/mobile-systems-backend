package models

import (
	"time"

	"github.com/guregu/null"
)

// Event ...
type Event struct {
	ID        int64     `json:"id" db:"id"`
	TutorID   int64     `json:"tutor" db:"tutor"`
	TuteeID   null.Int  `json:"tutee" db:"tutee"`
	StartTime time.Time `json:"start" db:"start"`
	EndTime   time.Time `json:"end" db:"end"`
}
