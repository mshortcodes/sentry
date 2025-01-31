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
			fmt.Printf("\terror validating input: %v\n\n", err)
			continue
		}

		input = strings.ToLower(input)

		cmd, ok := cmds[input]
		if !ok {
			fmt.Print("\tinvalid command\n\n")
			continue
		}

		err = cmd.callback(s)
		if err != nil {
			fmt.Printf("\t%v\n\n", err)
		}
	}
}

func printWelcomeMessage() {
	fmt.Print("\tWelcome to Sentry!🛡️\n")
	fmt.Print("\tType 'help' to view available commands.\n\n")
}
