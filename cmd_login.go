package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdLogin(s *state) error {
	isLoggedIn := validateUser(s) == nil
	if isLoggedIn {
		return errors.New("must be logged out")
	}

	user, password, err := getUserInfo(s)
	if err != nil {
		return fmt.Errorf("invalid user info: %v", err)
	}

	err = auth.CheckPasswordHash(password, user.Password)
	if err != nil {
		return fmt.Errorf("incorrect password: %v", err)
	}

	salt, err := hex.DecodeString(user.Salt)
	if err != nil {
		return fmt.Errorf("couldn't decode hex string: %v", err)
	}

	key, err := crypt.GenerateKey([]byte(password), salt)
	if err != nil {
		return fmt.Errorf("failed to generate key: %v", err)
	}

	s.user = user
	s.password = password
	s.key = key

	fmt.Printf("\tWelcome, %s\n\n", user.Username)
	return nil
}

func getUserInfo(s *state) (*database.User, string, error) {
	fmt.Print("\tusername: ")
	s.scanner.Scan()
	username := s.scanner.Text()

	user, err := s.db.GetUserByUsername(username)
	if err != nil {
		return &user, "", fmt.Errorf("couldn't get user: %v", err)
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password := s.scanner.Text()

	return &user, password, nil
}
