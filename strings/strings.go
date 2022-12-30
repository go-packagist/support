package strings

import (
	"regexp"
	_s "strings"
)

// Is returns true if the value matches the pattern.
// The pattern can contain the wildcard character *.
// example: Is("*.example.com", "www.example.com") // true
// example: Is("*.example.com", "example.com") // false
func Is(pattern, value string) bool {
	if pattern == value {
		return true
	}

	pattern = _s.ReplaceAll(pattern, "*", ".*")

	match, err := regexp.Match(pattern, []byte(value))
	if err != nil {
		return false
	}

	return match
}
