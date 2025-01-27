package main

import (
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdCreate(s *state) error {
	fmt.Print("\tusername: ")
	s.scanner.Scan()
	username := s.scanner.Text()
	username = validateInput(username)
	if username == "" {
		return errors.New("must enter a username")
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password := s.scanner.Text()
	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}

	err = s.db.CreateUser(database.CreateUserParams{
		Username: username,
		Password: hash,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %v", err)
	}

	fmt.Printf("\t%s has been created. Login to add passwords.\n\n", username)

	return nil
}
