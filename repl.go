package main

import (
	"fmt"
	"strings"
)

func repl(s *state) {
	cmds := getCmds()
	printWelcomeMessage()

	for {
		fmt.Print("sentry > ")
		s.scanner.Scan()
		input := s.scanner.Text()

		input, err := validateInput(input)
		if err != nil {
			fmt.Printf("\t%s %v\n\n", failure, err)
			continue
		}

		input = strings.ToLower(input)

		cmd, ok := cmds[input]
		if !ok {
			fmt.Printf("\t%s invalid command\n\n", failure)
			continue
		}

		err = cmd.callback(s)
		if err != nil {
			fmt.Printf("\t%s %v\n\n", failure, err)
		}
	}
}

func printWelcomeMessage() {
	fmt.Printf("\tWelcome to Sentry!%s\n", shieldEmoji)
	fmt.Print("\tType 'help' to view available commands.\n\n")
}
