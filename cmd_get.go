package main

import (
	"fmt"
	"slices"
	"strconv"
)

func cmdGet(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	if len(s.cache) == 0 {
		fmt.Print("\tNo saved passwords\n\n")
		return nil
	}

	for {
		s.printPasswords()
		pwNumber, err := s.getPasswordInput()
		if err != nil {
			fmt.Printf("\t%s error getting password input: %v\n\n", errEmoji, err)
			continue
		}

		pw, ok := s.cache[pwNumber]
		if !ok {
			fmt.Printf("\t%s invalid number\n\n", errEmoji)
			continue
		}

		fmt.Printf("\t%s %s\n\n", keyEmoji, pw.password)
		break
	}

	return nil
}

func (s *state) printPasswords() {
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

func (s *state) getPasswordInput() (int, error) {
	fmt.Print("\tnumber: ")
	s.scanner.Scan()
	pwNumber, err := strconv.Atoi(s.scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("must enter a number: %v", err)
	}

	return pwNumber, nil
}
