package enum

import (
	"fmt"
)

type Enum interface {
	Value(string) ValResult
	String(int) StrResult
}

type enumValue struct {
	name  string
	value int
}

type enumValues []enumValue

type enum struct {
	name   string
	values enumValues
}

type ValResult struct {
	val int
	err error
}

var _ Enum = &enum{}

func (e enum) Value(name string) ValResult {
	for _, v := range e.values {
		if v.name == name {
			return ValResult{v.value, nil}
		}
	}

	return ValResult{0, fmt.Errorf("value %s not found in enum %s", name, e.name)}
}

func (e enum) String(value int) (string, error) {
	for _, v := range e.values {
		if v.value == value {
			return v.name, nil
		}
	}
	return "", fmt.Errorf("value %d not found in enum %s", value, e.name)
}

func New(name string, values map[string]int) Enum {
	ev := make(enumValues, len(values))
	i := 0
	for k, v := range values {
		ev[i] = enumValue{name: k, value: v}
		i++
	}
	return &enum{name: name, values: ev}
}
