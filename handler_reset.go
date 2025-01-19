package main

import "fmt"

func handlerReset(s *state, args []string) error {
	fmt.Println("resetting db...")
	s.db.Reset()

	return nil
}
