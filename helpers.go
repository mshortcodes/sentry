package main

import (
	"errors"
	"strings"
)

var (
	shieldEmoji = "🛡️"
	checkEmoji  = "✔️"
	errEmoji    = "❌"
	keyEmoji    = "🔑"
	errNoName   = errors.New("must enter a name")
	errNoSpaces = errors.New("no spaces allowed")
)

func (s *state) validateUser() error {
	if s.user == nil {
		return errors.New("must be logged in")
	}
	return nil
}

func validateInput(input string) (string, error) {
	args := strings.Fields(input)

	switch {
	case len(args) < 1:
		return "", errNoName
	case len(args) > 1:
		return "", errNoSpaces
	}

	return args[0], nil
}

func (s *state) clearMemory() {
	s.user = nil
	s.username = ""
	s.password = ""
	s.key = nil
	s.cache = nil
}
