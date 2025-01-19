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

type CreatePasswordParams struct {
	Name     uuid.UUID
	Password string
	UserID   string
}

func (c Client) CreatePassword(params CreatePasswordParams) error {
	id := uuid.New()

	query := `
	INSERT INTO passwords (id, name, password, created_at, updated_at, user_id)
	VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
	`

	if _, err := c.db.Exec(
		query,
		id,
		params.Name,
		params.Password,
		params.UserID,
	); err != nil {
		return fmt.Errorf("failed to create password")
	}

	return nil
}
