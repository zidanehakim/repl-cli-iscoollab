package utils

import (
	"regexp"
)

func ValidateString(str string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return regex.MatchString(str)
}
