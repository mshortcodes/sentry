package main

import (
	"bufio"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mshortcodes/sentry/internal/database"
)

type state struct {
	db       *database.Client
	user     *database.User
	username string
	password string
	key      []byte
	scanner  *bufio.Scanner
	cache    map[int]passwordInfo
}

type passwordInfo struct {
	name     string
	password string
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

	scanner := bufio.NewScanner(os.Stdin)

	state := state{
		db:      &db,
		scanner: scanner,
	}

	repl(&state)
}
