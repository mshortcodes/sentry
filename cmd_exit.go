package main

import (
	"fmt"
	"os"
)

func cmdExit(s *state) error {
	fmt.Println("\tExiting Sentry...")

	if s.user != nil {
		fmt.Printf("\tGoodbye, %s!\n\n", s.user.Username)
		os.Exit(0)
		return nil
	}

	fmt.Print("\tGoodbye!\n\n")
	os.Exit(0)
	return nil
}
