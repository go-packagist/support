package main

import (
	"github.com/go-packagist/support/str"
)

func main() {
	println(str.InArray("abc", []string{"abc", "def"}))
	println(str.Is("ab*", "abc"))
	println(str.Md5("abc"))
	println(str.Strpos("aabbcc", "a"))
	println(str.Strrpos("aabbcc", "a"))
	println(str.Strrev("abc"))
	println(str.Strtr("aabbcc", "a", "b"))
	println(str.StrShuffle("abc"))

	println(str.Atoi("1").Val())
	println(str.Atoi("a").Err())
	println(str.Atoi("a").IsOk())
}
