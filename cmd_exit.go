package main

import (
	"fmt"
	"os"
)

func cmdExit(s *state) error {
	s.printGoodbyeMessage()
	os.Exit(0)
	return nil
}

func (s *state) printGoodbyeMessage() {
	fmt.Println("\tExiting Sentry...")

	if s.user != nil {
		fmt.Printf("\tGoodbye, %s!%s\n\n", s.username, shieldEmoji)
		return
	}

	fmt.Printf("\tGoodbye!%s\n\n", shieldEmoji)
}
