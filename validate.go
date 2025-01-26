package main

import "errors"

func validateUser(s *state) error {
	if s.user == nil {
		return errors.New("must be logged in")
	}
	return nil
}
