package main

import "fmt"

func cmdWipe(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	for {
		fmt.Print("\tWipe all passwords? [y/n] ")
		s.scanner.Scan()
		input := s.scanner.Text()
		input = cleanInput(input)

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
	fmt.Printf("\tWiping all passwords from %s...\n", s.user.Username)
	if err := s.db.WipePasswords(s.user.Id); err != nil {
		return err
	}
	s.cache = nil
	fmt.Print("\tPasswords have been wiped!\n\n")
	return nil
}
