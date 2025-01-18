package main

import (
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func handlerUsersCreate(s *state, args []string) error {
	if len(args) < 3 {
		return errors.New("must provide a username and password")
	}

	username := args[1]
	password := args[2]

	if username == "" {
		return errors.New("must enter a username")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}

	return s.db.CreateUser(database.CreateUserParams{
		Username: username,
		Password: hash,
	})
}
