package main

import "fmt"

func cmdReset(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	for {
		fmt.Print("\tAre you sure? [y/n] ")
		s.scanner.Scan()
		input := s.scanner.Text()
		input = cleanInput(input)

		switch input {
		case "y":
			fmt.Print("\tResetting passwords...\n")
			if err := s.db.Reset(s.user.Id); err != nil {
				return fmt.Errorf("failed to reset passwords: %v", err)
			}
			fmt.Print("\tPasswords have been reset!\n\n")
			return nil
		case "n":
			fmt.Println()
			return nil
		}
	}
}
