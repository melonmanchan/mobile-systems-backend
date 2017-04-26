package models

import (
	_ "database/sql"
	"log"

	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

// Client ...
type Client struct {
	DB *sqlx.DB
}

// PerformPendingMigrations ...
func PerformPendingMigrations(path string, connectionString string) []error {
	errors, ok := migrate.UpSync(connectionString, path)

	if !ok {
		return errors
	}

	return nil
}

// ConnectToDatabase ...
func ConnectToDatabase(connectionString string) (*Client, error) {
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
