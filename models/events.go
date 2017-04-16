package models

import (
	"time"

	"github.com/guregu/null"
)

// Message ...
type Event struct {
	ID         int64         `json:"id" db:"id"`
	TutorID	   int64         `json:"tutor" db: "tutor"`
	TuteeID    null.int64    `json: "tutee" db: "tutee"`
	StartTime  time.Time     `json:"start" db:"start"`
	EndTime    time.Time     `json:"end" db:"end"`
}


// CreateMessage ...
func (c Client) SetFreeTime ( TutorID int64, StartTime time.Time, EndTime time.Time) (Message, error) {
	ev := Event{
		TutorID:   TutorID,
		TuteeID:   null.From,
		StartTime: StartTime,
		EndTime:    EndTime,
		SentAt:     time.Now(),
	}

	res, err := c.DB.Exec(`
	INSERT INTO messages (sender, receiver, content)
	VALUES($1, $2, $3);
	`, senderID, receiverID, content)

	msg.ID, _ = res.LastInsertId()

	return msg, err
}
