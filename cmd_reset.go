package main

import (
	"fmt"
)

func cmdReset(s *state) error {
	fmt.Println("Resetting the database...")
	if err := s.db.Reset(); err != nil {
		return fmt.Errorf("failed to reset db: %v", err)
	}

	fmt.Println("Database has been reset!")
	return nil
}
