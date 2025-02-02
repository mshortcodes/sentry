package main

import (
	"bufio"
	"encoding/hex"
	"errors"
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

func (s *state) makeCache() {
	s.cache = make(map[int]passwordInfo)
}

func (s *state) addToCache(plaintext, pwName string, idx int) error {
	s.cache[idx] = passwordInfo{
		name:     pwName,
		password: plaintext,
	}
	return nil
}

func (s *state) invalidateCache() {
	s.cache = nil
}

func (s *state) setUser(user *database.User) {
	s.user = user
}

func (s *state) setPassword(password string) {
	s.password = password
}

func (s *state) setKey(key []byte) {
	s.key = key
}

func (s *state) setUsername(username string) {
	s.username = username
}

func (s *state) validateUser() error {
	if s.user == nil {
		return errors.New("must be logged in")
	}
	return nil
}

func (s *state) clearMemory() {
	s.user = nil
	s.username = ""
	s.password = ""
	s.key = nil
	s.cache = nil
}

func (s *state) decrypt(dbPassword database.Password) (plaintext string, err error) {
	dbPasswordNonce, err := hex.DecodeString(dbPassword.Nonce)
	if err != nil {
		return "", fmt.Errorf("error decoding nonce: %v", err)
	}

	ciphertext, err := hex.DecodeString(dbPassword.Password)
	if err != nil {
		return "", fmt.Errorf("error decoding ciphertext: %v", err)
	}

	plaintext, err = crypt.Decrypt(ciphertext, s.key, dbPasswordNonce)
	if err != nil {
		return "", fmt.Errorf("couldn't decrypt password: %v", err)
	}

	return plaintext, nil
}
