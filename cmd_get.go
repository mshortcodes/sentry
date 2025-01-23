package main

import (
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdGet() command {
	cmd := command{
		name:        "get",
		description: "retrieves passwords",
		callback:    handlerGet,
		flags:       flag.NewFlagSet("get", flag.ExitOnError),
	}

	cmd.flags.String("n", "", "(Optional) Specify password [n]ame")
	return cmd
}

func handlerGet(db database.Client, flags *flag.FlagSet, cmds commands) error {
	dbToken, err := auth.ValidateToken(db)
	if err != nil {
		return fmt.Errorf("must be logged in: %v", err)
	}

	dbPasswords, err := db.GetPasswords(dbToken.UserID)
	if err != nil {
		return fmt.Errorf("couldn't get passwords: %v", err)
	}

	if len(dbPasswords) == 0 {
		fmt.Println("no saved passwords")
		return nil
	}

	pwName := flags.Lookup("n").Value.String()

	for _, dbPassword := range dbPasswords {
		if pwName != "" && dbPassword.Name != pwName {
			continue
		}

		fmt.Printf("%s: %s\n", dbPassword.Name, dbPassword.Password)
	}

	return nil
}
