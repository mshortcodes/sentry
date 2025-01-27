package main

import "fmt"

func cmdLogout(s *state) error {
	if s.user == nil {
		fmt.Print("\tno user is logged in\n\n")
		return nil
	}

	fmt.Printf("\t%s has been logged out\n\n", s.user.Username)
	s.user = nil
	return nil
}
