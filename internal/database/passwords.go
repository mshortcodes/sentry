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

func (c Client) GetPasswords(userID int) ([]Password, error) {
	query := `
	SELECT name, password
	FROM passwords
	WHERE user_id = ?
	ORDER BY name ASC
	`

	rows, err := c.db.Query(query, userID)
	if err != nil {
		return []Password{}, fmt.Errorf("failed to get passwords: %v", err)
	}
	defer rows.Close()

	var passwords []Password

	for rows.Next() {
		password := Password{}
		if err := rows.Scan(&password.Name, &password.Password); err != nil {
			return []Password{}, err
		}

		passwords = append(passwords, password)
	}

	return passwords, nil
}
