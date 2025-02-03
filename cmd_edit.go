package main

import (
	"errors"
	"fmt"
	"strings"
)

func cmdEdit(s *state) error {
	err := s.validateUser()
	if err != nil {
		return err
	}

	s.printPasswords()

	pwIdx, err := s.getPasswordIdx()
	if err != nil {
		return fmt.Errorf("error getting input: %v", err)
	}

	oldPw, ok := s.cache[pwIdx]
	if !ok {
		return errors.New("invalid number")
	}

	newPwName, err := s.promptEditName()
	if err != nil {
		return fmt.Errorf("error getting new name: %v", err)
	}
	if newPwName == "" {
		newPwName = oldPw.name
	}

	newPw, err := s.promptEditPassword()
	if err != nil {
		return fmt.Errorf("error getting new password: %v", err)
	}
	if newPw == "" {
		newPw = oldPw.password
	}

	if !s.checkIfUpdated(oldPw, passwordInfo{newPwName, newPw}) {
		return nil
	}

	err = s.addToCache(newPw, newPwName, pwIdx)
	if err != nil {
		return fmt.Errorf("couldn't update password: %v", err)
	}

	fmt.Printf("\t%s Password has been updated.\n\n", success)
	return nil
}

func (s *state) checkIfUpdated(oldPw, newPw passwordInfo) bool {
	switch newPw {
	case oldPw:
		return false
	default:
		return true
	}
}

func (s *state) promptEditName() (pwName string, err error) {
	fmt.Print("\tUpdate name? [y/n] ")
	s.scanner.Scan()
	input := s.scanner.Text()
	input, err = validateInput(input)
	if err != nil {
		return "", fmt.Errorf("error validating input: %v", err)
	}
	input = strings.ToLower(input)

	switch input {
	case "y":
		return s.getNewName()
	default:
		fmt.Println()
		return "", nil
	}
}

func (s *state) promptEditPassword() (password string, err error) {
	fmt.Print("\tUpdate password? [y/n] ")
	s.scanner.Scan()
	input := s.scanner.Text()
	input, err = validateInput(input)
	if err != nil {
		return "", fmt.Errorf("error validating input: %v", err)
	}
	input = strings.ToLower(input)

	switch input {
	case "y":
		return s.getNewPassword()
	default:
		fmt.Println()
		return "", nil
	}
}

func (s *state) getNewName() (pwName string, err error) {
	fmt.Print("\tpassword name: ")
	s.scanner.Scan()
	pwName = s.scanner.Text()
	pwName, err = validateInput(pwName)
	if err != nil {
		return "", fmt.Errorf("error validating input: %v", err)
	}
	fmt.Println()

	return pwName, nil
}

func (s *state) getNewPassword() (password string, err error) {
	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password = s.scanner.Text()
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 chars")
	}
	fmt.Println()

	return password, nil
}
