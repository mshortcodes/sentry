package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mshortcodes/sentry/internal/database"
)

type handler func(db database.Client, flags *flag.FlagSet) error

type commands map[string]command

type command struct {
	name        string
	description string
	callback    handler
	flags       *flag.FlagSet
}

func (c commands) add(name string, cmd command) error {
	_, ok := c[name]
	if ok {
		return errors.New("command already exists")
	}

	c[name] = cmd
	return nil
}

func (c commands) run(name string, flags []string, db database.Client) error {
	cmd, ok := c[name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	if err := cmd.flags.Parse(flags); err != nil {
		return fmt.Errorf("couldn't parse flags: %v", err)
	}

	if err := cmd.callback(db, cmd.flags); err != nil {
		return fmt.Errorf("error running command: %v", err)
	}

	return nil
}
