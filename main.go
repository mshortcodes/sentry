package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cmds := make(commands)
	cmds.add("hello", handlerHello)

	if len(os.Args) < 2 {
		log.Fatal("a command is required")
	}

	args := os.Args[1:]
	err := cmds.run(args[0])
	if err != nil {
		log.Fatalf("error with args: %v", err)
	}

	key, _ := makeKey([]byte("hello"))
	fmt.Println(key)
}
