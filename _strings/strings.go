package _strings

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

// Is returns true if the value matches the pattern.
// The pattern can contain the wildcard character *.
//
// Example:
//
//	Is("*.example.com", "www.example.com") // true
//	Is("*.example.com", "example.com") // false
func Is(pattern, value string) bool {
	if pattern == value {
		return true
	}

	pattern = strings.ReplaceAll(pattern, "*", ".*")

	match, err := regexp.Match(pattern, []byte(value))
	if err != nil {
		return false
	}

	return match
}

// InArray checks if a string is in a string array.
//
// Example:
//
//	InArray("abc", []string{"abc", "def"}) // true
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

// Md5 returns the md5 hash of a string.
//
// Example:
//
//	Md5("abc") // 900150983cd24fb0d6963f7d28e17f72
func Md5(s string) string {
	sm := md5.Sum([]byte(s))

	return fmt.Sprintf("%x", sm)
}

// Strpos returns the position of the first occurrence of a substring in a string.
//
// Example:
//
//	Strpos("abc", "b") // 1
func Strpos(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

// Strrpos returns the position of the last occurrence of a substring in a string.
//
// Example:
//
//	Strrpos("abc", "b") // 1
func Strrpos(haystack, needle string) int {
	return strings.LastIndex(haystack, needle)
}

// Strrev returns a string with its characters in reverse order.
//
// Example:
//
//	Strrev("abc") // "cba"
func Strrev(s string) string {
	var reversed string

	for _, v := range s {
		reversed = string(v) + reversed
	}

	return reversed
}

// Strtr replaces all occurrences of the search string with the replacement string.
//
// Example:
//
//	Strtr("aabbcc", "a", "b") // "bbbbcc"
func Strtr(s, from, to string) string {
	return strings.NewReplacer(from, to).Replace(s)
}

// StrtrArray replaces all occurrences of the search strings with the replacement strings.
//
// Example:
//
//	StrtrArray("aabbcc", map[string]string{"a": "d", "b": "f"} // "ddffcc"
//	StrtrArray("aabbcc", map[string]string{"a": "b", "b": "f"}) // "ffffcc" (wrong) todo: to be fixed
// func StrtrArray(s string, rp map[string]string) string {
//
// 	strings.ReplaceAll(s, "a", "b")
// 	for k, v := range rp {
// 		s = strings.ReplaceAll(s, k, v)
// 	}
//
// 	return s
// }

// StrShuffle returns a string with its characters in random order.
//
// Example:
//
//	StrShuffle("abc") // "bca"
func StrShuffle(s string) string {
	ss := strings.Split(s, "")

	rand.Shuffle(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})

	return strings.Join(ss, "")
}
