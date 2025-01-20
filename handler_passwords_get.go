package main

import "fmt"

func handlerPasswordsGet(s *state, args []string) error {
	dbToken, err := validateToken(s)
	if err != nil {
		return fmt.Errorf("must be logged in: %v", err)
	}

	if err := s.db.GetPasswords(dbToken.UserID); err != nil {
		return fmt.Errorf("couldn't get passwords: %v", err)
	}

	return nil
}
