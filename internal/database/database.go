package database

import (
	"database/sql"
	"fmt"

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
		username TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		password TEXT NOT NULL
	);
	`

	_, err := c.db.Exec(usersTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	apiKeys := `
	CREATE TABLE IF NOT EXISTS api_keys (
		api_key TEXT PRIMARY KEY,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_id TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id);
	)
	`

	_, err = c.db.Exec(apiKeys)
	if err != nil {
		return fmt.Errorf("failed to create api_keys table: %v", err)
	}

	return nil
}

func (c *Client) Reset() error {
	if _, err := c.db.Exec("DELETE FROM users"); err != nil {
		return fmt.Errorf("failed to reset users table: %v", err)
	}

	if _, err := c.db.Exec("DELETE FROM api_keys"); err != nil {
		return fmt.Errorf("failed to reset api_keys table: %v", err)
	}

	return nil
}
