package main

import (
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdAdd(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	pwName, password, err := getPasswordInfo(s)
	if err != nil {
		return fmt.Errorf("error getting password info: %v", err)
	}

	nonce, err := crypt.GenerateNonce()
	if err != nil {
		return fmt.Errorf("failed to generate nonce: %v", err)
	}

	ciphertext, err := crypt.Encrypt([]byte(password), s.key, nonce)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %v", err)
	}

	err = s.db.AddPassword(database.AddPasswordParams{
		Name:     pwName,
		Password: fmt.Sprintf("%x", ciphertext),
		UserID:   s.user.Id,
		Nonce:    fmt.Sprintf("%x", nonce),
	})
	if err != nil {
		return fmt.Errorf("couldn't add password: %v", err)
	}

	s.cache = nil
	fmt.Print("\tpassword saved\n\n")
	return nil
}

func getPasswordInfo(s *state) (string, string, error) {
	fmt.Print("\tpassword name: ")
	s.scanner.Scan()
	pwName := s.scanner.Text()
	pwName, err := validateInput(pwName)
	if err != nil {
		return "", "", fmt.Errorf("error validating input: %v", err)
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password := s.scanner.Text()
	if len(password) < 8 {
		return "", "", errors.New("password must be at least 8 chars")
	}

	return pwName, password, nil
}
