package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mshortcodes/sentry/internal/database"
)

func cmdHelp() command {
	cmd := command{
		name:        "help",
		description: "Adds a new password",
		callback:    handlerHelp,
		flags:       flag.NewFlagSet("help", flag.ExitOnError),
	}

	cmd.flags.Bool("v", false, "[v]erbose")
	return cmd
}

func handlerHelp(db database.Client, flags *flag.FlagSet, cmds commands) error {
	verbose, err := strconv.ParseBool(flags.Lookup("v").Value.String())
	if err != nil {
		fmt.Println(err)
	}

	separator := strings.Repeat("=", 40)

	keys := make([]string, 0, len(cmds))
	for key := range cmds {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		fmt.Println("Name:")
		fmt.Printf("  %s - %s", cmds[key].name, cmds[key].description)
		fmt.Println()
		fmt.Println()
		if verbose {
			cmds[key].flags.Usage()
			fmt.Println()
			fmt.Println(separator)
			fmt.Println()
		}
	}

	return nil
}
