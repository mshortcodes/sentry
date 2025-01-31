package main

import "fmt"

func cmdLogout(s *state) error {
	if s.user == nil {
		fmt.Print("\tno user is logged in\n\n")
		return nil
	}

	printLogoutMessage(s)
	clearMemory(s)

	return nil
}

func printLogoutMessage(s *state) {
	fmt.Printf("\t%s %s has been logged out\n\n", checkEmoji, s.username)
}
