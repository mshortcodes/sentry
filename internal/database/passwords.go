package database

import (
	"fmt"
	"time"
)

type Password struct {
	Id        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
}

type AddPasswordParams struct {
	Name     string
	Password string
	UserID   int
}

func (c Client) AddPassword(params AddPasswordParams) error {
	query := `
	INSERT INTO passwords (name, password, created_at, updated_at, user_id)
	VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
	`

	if _, err := c.db.Exec(
		query,
		params.Name,
		params.Password,
		params.UserID,
	); err != nil {
		return fmt.Errorf("failed to add password")
	}

	return nil
}

func (c Client) GetPasswords(userID int) error {
	query := `
	SELECT name, password
	FROM passwords
	WHERE user_id = ?
	ORDER BY name ASC
	`

	rows, err := c.db.Query(query, userID)
	if err != nil {
		return fmt.Errorf("failed to get passwords: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var password string

		if err := rows.Scan(&name, &password); err != nil {
			return err
		}

		fmt.Printf("%s: %s\n", name, password)
	}

	return nil
}
