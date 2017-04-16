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
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
}

// CreateNewFreeEvent ...
func (c Client) CreateNewFreeEvent(user *User, start time.Time, end time.Time) (Event, error) {
	event := Event{
		TutorID:   user.ID,
		StartTime: start,
		EndTime:   end,
	}
	res, err := c.DB.Exec(`
	INSERT INTO events (tutor, start_time, end_time)
	VALUES($1, $2, $3);

	`, user.ID, start, end)

	if err != nil {
		return event, err

	}

	event.ID, _ = res.LastInsertId()
	return event, nil

}
