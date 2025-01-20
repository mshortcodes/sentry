package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Password struct {
	Id        uuid.UUID
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
}

type AddPasswordParams struct {
	Name     string
	Password string
	UserID   uuid.UUID
}

func (c Client) AddPassword(params AddPasswordParams) error {
	id := uuid.New()

	query := `
	INSERT INTO passwords (id, name, password, created_at, updated_at, user_id)
	VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
	`

	if _, err := c.db.Exec(
		query,
		id.String(),
		params.Name,
		params.Password,
		params.UserID,
	); err != nil {
		return fmt.Errorf("failed to add password")
	}

	return nil
}

func (c Client) GetPasswords(userID uuid.UUID) error {
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
