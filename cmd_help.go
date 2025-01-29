package main

import (
	"fmt"
	"slices"
)

func cmdHelp(s *state) error {
	cmds := getCmds()

	keys := make([]string, 0, len(cmds))
	for key := range cmds {
		keys = append(keys, key)
	}

	slices.Sort(keys)
	printHelpMsg(cmds, keys)
	return nil
}

func printHelpMsg(cmds map[string]cmd, keys []string) {
	for _, key := range keys {
		fmt.Printf("\t%s - %s\n", cmds[key].name, cmds[key].description)
	}
	fmt.Println()
}
