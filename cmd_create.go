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

	username, password, err := s.getCreateInfo()
	if err != nil {
		return fmt.Errorf("error getting user info: %v", err)
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}

	salt, err := crypt.GenerateSalt()
	if err != nil {
		return fmt.Errorf("couldn't generate salt: %v", err)
	}

	err = s.db.CreateUser(database.CreateUserParams{
		Username: strings.ToLower(username),
		Password: hash,
		Salt:     fmt.Sprintf("%x", salt),
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %v", err)
	}

	fmt.Println()
	fmt.Printf("\t%s %s has been created. Login to add passwords.\n\n", success, username)
	return nil
}

func (s *state) getCreateInfo() (username, password string, err error) {
	fmt.Print("\tusername: ")
	s.scanner.Scan()
	username = s.scanner.Text()
	username, err = validateInput(username)
	if err != nil {
		return "", "", fmt.Errorf("error validating input: %v", err)
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password = s.scanner.Text()
	if len(password) < 8 {
		return "", "", errors.New("password must be at least 8 chars")
	}

	return username, password, nil
}
