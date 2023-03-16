package str

type String string

// Is checks if the string matches the pattern.
func (s String) Is(pattern string) bool {
	return Is(pattern, string(s))
}

// InArray checks if a string is in a string array.
func (s String) InArray(arr []string) bool {
	return InArray(string(s), arr)
}

// Md5 returns the md5 hash of a string.
func (s String) Md5() string {
	return Md5(string(s))
}

// Strpos returns the position of the first occurrence of a substring in a string.
func (s String) Strpos(substr string) int {
	return Strpos(string(s), substr)
}

// Strrpos returns the position of the last occurrence of a substring in a string.
func (s String) Strrpos(substr string) int {
	return Strrpos(string(s), substr)
}

// Strrev returns a string with the characters in reverse order.
func (s String) Strrev() string {
	return Strrev(string(s))
}

// Strtr returns a string with the characters in reverse order.
func (s String) Strtr(from, to string) string {
	return Strtr(string(s), from, to)
}

// Atoi returns the integer value of a string.
func (s String) Atoi() AtoiRes {
	return Atoi(string(s))
}
