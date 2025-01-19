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
	_, err := c.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return fmt.Errorf("failed to enable foreign keys")
	}

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		password TEXT NOT NULL
	);
	`

	_, err = c.db.Exec(usersTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	passwords := `
	CREATE TABLE IF NOT EXISTS passwords (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_id TEXT NOT NULL,
		UNIQUE(user_id, name),
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	)
	`

	_, err = c.db.Exec(passwords)
	if err != nil {
		return fmt.Errorf("failed to created passwords table")
	}

	tokens := `
	CREATE TABLE IF NOT EXISTS tokens (
		token TEXT PRIMARY KEY,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_id TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`

	_, err = c.db.Exec(tokens)
	if err != nil {
		return fmt.Errorf("failed to create tokens table: %v", err)
	}

	return nil
}

func (c *Client) Reset() error {
	if _, err := c.db.Exec("DELETE FROM users"); err != nil {
		return fmt.Errorf("failed to reset the database: %v", err)
	}

	return nil
}
