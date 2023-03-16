package strs

import (
	"crypto/md5"
	"crypto/sha1"
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

// Sha1 returns the sha1 hash of a string.
//
// Example:
//
//	Sha1("abc") // a9993e364706816aba3e25717850c26c9cd0d89d
func Sha1(s string) string {
	sm := sha1.Sum([]byte(s))

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

// Shuffle returns a string with its characters in random order.
//
// Example:
//
//	Shuffle("abc") // "bca"
func Shuffle(s string) string {
	ss := strings.Split(s, "")

	rand.Shuffle(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})

	return strings.Join(ss, "")
}

// RandomString returns a random string with the specified length.
//
// Example:
//
//	RandomString(10) // "qujrlkhyqr"
func RandomString(length int) string {
	var letters = []rune(randomStringLetters)
	var lettersLength = len(letters)

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(lettersLength)]
	}

	return string(b)
}

// StrPad returns an input string padded on the left or right to specified length with pad string.
//
// Example:
//
//	StrPad("abc", 6, " ", StrPadLeft) // "   abc"
//	StrPad("abc", 6, " ", StrPadRight) // "abc   "
//	StrPad("abc", 6, " ", StrPadBoth) // " abc  "
func StrPad(s string, length int, pad string, padType StrPadType) string {
	if len(s) >= length {
		return s
	}

	switch padType {
	case StrPadLeft:
		return strings.Repeat(pad, length-len(s)) + s
	case StrPadRight:
		return s + strings.Repeat(pad, length-len(s))
	case StrPadBoth:
		length = length - len(s)
		paddingLeft := strings.Repeat(pad, (length)/2)
		paddingRight := strings.Repeat(pad, (length)/2+(length)%2)
		return paddingLeft + s + paddingRight
	default:
		return s
	}
}

// Length returns the length of a string.
// support chinese characters.
//
// Example:
//
//	Length("abc") // 3
//	Length("张三李四") // 4
func Length(s string) int {
	return len([]rune(s))
}

// Strcut returns a string with a specified length starting from a specified position.
// support chinese characters.
//
// Example:
//
//	Strcut("abc", 1, 1) // "b"
//	Strcut("张三李四", 1, 2) // "三李
func Strcut(s string, start, length int) string {
	sr := []rune(s)

	if len(sr) <= start {
		return ""
	}

	if len(sr) <= start+length {
		return string(sr[start:])
	}

	return string(sr[start : start+length])
}

// Limit returns a string with a specified length.
// support chinese characters.
//
// Example:
//
//	Limit("abc", 2) // "ab"
//	Limit("张三李四", 2) // "张三"
//	Limit("张三李四", 2, "...") // "张三..."
func Limit(s string, length int, suffix ...string) string {
	sr, sf := []rune(s), ""

	if len(sr) <= length {
		return s
	}

	if len(suffix) > 0 {
		sf = suffix[0]
	}

	return string(sr[:length]) + sf
}
