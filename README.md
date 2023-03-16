# Support

[![Go Version](https://badgen.net/github/release/go-packagist/support/stable)](https://github.com/go-packagist/support/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/support)](https://pkg.go.dev/github.com/go-packagist/support)
[![codecov](https://codecov.io/gh/go-packagist/support/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/support)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/support)](https://goreportcard.com/report/github.com/go-packagist/support)
[![tests](https://github.com/go-packagist/support/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/support/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/support
```

## Usage

```go
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
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.