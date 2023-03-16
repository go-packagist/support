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

// Sha1 returns the sha1 hash of a string.
func (s String) Sha1() string {
	return Sha1(string(s))
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

// StrShuffle returns a string with the characters in reverse order.
func (s String) StrShuffle() string {
	return StrShuffle(string(s))
}

// Atoi returns the integer value of a string.
func (s String) Atoi() AtoiRes {
	return Atoi(string(s))
}

// StrPad returns a string padded to a certain length with another string.
func (s String) StrPad(length int, pad string, padType StrPadType) string {
	return StrPad(string(s), length, pad, padType)
}

// Bytes returns the bytes of the string.
func (s String) Bytes() []byte {
	return []byte(s)
}

// Length returns the length of the string.
func (s String) Length() int {
	return Length(string(s))
}

// Strcut returns part of a string.
func (s String) Strcut(start, length int) string {
	return Strcut(string(s), start, length)
}

// Limit returns part of a string.
func (s String) Limit(length int, suffix ...string) string {
	return Limit(string(s), length, suffix...)
}
