package models

import (
	_ "database/sql"
	"log"

	_ "github.com/lib/pq"

	"../config"
	"github.com/jmoiron/sqlx"
)

// Client ...
type Client struct {
	DB *sqlx.DB
}

// ConnectToDatabase ...
func ConnectToDatabase(pgConf config.PostgresConfig) (*Client, error) {
	connectionString := pgConf.PostgresConfigToConnectionString()
	log.Print("Attempting to connect to " + connectionString)
	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return &Client{DB: db}, nil
}

// Close ...
func (c *Client) Close() {
	c.DB.Close()
}
