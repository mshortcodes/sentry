package main

import "fmt"

func handlerLogin(s *state, args []string) error {
	s.currentUser = args[1]
	fmt.Printf("Welcome, %s\n", s.currentUser)

	return nil
}
