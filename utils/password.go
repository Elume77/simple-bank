package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashedPassword retuns the bcrypt hash password string
func HashedPassword(password string) (string, error) {

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedpassword), nil
}

// CheckPassword checks if password is correct or not
func CheckPassword(password string, hashedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}
