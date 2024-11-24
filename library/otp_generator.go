package library

import (
	"fmt"
	"math/rand"
)

func GenerateOTP(length int) string {
	otp := ""
	for i := 0; i < length; i++ {

		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}
