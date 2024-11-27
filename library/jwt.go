package library

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenSecret = "super secret"

func GenerateJWT(data string, validity time.Duration) (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(validity).Unix(), // validity duration
		"iat":  time.Now().Unix(),               // Issue time
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Sign the token with the secret key
	// @TODO add strong token secrect on
	return token.SignedString(tokenSecret)
}

func ValidateJWT(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return tokenSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Validate token and extracts claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract data from claims
		data := claims["data"].(string)
		return data, nil
	}

	return "", errors.New("invalid token")
}
