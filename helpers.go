package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateUser(s *state) error {
	if s.user == nil {
		return errors.New("must be logged in")
	}
	return nil
}

func cleanInput(input string) string {
	lowered := strings.ToLower(input)
	args := strings.Fields(lowered)
	if len(args) < 1 {
		return ""
	}
	return args[0]
}

func validateInput(input string) string {
	args := strings.Fields(input)
	if len(args) < 1 {
		return ""
	}
	return args[0]
}

func printWelcomeMessage() {
	fmt.Print("\tWelcome to Sentry!🛡️\n")
	fmt.Print("\tType 'help' to view available comands\n\n")
}
