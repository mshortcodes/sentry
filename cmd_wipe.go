package main

import (
	"fmt"
	"strings"
)

func cmdWipe(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	input := s.getInput("Wipe all passwords? [y/n] ")
	input, err = validateInput(input)
	if err != nil {
		return err
	}
	input = strings.ToLower(input)

	switch input {
	case "y":
		return s.wipePasswords()
	default:
		fmt.Println()
		return nil
	}
}

func (s *state) wipePasswords() error {
	fmt.Printf("\tWiping all passwords from %s...\n", s.username)
	if err := s.db.WipePasswords(s.user.Id); err != nil {
		return err
	}
	fmt.Println()
	s.invalidateCache()
	fmt.Printf("\t%s Passwords have been wiped!\n\n", success)
	return nil
}
