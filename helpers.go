package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	shieldEmoji = "🛡️"
	success     = brightGreen + checkMark + reset
	failure     = red + crossMark + reset
	brightGreen = "\033[92m"
	red         = "\033[31m"
	reset       = "\033[0m"
	checkMark   = "\u2713"
	crossMark   = "\u2717"
)

var (
	errEmpty       = errors.New("input can't be empty")
	errNoSpaces    = errors.New("no spaces allowed")
	errLoggedIn    = errors.New("must be logged in")
	errLoggedOut   = errors.New("must be logged out")
	errPassword    = errors.New("password must be at least 8 chars")
	errNoPasswords = errors.New("no saved passwords")
	errEnterNum    = errors.New("must enter a number")
	errInvalidNum  = errors.New("invalid number")
)

func (s *state) getInput(prompt string) string {
	fmt.Printf("\t%s: ", prompt)
	s.scanner.Scan()
	input := s.scanner.Text()
	return input
}

func validateInput(input string) (string, error) {
	args := strings.Fields(input)

	switch {
	case len(args) < 1:
		return "", errEmpty
	case len(args) > 1:
		return "", errNoSpaces
	}

	return args[0], nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errPassword
	}
	return nil
}
