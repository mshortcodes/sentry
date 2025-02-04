package main

import (
	"fmt"
	"strconv"
)

func cmdGet(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	if len(s.cache) == 0 {
		return errNoPasswords
	}

	for {
		s.printPasswords()

		pwIdxStr := s.getInput("number: ")
		pwIdxStr, err = validateInput(pwIdxStr)
		if err != nil {
			fmt.Printf("\t%s %v\n\n", failure, err)
			continue
		}

		pwIdx, err := strconv.Atoi(pwIdxStr)
		if err != nil {
			fmt.Printf("\t%s %v\n\n", failure, errEnterNum)
			continue
		}

		pw, ok := s.cache[pwIdx]
		if !ok {
			fmt.Printf("\t%s %v\n\n", failure, errInvalidNum)
			continue
		}

		fmt.Printf("\t%s %s\n\n", success, pw.password)
		break
	}

	return nil
}
