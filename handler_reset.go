package main

import "fmt"

func handlerReset(s *state, args []string) error {
	fmt.Println("resetting db...")
	if err := s.db.Reset(); err != nil {
		return fmt.Errorf("failed to reset db: %v", err)
	}

	return nil
}
