package strings

import (
	"regexp"
	str "strings"
)

func Is(pattern, value string) bool {
	if pattern == value {
		return true
	}

	pattern = str.ReplaceAll(pattern, "*", ".*")

	match, err := regexp.Match(pattern, []byte(value))
	if err != nil {
		return false
	}

	return match
}
