package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func handlerLogin(s *state, args []string) error {
	if len(args) < 3 {
		return errors.New("must enter username and password")
	}

	username := args[1]
	password := args[2]

	user, err := s.db.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("couldn't get user: %v", err)
	}

	if err = auth.CheckPasswordHash(password, user.Password); err != nil {
		return fmt.Errorf("incorrect password: %v", err)
	}

	token, err := auth.GenerateToken()
	if err != nil {
		return fmt.Errorf("error generating token: %v", err)
	}

	dbToken, err := s.db.GetToken(user.Id)
	fmt.Printf("%+v", dbToken)

	if err = s.db.CreateToken(database.CreateTokenParams{
		Token:     token,
		UserID:    user.Id,
		ExpiresAt: time.Now().UTC().Add(5 * time.Minute),
	}); err != nil {
		return fmt.Errorf("couldn't create session: %v", err)
	}

	fmt.Printf("Welcome, %s\n", user.Username)
	fmt.Printf("Token: %s", token)
	fmt.Printf("Expires: %s\n", time.Now().UTC().Add(5*time.Minute).Format(time.RFC1123))

	return nil
}
