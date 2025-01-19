package main

import (
	"errors"
	"fmt"

	"github.com/mshortcodes/sentry/internal/database"
)

func handlerPasswordsAdd(s *state, args []string) error {
	if len(args) < 3 {
		return errors.New("must provide a name and password")
	}

	dbToken, err := validateToken(s)
	if err != nil {
		return fmt.Errorf("couldn't validate token: %v", err)
	}

	name := args[1]
	password := args[2]

	if name == "" {
		return errors.New("password name can't be empty")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
	}

	err = s.db.AddPassword(database.AddPasswordParams{
		Name:     name,
		Password: password,
		UserID:   dbToken.UserID,
	})

	if err != nil {
		return fmt.Errorf("couldn't add password: %v", err)
	}

	return nil
}
