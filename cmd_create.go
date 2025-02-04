package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdCreate(s *state) error {
	isLoggedIn := s.validateUser() == nil
	if isLoggedIn {
		return errors.New("must be logged out")
	}

	username := s.getInput("username")
	username, err := validateInput(username)
	if err != nil {
		return err
	}

	password := s.getInput("password")
	err = validatePassword(password)
	if err != nil {
		return err
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	salt, err := crypt.GenerateSalt()
	if err != nil {
		return err
	}

	err = s.db.CreateUser(database.CreateUserParams{
		Username: strings.ToLower(username),
		Password: hash,
		Salt:     fmt.Sprintf("%x", salt),
	})
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("\t%s %s has been created. Login to add passwords.\n\n", success, username)
	return nil
}
