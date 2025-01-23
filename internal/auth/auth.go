package auth

import (
	"crypto/rand"
	"fmt"

	"github.com/mshortcodes/sentry/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateToken() (string, error) {
	token := make([]byte, 16)
	if _, err := rand.Read(token); err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return fmt.Sprintf("%x", token), nil
}

func ValidateToken(db database.Client) (database.Token, error) {
	dbToken, err := db.GetToken()
	if err != nil {
		return database.Token{}, fmt.Errorf("invalid token")
	}

	return dbToken, nil
}
