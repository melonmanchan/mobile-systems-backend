package models

import (
	"errors"
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
	if user.UserType != TutorType {
		return Event{}, errors.New("user is not a tutor")
	}

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

// GetTutorOwnTimes ...
func (c Client) GetTutorOwnTimes(user *User) ([]Event, error) {
	events := []Event{}

	if user.UserType != TutorType {
		return events, errors.New("user is not a tutor")
	}

	err := c.DB.Select(&events, `
	SELECT events.* FROM events
	WHERE events.tutor = $1;
	;`, user.ID)

	if err != nil {
		return events, err
	}

	return events, nil
}
