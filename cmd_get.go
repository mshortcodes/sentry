package main

import (
	"fmt"
)

func cmdGet(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	dbPasswords, err := s.db.GetPasswords(s.user.Id)
	if err != nil {
		return fmt.Errorf("couldn't get passwords: %v", err)
	}

	if len(dbPasswords) == 0 {
		fmt.Print("\tno saved passwords\n\n")
		return nil
	}

	dbPasswordsCache := make(map[string]string)

	for _, dbPassword := range dbPasswords {
		fmt.Printf("\t%s\n", dbPassword.Name)
		dbPasswordsCache[dbPassword.Name] = dbPassword.Password
	}

	for {
		fmt.Print("\tpassword name: ")
		s.scanner.Scan()
		pwName := s.scanner.Text()

		pw, ok := dbPasswordsCache[pwName]
		if !ok {
			fmt.Println("\tinvalid password name")
			continue
		}

		fmt.Printf("\t%s\n\n", pw)
		break
	}

	return nil
}
