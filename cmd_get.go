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
		pwIdx, err := s.getPasswordIdx()
		if err != nil {
			fmt.Printf("\t%s error getting password input: %v\n\n", failure, err)
			continue
		}

		pw, ok := s.cache[pwIdx]
		if !ok {
			fmt.Printf("\t%s invalid number\n\n", failure)
			continue
		}

		fmt.Printf("\t%s %s\n\n", success, pw.password)
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

func (s *state) getPasswordIdx() (pwIdx int, err error) {
	fmt.Print("\tnumber: ")
	s.scanner.Scan()
	input := s.scanner.Text()

	input, err = validateInput(input)
	if err != nil {
		return 0, err
	}

	pwIdx, err = strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("must enter a number: %v", err)
	}

	return pwIdx, nil
}
