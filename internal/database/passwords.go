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
	Nonce     string
}

type AddPasswordParams struct {
	Name     string
	Password string
	UserID   int
	Nonce    string
}

func (c Client) AddPassword(params AddPasswordParams) error {
	query := `
	INSERT INTO passwords (name, password, created_at, updated_at, user_id, nonce)
	VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?)
	`

	if _, err := c.db.Exec(
		query,
		params.Name,
		params.Password,
		params.UserID,
		params.Nonce,
	); err != nil {
		return fmt.Errorf("failed to add password")
	}

	return nil
}

func (c Client) GetPasswords(userID int) ([]Password, error) {
	query := `
	SELECT name, password, nonce
	FROM passwords
	WHERE user_id = ?
	ORDER BY name ASC
	`

	rows, err := c.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get passwords: %v", err)
	}
	defer rows.Close()

	var passwords []Password

	for rows.Next() {
		password := Password{}
		if err := rows.Scan(&password.Name, &password.Password, &password.Nonce); err != nil {
			return nil, err
		}

		passwords = append(passwords, password)
	}

	return passwords, nil
}

func (c *Client) DeletePassword(userID int, pwName string) error {
	query := `
	DELETE FROM passwords
	WHERE user_id = ?
	AND name = ?
	`

	if _, err := c.db.Exec(query, userID, pwName); err != nil {
		return fmt.Errorf("failed to delete password: %v", err)
	}

	return nil
}

func (c *Client) WipePasswords(userID int) error {
	query := `
	DELETE FROM passwords
	WHERE user_id = ?
	`
	if _, err := c.db.Exec(query, userID); err != nil {
		return fmt.Errorf("failed to wipe passwords: %v", err)
	}

	return nil
}

func (c *Client) UpdatePassword(userID int, oldName, newName, password, nonce string) error {
	query := `
	UPDATE passwords
	SET name = ?, password = ?, nonce = ?, updated_at = CURRENT_TIMESTAMP
	WHERE user_id = ?
	AND name = ?
	`

	if _, err := c.db.Exec(
		query,
		newName,
		password,
		nonce,
		userID,
		oldName); err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}
