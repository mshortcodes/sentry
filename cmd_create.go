package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdCreate() command {
	cmd := command{
		name:        "create",
		description: "Creates a new user",
		callback:    handlerCreate,
		flags:       flag.NewFlagSet("create", flag.ExitOnError),
	}

	cmd.flags.String("u", "", "[u]sername")
	cmd.flags.String("p", "", "[p]assword")
	return cmd
}

func handlerCreate(db database.Client, flags *flag.FlagSet) error {
	username := flags.Lookup("u").Value.String()
	if username == "" {
		return errors.New("must enter a username")
	}

	password := flags.Lookup("p").Value.String()
	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}

	err = db.CreateUser(database.CreateUserParams{
		Username: username,
		Password: hash,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %v", err)
	}

	fmt.Printf("%s has been created. Login to add passwords.\n", username)

	return nil
}
