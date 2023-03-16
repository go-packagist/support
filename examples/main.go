package main

import (
	"github.com/go-packagist/support/str"
)

func main() {
	// strings
	println(str.InArray("abc", []string{"abc", "def"}))
	println(str.Is("ab*", "abc"))
	println(str.Md5("abc"))
	println(str.Strpos("aabbcc", "a"))
	println(str.Strrpos("aabbcc", "a"))
	println(str.Strrev("abc"))
	println(str.Strtr("aabbcc", "a", "b"))
	println(str.StrShuffle("abc"))
	println(str.StrPad("abc", 6, " ", str.StrPadLeft))
	println(str.Length("张三")) // 2
	println(str.Strcut("abc", 0, 1))
	println(str.Limit("abc", 1, "..."))

	// string Atoi
	println(str.Atoi("1").Val())
	println(str.Atoi("a").Err())
	println(str.Atoi("a").IsOk())

	// Type String
	println(str.String("abc").Is("ab*"))
	println(str.String("abc").InArray([]string{"abc", "def"}))
	println(str.String("abc").Md5())
	println(str.String("abc").Sha1())
	println(str.String("aabbcc").Strpos("a"))
	println(str.String("aabbcc").Strrpos("a"))
	println(str.String("abc").Strrev())
	println(str.String("aabbcc").Strtr("a", "b"))
	println(str.String("abc").StrShuffle())
	println(str.String("1").Atoi().Val())
	println(str.String("a").Atoi().Err())
	println(str.String("a").Atoi().IsOk())
	println(str.String("abc").Bytes())
	println(str.String("abc").StrPad(6, " ", str.StrPadLeft))

}
