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
	errEmpty    = errors.New("input can't be empty")
	errNoSpaces = errors.New("no spaces allowed")
)

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
