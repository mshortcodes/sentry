package main

import (
	"errors"
	"fmt"
	"strconv"
)

func cmdDelete(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	if len(s.cache) == 0 {
		fmt.Print("\tNo saved passwords\n\n")
		return nil
	}

	s.printPasswords()

	pwIdxStr := s.getInput("number")
	pwIdxStr, err = validateInput(pwIdxStr)
	if err != nil {
		return err
	}

	pwIdx, err := strconv.Atoi(pwIdxStr)
	if err != nil {
		return errors.New("must enter a number")
	}

	pw, ok := s.cache[pwIdx]
	if !ok {
		return errors.New("invalid number")
	}

	err = s.db.DeletePassword(s.user.Id, pw.name)
	if err != nil {
		return err
	}

	s.deleteFromCache(pwIdx)

	fmt.Printf("\t%s Password for %s has been deleted.\n\n", success, pw.name)
	return nil
}

func (s *state) deleteFromCache(pwIdx int) {
	delete(s.cache, pwIdx)
}
