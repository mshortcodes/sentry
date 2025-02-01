package main

import (
	"bufio"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mshortcodes/sentry/internal/database"
)

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
