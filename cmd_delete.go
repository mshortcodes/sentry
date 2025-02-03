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

	pwIdx, err := s.getPasswordIdx()
	if err != nil {
		return err
	}

	pw, ok := s.cache[pwIdx]
	if !ok {
		return errors.New("invalid number")
	}

	err = s.db.DeletePassword(s.user.Id, pw.name)
	if err != nil {
		return fmt.Errorf("error deleting password: %v", err)
	}

	s.deleteFromCache(pwIdx)

	fmt.Printf("\t%s Password for %s has been deleted.\n\n", success, pw.name)
	return nil
}

func (s *state) deleteFromCache(pwIdx int) {
	delete(s.cache, pwIdx)
}
