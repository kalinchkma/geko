package library

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generate a ECDSA private key
func GenerateECDSAKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func GenerateJWT(data string, validity time.Duration, privateKey *ecdsa.PrivateKey) (string, error) {
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
	return token.SignedString(privateKey)
}

func ValidateJWT(tokenString string, publicKey *ecdsa.PublicKey) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
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
