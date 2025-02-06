package main

import (
	"bufio"
	"log"
	"os"

	"github.com/mshortcodes/sentry/internal/database"
)

func main() {
	err := createProjectDir()
	if err != nil {
		log.Fatalf("couldn't create project directory: %v", err)
	}

	dbPath, err := getDBPath()
	if err != nil {
		log.Fatalf("couldn't get database path: %v", err)
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
