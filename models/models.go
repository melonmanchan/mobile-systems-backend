package models

import (
	"github.com/jmoiron/sqlx"
)

// Client ...
type Client struct {
	DB *sqlx.DB
}

// Datastore ...
type Datastore interface {
	GetUser() (*string, error)
}

// ConnectToDatabase ...
func ConnectToDatabase() (*Client, error) {
	return &Client{DB: nil}, nil
}

// Close ...
func (c *Client) Close() {
	c.DB.Close()
}
