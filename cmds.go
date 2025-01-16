package main

import (
	"errors"
	"fmt"
)

type commands map[string]func(*state, []string) error

func (c commands) add(name string, f func(s *state, args []string) error) {
	c[name] = f
}

func (c commands) run(name string, s *state, args []string) error {
	f, ok := c[name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	err := f(s, args)
	if err != nil {
		return fmt.Errorf("error running command: %v", err)
	}

	return nil
}
