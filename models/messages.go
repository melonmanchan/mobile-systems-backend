package models

import (
	"time"

	"github.com/guregu/null"
)

// Message ...
type Message struct {
	ID         int64       `json:"id" db:"id"`
	SenderID   int64       `json:"sender" db:"sender"`
	ReceiverID int64       `json:"receiver" db:"receiver"`
	Content    null.String `json:"content" db:"content"`
	SentAt     time.Time   `json:"sent_at" db:"sent_at"`
}

// CreateMessage ...
func (c Client) CreateMessage(senderID int64, receiverID int64, content string) (Message, error) {
	msg := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    null.StringFrom(content),
		SentAt:     time.Now(),
	}

	res, err := c.DB.Exec(`
	INSERT INTO messages (sender, receiver, content)
	VALUES($1, $2, $3);
	`, senderID, receiverID, content)

	msg.ID, _ = res.LastInsertId()

	return msg, err
}

// GetConversation ...
func (c Client) GetConversation(firstID int64, secondID int64) ([]Message, error) {
	messages := []Message{}

	err := c.DB.Select(&messages, `
	SELECT messages.* FROM messages
	WHERE
	(messages.sender = $1
		AND
	messages.receiver = $2)
		OR
	(messages.sender = $2
		AND
	messages.receiver = $1);`, firstID, secondID)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c Client) GetUserLatestReceivedMessages(user *User) ([]Message, error) {
	receivedMessages := []Message{}
	sentMessages := []Message{}
	out := []Message{}

	// Map indices to a message
	var uniqueMap map[[2]int64]Message
	uniqueMap = make(map[[2]int64]Message)

	err := c.DB.Select(&receivedMessages, `
	SELECT DISTINCT ON(sender) sender, id, receiver, content, sent_at
	FROM messages
	WHERE receiver = $1
	ORDER BY sender, id DESC;
	`, user.ID)

	if err != nil {
		return nil, err
	}

	err = c.DB.Select(&sentMessages, `
	SELECT DISTINCT ON(receiver) sender, id, receiver, content, sent_at
	FROM messages
	WHERE sender = $1
	ORDER BY receiver, id DESC;
	`, user.ID)

	if err != nil {
		return nil, err
	}

	allMessages := append(receivedMessages, sentMessages...)

	for _, m := range allMessages {
		sender := m.SenderID
		receiver := m.ReceiverID
		arr := [2]int64{}

		if receiver < sender {
			arr[0] = receiver
			arr[1] = sender
		} else {
			arr[0] = sender
			arr[1] = receiver
		}

		msg, exists := uniqueMap[arr]

		if exists {
			if msg.ID < m.ID {
				uniqueMap[arr] = m
			}
		} else {
			uniqueMap[arr] = m
		}
	}

	for _, val := range uniqueMap {
		out = append(out, val)
	}

	return out, nil
}
