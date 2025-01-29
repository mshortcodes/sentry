package main

import (
	"encoding/hex"
	"fmt"
	"slices"
	"strconv"

	"github.com/mshortcodes/sentry/internal/crypt"
)

func cmdGet(s *state) error {
	err := validateUser(s)
	if err != nil {
		return err
	}

	if s.cache != nil {
		return getPassword(s)
	}

	dbPasswords, err := s.db.GetPasswords(s.user.Id)
	if err != nil {
		return fmt.Errorf("couldn't get passwords: %v", err)
	}

	if len(dbPasswords) == 0 {
		fmt.Print("\tno saved passwords\n\n")
		return nil
	}

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

	return getPasswordFromCache(s)
}

func getPasswordFromCache(s *state) error {
	for {
		printPasswords(s)
		pwNumber, err := getPasswordInput(s)
		if err != nil {
			fmt.Printf("\terror getting password input: %v\n\n", err)
			continue
		}

		pw, ok := s.cache[pwNumber]
		if !ok {
			fmt.Print("\tinvalid number\n\n")
			continue
		}

		fmt.Printf("\t%s\n\n", pw.password)
		break
	}

	return nil
}

func printPasswords(s *state) {
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

func getPasswordInput(s *state) (int, error) {
	fmt.Print("\tnumber: ")
	s.scanner.Scan()
	pwNumber, err := strconv.Atoi(s.scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("must enter a number: %v", err)
	}

	return pwNumber, nil
}
