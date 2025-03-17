package lib

import (
	"os"
	"strconv"
)

func GetEnvString(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func GetEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intValue
}

func GetEnvFloat32(key string, fallback float32) float32 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	floatValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(err)
	}
	return float32(floatValue)
}

func GetEnvFloat64(key string, fallback float64) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return floatValue
}

func GetEnvBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return boolVal
	}

	return boolVal
}
