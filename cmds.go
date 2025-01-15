package main

import (
	"errors"
	"fmt"
)

type commands map[string]func() error
type command func() error

func (c commands) add(name string, f command) {
	c[name] = f
}

func (c commands) run(name string) error {
	f, ok := c[name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	err := f()
	if err != nil {
		return fmt.Errorf("error running command: %v", err)
	}

	return nil
}
