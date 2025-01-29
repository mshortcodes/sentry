package main

import "fmt"

func cmdReset(s *state) error {
	fmt.Print("\tResetting database...\n")

	if err := s.db.Reset(); err != nil {
		return fmt.Errorf("failed to reset database: %v", err)
	}

	clearMemory(s)
	fmt.Print("\tDatabase has been reset!\n\n")
	return nil
}
