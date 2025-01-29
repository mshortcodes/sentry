package main

import "fmt"

func cmdReset(s *state) error {
	fmt.Print("\tResetting database...\n")
	if err := s.db.Reset(); err != nil {
		return fmt.Errorf("failed to reset database: %v", err)
	}

	s.user = nil
	s.password = ""
	s.key = nil
	s.cache = nil

	fmt.Print("\tDatabase has been reset!\n\n")
	return nil
}
