package main

import (
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/database"
)

func cmdReset() command {
	cmd := command{
		name:        "reset",
		description: "Resets the database",
		callback:    handlerReset,
		flags:       flag.NewFlagSet("reset", flag.ExitOnError),
	}

	cmd.flags.Bool("c", false, "[c]onfirm resetting the database")
	return cmd
}

func handlerReset(db database.Client, flags *flag.FlagSet) error {
	if flags.Lookup("c").Value.String() != "true" {
		fmt.Println("must provide the confirm flag")
		return nil
	}

	fmt.Println("Resetting the database...")
	if err := db.Reset(); err != nil {
		return fmt.Errorf("failed to reset db: %v", err)
	}

	fmt.Println("Database has been reset!")
	return nil
}
