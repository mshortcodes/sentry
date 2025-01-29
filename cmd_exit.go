package main

import (
	"fmt"
	"os"
)

func cmdExit(s *state) error {
	printGoodbyeMsg(s)
	os.Exit(0)
	return nil
}

func printGoodbyeMsg(s *state) {
	fmt.Println("\tExiting Sentry...")

	if s.user != nil {
		fmt.Printf("\tGoodbye, %s!\n\n", s.user.Username)
		return
	}

	fmt.Print("\tGoodbye!\n\n")
}
