package main

import (
	"bufio"
	"encoding/hex"
	"fmt"

	"github.com/mshortcodes/sentry/internal/crypt"
	"github.com/mshortcodes/sentry/internal/database"
)

type state struct {
	db       *database.Client
	user     *database.User
	username string
	password string
	key      []byte
	scanner  *bufio.Scanner
	cache    map[int]passwordInfo
}

type passwordInfo struct {
	name     string
	password string
}

func (s *state) invalidateCache() {
	s.cache = nil
}

func (s *state) makeCache() {
	s.cache = make(map[int]passwordInfo)
}

func (s *state) addToCache(dbPasswords []database.Password) error {
	s.makeCache()

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
