package main

import (
	"fmt"

	"github.com/mshortcodes/sentry/internal/database"
)

func validateToken(s *state) (database.Token, error) {
	dbToken, err := s.db.GetToken()
	if err != nil {
		return database.Token{}, fmt.Errorf("invalid token")
	}

	return dbToken, nil
}
