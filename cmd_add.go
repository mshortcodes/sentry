package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdAdd() command {
	cmd := command{
		name:        "add",
		description: "Adds a new password",
		callback:    handlerAdd,
		flags:       flag.NewFlagSet("add", flag.ExitOnError),
	}

	cmd.flags.String("n", "", "password [n]ame")
	cmd.flags.String("p", "", "[p]assword")
	return cmd
}

func handlerAdd(db database.Client, flags *flag.FlagSet, cmds commands) error {
	pwName := flags.Lookup("n").Value.String()
	if pwName == "" {
		return errors.New("password name can't be empty")
	}

	password := flags.Lookup("p").Value.String()
	if len(password) < 8 {
		return errors.New("password must be at least 8 chars")
	}

	dbToken, err := auth.ValidateToken(db)
	if err != nil {
		return fmt.Errorf("must be logged in: %v", err)
	}

	err = db.AddPassword(database.AddPasswordParams{
		Name:     pwName,
		Password: password,
		UserID:   dbToken.UserID,
	})

	if err != nil {
		return fmt.Errorf("couldn't add password: %v", err)
	}

	return nil
}
