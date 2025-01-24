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
		description: "lists available commands",
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

	fmt.Println("Syntax:")
	fmt.Printf("  [CMD] [FLAG] [VALUE]\n")
	fmt.Println("  'help -h'")
	fmt.Println("  'add -n phone -p 12345678'")
	fmt.Println("  Run [CMD] -h to view its flags")
	fmt.Println()
	fmt.Println(separator)
	fmt.Println()

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
