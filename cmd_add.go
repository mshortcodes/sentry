package main

import (
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdAdd(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	pwName, password, err := s.getPasswordInfo()
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

	if s.cache == nil {
		s.makeCache()
	}

	pwIdx := len(s.cache) + 1
	s.addToCache(password, pwName, pwIdx)

	fmt.Println()
	fmt.Printf("\t%s Password has been saved.\n\n", success)
	return nil
}

func (s *state) getPasswordInfo() (password, pwName string, err error) {
	fmt.Print("\tpassword name: ")
	s.scanner.Scan()
	pwName = s.scanner.Text()
	pwName, err = validateInput(pwName)
	if err != nil {
		return "", "", fmt.Errorf("error validating input: %v", err)
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password = s.scanner.Text()
	if len(password) < 8 {
		return "", "", errors.New("password must be at least 8 chars")
	}

	return pwName, password, nil
}
