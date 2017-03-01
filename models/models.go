package models

import (
	"github.com/jmoiron/sqlx"
)

// Client ...
type Client struct {
	DB *sqlx.DB
}

// ConnectToDatabase ...
func ConnectToDatabase() (*Client, error) {
	return &Client{DB: nil}, nil
}

// Close ...
func (c *Client) Close() {
	c.DB.Close()
}
