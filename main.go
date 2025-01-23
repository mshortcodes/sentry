package main

import (
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

	cmds := make(commands)
	cmds.add("create", cmdCreate())
	cmds.add("login", cmdLogin())
	cmds.add("add", cmdAdd())
	cmds.add("get", cmdGet())
	cmds.add("reset", cmdReset())

	if len(os.Args) < 2 {
		log.Fatal("a command is required")
	}

	cmd := os.Args[1]
	flags := os.Args[2:]

	err = cmds.run(cmd, flags, db)
	if err != nil {
		log.Fatalf("error with args: %v", err)
	}
}
