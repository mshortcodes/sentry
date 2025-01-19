package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	ExpiresAt time.Time
}

type CreateTokenParams struct {
	Token     string
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func (c Client) CreateToken(params CreateTokenParams) error {
	query := `
	INSERT INTO tokens (token, created_at, updated_at, user_id, expires_at)
	VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?)
	`

	if _, err := c.db.Exec(query, params.Token, params.UserID, params.ExpiresAt); err != nil {
		return fmt.Errorf("failed to create token: %v", err)
	}

	return nil
}

func (c Client) GetToken(userID uuid.UUID) (Token, error) {
	var token Token

	query := `
	SELECT * FROM tokens
	WHERE expires_at > CURRENT_TIMESTAMP
	AND user_id = ?
	ORDER BY expires_at DESC
	`

	if err := c.db.QueryRow(query, userID).Scan(
		&token.Token,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.UserID,
		&token.ExpiresAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Token{}, fmt.Errorf("invalid token")
		}
		return Token{}, fmt.Errorf("error getting token: %v", err)
	}

	return token, nil
}
