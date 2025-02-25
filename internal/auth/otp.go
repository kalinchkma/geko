package auth

import (
	"fmt"
	"math/rand"
)

// Generate otp
func GenerateOTP(length int) string {
	otp := ""
	for range length {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}
