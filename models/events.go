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
func (c Client) CreateNewFreeEvent(user *User, start time.Time) (Event, error) {
	if user.UserType != TutorType {
		return Event{}, errors.New("user is not a tutor")
	}

	end := start.UTC().Add(1 * time.Hour)

	event := Event{
		TutorID:   user.ID,
		StartTime: start,
		EndTime:   end,
	}

	lastInsertID := int64(0)

	err := c.DB.QueryRow(`
	INSERT INTO events (tutor, start_time, end_time)
	VALUES($1, $2, $3) RETURNING id;
	`, user.ID, start, end).Scan(&lastInsertID)

	if lastInsertID != 0 {
		event.ID = lastInsertID
	}

	if err != nil {
		return event, err
	}

	return event, nil
}

// GetTutorOwnTimes ...
func (c Client) GetTutorOwnTimes(user *User) ([]Event, error) {
	events := []Event{}

	err := c.DB.Select(&events, `
	SELECT events.* FROM events
	WHERE events.tutor = $1;
	`, user.ID)

	if err != nil {
		return events, err
	}

	return events, nil
}

// RemoveTime ...
func (c Client) RemoveTime(user *User, event *Event) error {
	if user.UserType != TutorType {
		return errors.New("user is not a tutor")
	}

	_, err := c.DB.Exec(`
		DELETE FROM EVENTS
		WHERE events.start_time = $1 AND events.end_time = $2 AND events.tutor = $3;
	`, event.StartTime, event.EndTime, user.ID)

	return err
}

// GetTutorFreeTimes ...
func (c Client) GetTutorFreeTimes(tutorID int64) ([]Event, error) {
	events := []Event{}

	err := c.DB.Select(&events, `
	SELECT events.* FROM events
	WHERE events.tutor = $1 AND events.tutee IS NULL;
	`, tutorID)

	if err != nil {
		return events, err
	}

	return events, nil
}

// ReserveTimeForUser ...
func (c Client) ReserveTimeForUser(user *User, event *Event) error {
	if event.TuteeID.Valid {
		return errors.New("event is already reserver")
	}

	_, err := c.DB.Exec(`
		UPDATE events
		SET tutee = $1 WHERE events.id = $2;
	`, user.ID, event.ID)

	if err != nil {
		return err
	}

	event.TuteeID = null.IntFrom(user.ID)

	return nil
}

// GetTuteeTimes ...
func (c Client) GetTuteeTimes(user *User) ([]Event, error) {
	events := []Event{}

	err := c.DB.Select(&events, `
	SELECT events.* FROM events
	WHERE events.tutee IS NOT NULL AND events.tutee = $1;
	`, user.ID)

	if err != nil {
		return events, err
	}

	return events, nil
}

// GetEventByID ...
func (c Client) GetEventByID(ID int64) (Event, error) {
	ev := Event{}

	err := c.DB.Get(&ev, `
	SELECT events.* FROM events
	WHERE events.id = $1;
	`, ID)

	return ev, err
}
