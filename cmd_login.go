package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/mshortcodes/sentry/internal/auth"
	"github.com/mshortcodes/sentry/internal/database"
)

func cmdLogin() command {
	cmd := command{
		name:        "login",
		description: "Logs a user in",
		callback:    handlerLogin,
		flags:       flag.NewFlagSet("login", flag.ExitOnError),
	}

	cmd.flags.String("u", "", "[u]sername")
	cmd.flags.String("p", "", "[p]assword")
	return cmd
}

func handlerLogin(db database.Client, flags *flag.FlagSet, cmds commands) error {
	username := flags.Lookup("u").Value.String()
	if username == "" {
		return errors.New("username can't be empty")
	}

	password := flags.Lookup("p").Value.String()
	if password == "" {
		return errors.New("password can't be empty")
	}

	user, err := db.GetUserByUsername(username)
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

	if err = db.CreateToken(database.CreateTokenParams{
		Token:     token,
		UserID:    user.Id,
		ExpiresAt: time.Now().UTC().Add(5 * time.Minute),
	}); err != nil {
		return fmt.Errorf("couldn't create session: %v", err)
	}

	fmt.Printf("Welcome, %s\n", user.Username)
	fmt.Printf("Token: %s\n", token)
	fmt.Printf("Expires: %s\n", time.Now().UTC().Add(5*time.Minute).Format(time.RFC1123))

	return nil
}
