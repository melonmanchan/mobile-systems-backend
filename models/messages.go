package models

import (
	"time"

	"github.com/guregu/null"
)

// Message ...
type Message struct {
	ID         int64       `json:"id" db:"id"`
	SenderID   int64       `json:"sender" db:"sender"`
	ReceiverID int64       `json":receiver" db:"receiver"`
	Content    null.String `json":content" db:"content"`
	SentAt     time.Time   `json:"sent_at" db:"sent_at"`
}

// CreateMessage ...
func (c Client) CreateMessage(senderID int64, receiverID int64, content string) error {
	_, err := c.DB.Exec(`
	INSERT INTO messages (sender, receiver, content)
	VALUES($1, $2, $3);
	`, senderID, receiverID, content)

	return err
}

// GetConversation ...
func (c Client) GetConversation(firstID int64, secondID int64) ([]Message, error) {
	messages := []Message{}

	err := c.DB.Select(&messages, `
	SELECT messages.* FROM messages
	WHERE messages.sender = $1
	OR messages.receiver = $1
	OR messages.sender = $2
	OR messages.receiver = $2;`, firstID, secondID)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c Client) GetUserLatestReceivedMessages(user *User) ([]Message, error) {
	messages := []Message{}

	err := c.DB.Select(&messages, `
	SELECT distinct sender, id, receiver, content, sent_at
	FROM messages
	WHERE messages.receiver = $1
	ORDER BY messages.sent_at;`, firstID, secondID)

	if err != nil {
		return nil, err
	}
	return messages, nil
}
