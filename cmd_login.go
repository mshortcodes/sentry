package main

import (
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
)

func cmdLogin(s *state) error {
	fmt.Print("\tusername: ")
	s.scanner.Scan()
	username := s.scanner.Text()

	user, err := s.db.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("couldn't get user: %v", err)
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password := s.scanner.Text()

	err = auth.CheckPasswordHash(password, user.Password)
	if err != nil {
		return fmt.Errorf("incorrect password: %v", err)
	}

	s.user = &user
	s.password = password

	fmt.Printf("\tWelcome, %s\n\n", user.Username)
	return nil
}
