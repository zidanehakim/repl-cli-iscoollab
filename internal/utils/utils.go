package utils

import (
	"regexp"
	"strings"
)

func ValidateString(str string) bool {
	// Alphanumeric characters, dots, underscores, and hyphens, within 1-255 characters
	regex := regexp.MustCompile(`^(['"][a-zA-Z0-9._\s-]{1,255}['"]|[a-zA-Z0-9._-]{1,255})$`)
	return regex.MatchString(str)
}

func ParseInput(input string) []string {
	// Nongreedy or sequence non-whitespace chars regex
	regex := regexp.MustCompile(`["'](.*?)["']|\S+`)
	matches := regex.FindAllStringSubmatch(input, -1)

	var args []string
	for _, match := range matches {
		match[0] = strings.TrimSpace(match[0])
		if len(match) > 1 {
			match[1] = strings.TrimSpace(match[1])

		}

		if match[0][0] == '\'' || match[0][0] == '"' {
			// Check if the string contains spaces, if not then quotes isnt needed
			if !strings.Contains(match[1], " ") {
				args = append(args, match[1])
			} else {
				args = append(args, `"`+match[1]+`"`)
			}
		} else {
			args = append(args, match[0])
		}
	}

	return args
}
