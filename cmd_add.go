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

	fmt.Print("\tpassword name: ")
	s.scanner.Scan()
	pwName := s.scanner.Text()
	pwName = validateInput(pwName)
	if pwName == "" {
		return errors.New("must enter a password name")
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password := s.scanner.Text()
	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
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

	fmt.Print("\tpassword saved\n\n")
	return nil
}
