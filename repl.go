package main

import (
	"fmt"
)

func repl(s *state) {
	cmds := getCmds()
	printWelcomeMessage()

	for {
		fmt.Print("sentry > ")
		s.scanner.Scan()
		input := s.scanner.Text()

		cleaned := cleanInput(input)
		if cleaned == "" {
			fmt.Printf("\tenter a command\n\n")
			continue
		}

		cmd, ok := cmds[cleaned]
		if !ok {
			fmt.Print("\tinvalid command\n\n")
			continue
		}

		err := cmd.callback(s)
		if err != nil {
			fmt.Printf("\t%v\n\n", err)
		}
	}
}

func printWelcomeMessage() {
	fmt.Print("\tWelcome to Sentry!🛡️\n")
	fmt.Print("\tType 'help' to view available comands\n\n")
}
