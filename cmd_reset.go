package main

import (
	"fmt"
)

func cmdReset(s *state) error {
	fmt.Print("\tResetting the database...\n")
	if err := s.db.Reset(); err != nil {
		return fmt.Errorf("failed to reset db: %v", err)
	}

	fmt.Print("\tDatabase has been reset!\n\n")
	return nil
}
