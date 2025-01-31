package main

import (
	"fmt"
	"os"
)

func cmdExit(s *state) error {
	printGoodbyeMessage(s)
	os.Exit(0)
	return nil
}

func printGoodbyeMessage(s *state) {
	fmt.Println("\tExiting Sentry...")

	if s.user != nil {
		fmt.Printf("\tGoodbye, %s!%s\n\n", s.username, sentryLogo)
		return
	}

	fmt.Printf("\tGoodbye!%s\n\n", sentryLogo)
}
