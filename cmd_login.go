package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdLogin(s *state) error {
	isLoggedIn := s.validateUser() == nil
	if isLoggedIn {
		return errLoggedOut
	}

	user, username, password, err := s.getUserInfo()
	if err != nil {
		return err
	}

	err = auth.CheckPasswordHash(password, user.Password)
	if err != nil {
		return err
	}

	salt, err := hex.DecodeString(user.Salt)
	if err != nil {
		return err
	}

	key, err := crypt.GenerateKey([]byte(password), salt)
	if err != nil {
		return err
	}

	s.setUser(user)
	s.setPassword(password)
	s.setKey(key)
	s.setUsername(username)

	dbPasswords, err := s.fetchPasswords()
	if err != nil {
		return err
	}

	s.makeCache()

	for i, dbPassword := range dbPasswords {
		dbPasswordNonce, err := hex.DecodeString(dbPassword.Nonce)
		if err != nil {
			return errors.New("error decoding nonce")
		}

		ciphertext, err := hex.DecodeString(dbPassword.Password)
		if err != nil {
			return errors.New("error decoding ciphertext")
		}

		plaintext, err := crypt.Decrypt(ciphertext, s.key, dbPasswordNonce)
		if err != nil {
			return errors.New("couldn't decrypt password")
		}

		s.addToCache(plaintext, dbPassword.Name, i+1)
	}

	s.printLoginMessage()
	return nil
}

func (s *state) getUserInfo() (user *database.User, username, password string, err error) {
	fmt.Print("\tusername: ")
	s.scanner.Scan()
	username = s.scanner.Text()
	username, err = validateInput(username)
	if err != nil {
		return nil, "", "", err
	}

	dbUser, err := s.db.GetUserByUsername(strings.ToLower(username))
	if err != nil {
		return nil, "", "", err
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password = s.scanner.Text()
	fmt.Println()

	return &dbUser, username, password, nil
}

func (s *state) fetchPasswords() ([]database.Password, error) {
	dbPasswords, err := s.db.GetPasswords(s.user.Id)
	if err != nil {
		return nil, err
	}

	if dbPasswords == nil {
		return []database.Password{}, nil
	}

	return dbPasswords, nil
}

func (s *state) printLoginMessage() {
	fmt.Printf("\t%s Hello, %s.\n", success, s.username)

	switch len(s.cache) {
	case 1:
		fmt.Printf("\tYou have %d password saved.\n\n", len(s.cache))
	default:
		fmt.Printf("\tYou have %d passwords saved.\n\n", len(s.cache))
	}
}
