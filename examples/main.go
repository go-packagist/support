package main

import (
	"github.com/go-packagist/support/_strings"
)

func main() {
	println(_strings.InArray("abc", []string{"abc", "def"}))
	println(_strings.Is("ab*", "abc"))
}
