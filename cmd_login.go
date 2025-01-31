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
	isLoggedIn := validateUser(s) == nil
	if isLoggedIn {
		return errors.New("must be logged out")
	}

	user, username, password, err := getUserInfo(s)
	if err != nil {
		return fmt.Errorf("invalid user info: %v", err)
	}

	err = auth.CheckPasswordHash(password, user.Password)
	if err != nil {
		return fmt.Errorf("incorrect password: %v", err)
	}

	salt, err := hex.DecodeString(user.Salt)
	if err != nil {
		return fmt.Errorf("couldn't decode hex string: %v", err)
	}

	key, err := crypt.GenerateKey([]byte(password), salt)
	if err != nil {
		return fmt.Errorf("failed to generate key: %v", err)
	}

	s.user = user
	s.password = password
	s.key = key
	s.username = username

	dbPasswords, err := fetchPasswords(s)
	if err != nil {
		return fmt.Errorf("error fetching passwords: %v", err)
	}

	err = addToCache(s, dbPasswords)
	if err != nil {
		return fmt.Errorf("failed to add to cache: %v", err)
	}

	printLoginMessage(s)
	return nil
}

func getUserInfo(s *state) (user *database.User, username, password string, err error) {
	fmt.Print("\tusername: ")
	s.scanner.Scan()
	username = s.scanner.Text()
	username, err = validateInput(username)
	if err != nil {
		return nil, "", "", fmt.Errorf("error validating input: %v", err)
	}

	dbUser, err := s.db.GetUserByUsername(strings.ToLower(username))
	if err != nil {
		return nil, "", "", fmt.Errorf("couldn't get user: %v", err)
	}

	fmt.Print("\tpassword: ")
	s.scanner.Scan()
	password = s.scanner.Text()
	fmt.Println()

	return &dbUser, username, password, nil
}

func fetchPasswords(s *state) ([]database.Password, error) {
	dbPasswords, err := s.db.GetPasswords(s.user.Id)
	if err != nil {
		return nil, fmt.Errorf("couldn't get passwords: %v", err)
	}

	if dbPasswords == nil {
		return []database.Password{}, nil
	}

	return dbPasswords, nil
}

func addToCache(s *state, dbPasswords []database.Password) error {
	s.cache = make(map[int]passwordInfo)

	for i, dbPassword := range dbPasswords {
		dbPasswordNonce, err := hex.DecodeString(dbPassword.Nonce)
		if err != nil {
			return fmt.Errorf("error decoding nonce: %v", err)
		}

		ciphertext, err := hex.DecodeString(dbPassword.Password)
		if err != nil {
			return fmt.Errorf("error decoding ciphertext: %v", err)
		}

		plaintext, err := crypt.Decrypt(ciphertext, s.key, dbPasswordNonce)
		if err != nil {
			return fmt.Errorf("couldn't decrypt password: %v", err)
		}

		s.cache[i+1] = passwordInfo{
			name:     dbPassword.Name,
			password: plaintext,
		}
	}

	return nil
}

func printLoginMessage(s *state) {
	fmt.Printf("\t%s Hello, %s.\n", checkMark, s.username)

	switch len(s.cache) {
	case 1:
		fmt.Printf("\tYou have %d password saved.\n\n", len(s.cache))
	default:
		fmt.Printf("\tYou have %d passwords saved.\n\n", len(s.cache))
	}
}
