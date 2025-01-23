package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdRegister() command {
	cmd := command{
		name:        "register",
		description: "Registers a new user",
		callback:    handlerRegister,
		flags:       flag.NewFlagSet("register", flag.ExitOnError),
	}

	cmd.flags.String("user", "", "username")
	cmd.flags.String("password", "", "password")
	return cmd
}

func handlerRegister(db database.Client, flags *flag.FlagSet) error {
	username := flags.Lookup("user").Value.String()
	if username == "" {
		return errors.New("must enter a username")
	}

	password := flags.Lookup("password").Value.String()
	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}

	return db.CreateUser(database.CreateUserParams{
		Username: username,
		Password: hash,
	})
}
