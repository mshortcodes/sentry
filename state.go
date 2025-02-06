package main

import (
	"bufio"
	"fmt"
	"slices"

	"github.com/mshortcodes/sentry/internal/database"
)

type state struct {
	db       *database.Client
	user     *database.User
	username string
	password string
	admin    bool
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

func (s *state) addToCache(plaintext, pwName string, idx int) {
	s.cache[idx] = passwordInfo{
		name:     pwName,
		password: plaintext,
	}
}

func (s *state) invalidateCache() {
	s.cache = nil
}

func (s *state) printPasswords() {
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
		return errLoggedIn
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
