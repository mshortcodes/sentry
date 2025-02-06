package main

import (
	"errors"
	"fmt"
)

func cmdReset(s *state) error {
	if !s.admin {
		return errors.New("only allowed in admin mode")
	}

	fmt.Print("\tResetting database...\n")

	if err := s.db.Reset(); err != nil {
		return err
	}

	s.clearMemory()
	fmt.Printf("\t%s Database has been reset!\n\n", success)
	return nil
}
