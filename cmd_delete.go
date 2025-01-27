package main

import "fmt"

func cmdDelete(s *state) error {
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
			fmt.Print("\tDeleting passwords...\n")
			if err := s.db.DeletePasswords(s.user.Id); err != nil {
				return fmt.Errorf("failed to delete passwords: %v", err)
			}
			fmt.Print("\tPasswords have been deleted!\n\n")
			return nil
		default:
			fmt.Println()
			return nil
		}
	}
}
