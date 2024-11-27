package library

import "golang.org/x/crypto/bcrypt"

// @TODO implement a password hash method
func HashPassword(passwordString string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// @TODO implement a compare password method
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
