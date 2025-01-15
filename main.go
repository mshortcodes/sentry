package main

import (
	"log"
	"os"
)

func main() {
	cmds := make(commands)
	cmds.add("hello", handlerHello)
	args := os.Args
	err := cmds.run(args[1])
	if err != nil {
		log.Fatalf("error with args: %v", err)
	}
}
