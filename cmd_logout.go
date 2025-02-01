package main

import "fmt"

func cmdLogout(s *state) error {
	if s.user == nil {
		fmt.Print("\tno user is logged in\n\n")
		return nil
	}

	s.printLogoutMessage()
	s.clearMemory()

	return nil
}

func (s *state) printLogoutMessage() {
	fmt.Printf("\t%s %s has been logged out\n\n", checkEmoji, s.username)
}
