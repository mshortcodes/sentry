package main

import (
	"errors"
	"fmt"
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

	pwIdx, err := s.getPasswordInput()
	if err != nil {
		return err
	}

	pw, ok := s.cache[pwIdx]
	if !ok {
		return errors.New("invalid number")
	}

	s.deleteFromCache(pwIdx)
	fmt.Printf("\t%s Password for %s has been deleted.\n\n", checkEmoji, pw.name)

	return nil
}

func (s *state) deleteFromCache(pwIdx int) {
	delete(s.cache, pwIdx)
}
