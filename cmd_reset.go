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

	cmd.flags.Bool("f", false, "force reset the database")
	return cmd
}

func handlerReset(db database.Client, flags *flag.FlagSet) error {
	if flags.Lookup("f").Value.String() != "true" {
		fmt.Println("must provide the force flag")
		return nil
	}

	fmt.Println("resetting db...")
	if err := db.Reset(); err != nil {
		return fmt.Errorf("failed to reset db: %v", err)
	}
	return nil
}
