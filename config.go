package main

import (
	"os"
	"path/filepath"
)

const projectName = "sentry"
const dbFileName = "sentry.db"

func createProjectDir() error {
	projectPath, err := getProjectPath()
	if err != nil {
		return err
	}
	return os.MkdirAll(projectPath, 0700)
}

func getDBPath() (string, error) {
	projectPath, err := getProjectPath()
	if err != nil {
		return "", err
	}
	dbPath := filepath.Join(projectPath, dbFileName)
	return dbPath, nil
}

func getProjectPath() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	projectPath := filepath.Join(configPath, projectName)
	return projectPath, nil
}
