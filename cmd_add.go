package main

import (
	"fmt"

	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdAdd(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	pwName := s.getInput("password name: ")
	pwName, err = validateInput(pwName)
	if err != nil {
		return err
	}

	password := s.getInput("password: ")
	fmt.Println()
	err = validatePassword(password)
	if err != nil {
		return err
	}

	nonce, err := crypt.GenerateNonce()
	if err != nil {
		return err
	}

	ciphertext, err := crypt.Encrypt([]byte(password), s.key, nonce)
	if err != nil {
		return err
	}

	err = s.db.AddPassword(database.AddPasswordParams{
		Name:     pwName,
		Password: fmt.Sprintf("%x", ciphertext),
		UserID:   s.user.Id,
		Nonce:    fmt.Sprintf("%x", nonce),
	})
	if err != nil {
		return err
	}

	if s.cache == nil {
		s.makeCache()
	}

	pwIdx := len(s.cache) + 1
	s.addToCache(password, pwName, pwIdx)

	fmt.Printf("\t%s Password has been saved.\n\n", success)
	return nil
}
