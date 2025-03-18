package lib

import (
	"regexp"
	"strings"
)

// This function make any camelcase string to snake case
func CamelToSnake(s string) (string, error) {
	re, err := regexp.Compile("([a-z0-9])([A-Z])")
	if err != nil {
		return "", err
	}
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake), nil
}
