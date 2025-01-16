package main

import "fmt"

func handlerHello(s *state, args []string) error {
	fmt.Println("hello")
	return nil
}
