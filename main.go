package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mshortcodes/sentry/internal/database"
)

type state struct {
	db          database.Client
	currentUser string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("couldn't load .env file: %v", err)
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH must be set")
	}

	db, err := database.NewClient(dbPath)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}

	state := &state{
		db: db,
	}

	cmds := make(commands)
	cmds.add("hello", handlerHello)
	cmds.add("login", handlerLogin)
	cmds.add("register", handlerUsersCreate)

	if len(os.Args) < 2 {
		log.Fatal("a command is required")
	}

	args := os.Args[1:]

	err = cmds.run(args[0], state, args)
	if err != nil {
		log.Fatalf("error with args: %v", err)
	}
}
