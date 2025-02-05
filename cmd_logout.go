package main

import (
	"errors"
	"fmt"
)

func cmdLogout(s *state) error {
	if s.user == nil {
		return errors.New("no user is logged in")
	}

	s.printLogoutMessage()
	s.clearMemory()
	return nil
}

func (s *state) printLogoutMessage() {
	fmt.Printf("\t%s %s has been logged out\n\n", success, s.username)
}
