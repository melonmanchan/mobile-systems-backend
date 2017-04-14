package models

import (
	"time"

	"github.com/guregu/null"
)

// Message ...
type Message struct {
	ID         int64       `json:"id" db:"id"`
	SenderID   int64       `json:"sender_id" db:"sender_id"`
	ReceiverID int64       `json":receiver_id" db:"receiver_id"`
	Content    null.String `json":content" db:"content"`
	SentAt     time.Time   `json:"sent_at" db:"sent_at"`
}
