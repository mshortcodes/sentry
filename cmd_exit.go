package main

import (
	"fmt"
	"os"
)

func cmdExit(s *state) error {
	fmt.Println("\tExiting Sentry...")
	fmt.Print("\tGoodbye!\n\n")
	os.Exit(0)
	return nil
}
