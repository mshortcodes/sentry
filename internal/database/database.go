package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	db *sql.DB
}

func NewClient(dbPath string) (Client, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return Client{}, err
	}

	c := Client{db}
	if err := c.migrate(); err != nil {
		return Client{}, err
	}

	return c, nil
}

func (c *Client) migrate() error {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	password TEXT NOT NULL
	);
	`

	_, err := c.db.Exec(usersTable)
	if err != nil {
		return err
	}

	return nil
}
