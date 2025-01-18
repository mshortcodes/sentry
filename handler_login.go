package main

import (
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
)

func handlerLogin(s *state, args []string) error {
	if len(args) < 3 {
		return errors.New("must enter username and password")
	}

	username := args[1]
	password := args[2]

	user, err := s.db.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("couldn't get user: %v", err)
	}

	if err = auth.CheckPasswordHash(password, user.Password); err != nil {
		return fmt.Errorf("incorrect password: %v", err)
	}

	apiKey, err := auth.GenerateAPIKey()
	if err != nil {
		return fmt.Errorf("error generating API key: %v", err)
	}

	fmt.Printf("Welcome, %s\n", user.Username)
	fmt.Printf("API key: %s", apiKey)
	return nil
}
