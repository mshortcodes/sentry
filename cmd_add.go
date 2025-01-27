package main

import (
	"errors"
	"fmt"

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

	err = s.db.AddPassword(database.AddPasswordParams{
		Name:     pwName,
		Password: password,
		UserID:   s.user.Id,
	})
	if err != nil {
		return fmt.Errorf("couldn't add password: %v", err)
	}

	fmt.Print("\tpassword saved\n\n")
	return nil
}
