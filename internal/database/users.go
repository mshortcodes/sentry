package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Id        int
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string
	Salt      string
}

type CreateUserParams struct {
	Username string
	Password string
	Salt     string
}

func (c Client) CreateUser(params CreateUserParams) error {
	query := `
	INSERT INTO users (username, created_at, updated_at, password, salt)
	VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?)
	`

	if _, err := c.db.Exec(
		query,
		params.Username,
		params.Password,
		params.Salt,
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
		&user.Password,
		&user.Salt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.New("no users with that username")
		}
		return User{}, err
	}

	return user, nil
}
