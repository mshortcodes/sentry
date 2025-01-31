package main

import (
	"fmt"
	"strings"
)

func cmdWipe(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	for {
		fmt.Print("\tWipe all passwords? [y/n] ")
		s.scanner.Scan()
		input := s.scanner.Text()
		input, err = validateInput(input)
		if err != nil {
			return fmt.Errorf("error validating input: %v", err)
		}
		input = strings.ToLower(input)

		switch input {
		case "y":
			return wipePasswords(s)
		default:
			fmt.Println()
			return nil
		}
	}
}

func wipePasswords(s *state) error {
	fmt.Printf("\tWiping all passwords from %s...\n", s.username)
	if err := s.db.WipePasswords(s.user.Id); err != nil {
		return err
	}
	s.cache = nil
	fmt.Printf("\t%s Passwords have been wiped!\n\n", checkEmoji)
	return nil
}
