package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string
}

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

	if _, err := c.db.Exec(
		query,
		id,
		params.Username,
		params.Password,
	); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (c Client) GetUserByUsername(username string) (User, error) {
	query := `
	SELECT * FROM users
	WHERE username = ?
	`

	var user User

	if err := c.db.QueryRow(query, username).Scan(
		&user.Id,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, fmt.Errorf("no users with that username: %v", err)
		}
		return User{}, err
	}

	return user, nil
}
