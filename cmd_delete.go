package main

import (
	"fmt"
	"strconv"
)

func cmdDelete(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	if len(s.cache) == 0 {
		return errNoPasswords
	}

	s.printPasswords()

	pwIdxStr := s.getInput("number: ")
	pwIdxStr, err = validateInput(pwIdxStr)
	if err != nil {
		return err
	}

	pwIdx, err := strconv.Atoi(pwIdxStr)
	if err != nil {
		return errEnterNum
	}

	pw, ok := s.cache[pwIdx]
	if !ok {
		return errInvalidNum
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
