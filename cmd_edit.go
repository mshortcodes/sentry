package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mshortcodes/sentry/internal/crypt"
)

func cmdEdit(s *state) error {
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

	oldPw, ok := s.cache[pwIdx]
	if !ok {
		return errInvalidNum
	}

	input := s.getInput("update name? [y/n] ")
	input, err = validateInput(input)
	if err != nil {
		return err
	}
	input = strings.ToLower(input)

	var newPwName string
	switch input {
	case "y":
		newPwName = s.getInput("new name: ")
		newPwName, err = validateInput(newPwName)
		if err != nil {
			return err
		}
	case "n":
		newPwName = oldPw.name
	default:
		return errors.New("invalid input")
	}

	input = s.getInput("update password? [y/n] ")
	input, err = validateInput(input)
	if err != nil {
		return err
	}
	input = strings.ToLower(input)

	var newPw string
	switch input {
	case "y":
		newPw = s.getInput("password: ")
		err = validatePassword(newPw)
		if err != nil {
			return err
		}
	case "n":
		newPw = oldPw.password
	default:
		return errors.New("invalid input")
	}

	fmt.Println()

	if !s.checkIfUpdated(oldPw, passwordInfo{newPwName, newPw}) {
		return nil
	}

	nonce, err := crypt.GenerateNonce()
	if err != nil {
		return err
	}

	ciphertext, err := crypt.Encrypt([]byte(newPw), s.key, nonce)
	if err != nil {
		return err
	}

	err = s.db.UpdatePassword(
		s.user.Id,
		oldPw.name,
		newPwName,
		fmt.Sprintf("%x", ciphertext),
		fmt.Sprintf("%x", nonce))
	if err != nil {
		return err
	}

	s.addToCache(newPw, newPwName, pwIdx)
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
