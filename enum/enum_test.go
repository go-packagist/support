package enum

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	e := New("Colors", map[string]int{
		"Red":   1,
		"Green": 2,
		"Blue":  3,
	})

	value, err := e.Value("Green")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Value of Green is %d\n", value)
	name, err := e.NameOf(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Name of value 2 is %s\n", name)
}
