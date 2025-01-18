package database

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	Username string
	Password string
}

func (c Client) CreateUser(params CreateUserParams) error {
	id := uuid.New()

	query := `
	INSERT INTO users (id, username, created_at, updated_at, password)
	VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
	`

	if _, err := c.db.Exec(query, id, params.Username, params.Password); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
