package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash user password
func HashPassword(passwordString string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Compate hashed password
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
