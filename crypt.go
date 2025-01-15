package main

import (
	"crypto/rand"
	"fmt"
)

func makeKey(plaintext []byte) ([]byte, error) {
	key := make([]byte, len(plaintext))

	_, err := rand.Read(key)
	if err != nil {
		return key, fmt.Errorf("error making key")
	}

	return key, nil
}
