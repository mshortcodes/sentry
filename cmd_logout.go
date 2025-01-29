package main

import "fmt"

func cmdLogout(s *state) error {
	if s.user == nil {
		fmt.Print("\tno user is logged in\n\n")
		return nil
	}

	printLogoutMessage(s)

	s.user = nil
	s.password = ""
	s.key = nil
	s.cache = nil

	return nil
}

func printLogoutMessage(s *state) {
	fmt.Printf("\t%s has been logged out\n\n", s.user.Username)
}
