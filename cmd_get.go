package main

import (
	"fmt"
	"slices"
	"strconv"
)

func cmdGet(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	if len(s.cache) == 0 {
		fmt.Print("\tNo saved passwords\n\n")
		return nil
	}

	for {
		printPasswords(s)
		pwNumber, err := getPasswordInput(s)
		if err != nil {
			fmt.Printf("\terror getting password input: %v\n\n", err)
			continue
		}

		pw, ok := s.cache[pwNumber]
		if !ok {
			fmt.Print("\tinvalid number\n\n")
			continue
		}

		fmt.Printf("\t%s\n\n", pw.password)
		break
	}

	return nil
}

func printPasswords(s *state) {
	keys := make([]int, 0, len(s.cache))

	for key := range s.cache {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		fmt.Printf("\t[%d] %s\n", key, s.cache[key].name)
	}

	fmt.Println()
}

func getPasswordInput(s *state) (int, error) {
	fmt.Print("\tnumber: ")
	s.scanner.Scan()
	pwNumber, err := strconv.Atoi(s.scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("must enter a number: %v", err)
	}

	return pwNumber, nil
}
