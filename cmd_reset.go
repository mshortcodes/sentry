package main

import "fmt"

func cmdReset(s *state) error {
	fmt.Print("\tResetting database...\n")

	if err := s.db.Reset(); err != nil {
		return fmt.Errorf("failed to reset database: %v", err)
	}

	s.clearMemory()
	fmt.Printf("\t%s Database has been reset!\n\n", success)
	return nil
}
