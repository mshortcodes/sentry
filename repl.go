package main

import (
	"fmt"
	"strings"
)

func repl(s *state) {
	cmds := getCmds()

	for {
		fmt.Print("sentry > ")
		s.scanner.Scan()
		input := s.scanner.Text()

		inputCmd := cleanInputCmd(input)
		if inputCmd == "" {
			fmt.Printf("\tenter a command\n\n")
		}

		cmd, ok := cmds[inputCmd]
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

func cleanInputCmd(input string) string {
	lowered := strings.ToLower(input)
	args := strings.Fields(lowered)
	if len(args) < 1 {
		return ""
	}
	return args[0]
}
