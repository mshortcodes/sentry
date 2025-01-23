package main

import (
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/database"
)

func cmdHello() command {
	cmd := command{
		name:        "hello",
		description: "Displays a welcome message",
		callback:    handlerHello,
		flags:       flag.NewFlagSet("hello", flag.ExitOnError),
	}

	cmd.flags.String("w", "hellooo", "add hellooo")
	return cmd
}

func handlerHello(db database.Client, flags *flag.FlagSet) error {
	fmt.Println("hello")
	return nil
}
