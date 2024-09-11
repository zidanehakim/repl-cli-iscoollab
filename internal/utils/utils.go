package utils

import (
	"regexp"
)

func ValidateString(str string) bool {
	// Alphanumeric characters, dots, underscores, and hyphens, within 1-255 characters
	regex := regexp.MustCompile(`^[a-zA-Z0-9._-]{1,255}$`)
	return regex.MatchString(str)
}

func ParseInput(input string) []string {
	// Nongreedy or sequence non-whitespace chars regex
	re := regexp.MustCompile(`"(.*?)"|\S+`)
	matches := re.FindAllStringSubmatch(input, -1)

	var args []string
	for _, match := range matches {
		if len(match) > 1 && match[1] != "" {
			args = append(args, match[1])
		} else {
			args = append(args, match[0])
		}
	}

	return args
}
