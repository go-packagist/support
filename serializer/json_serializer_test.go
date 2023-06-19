package serializer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonSerializer_String(t *testing.T) {
	serialized, err := Json.Serialize("foo")
	assert.Nil(t, err)
	assert.Equal(t, "\"foo\"", string(serialized))

	var unserialize string
	assert.Nil(t, Json.Unserialize(serialized, &unserialize))
	assert.Equal(t, "foo", unserialize)
}

func TestJsonSerializer_Int(t *testing.T) {
	serialized, err := Json.Serialize(1)
	assert.Nil(t, err)
	assert.Equal(t, "1", string(serialized))

	var unserialize int
	assert.Nil(t, Json.Unserialize(serialized, &unserialize))
	assert.Equal(t, 1, unserialize)
}

func TestJsonSerializer_Struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	user := &User{"Flc", 18}

	serialized, err := Json.Serialize(user)
	assert.Nil(t, err)
	fmt.Println(string(serialized))
	assert.Equal(t, "{\"Name\":\"Flc\",\"Age\":18}", string(serialized))

	var unserialize *User
	assert.Nil(t, Json.Unserialize(serialized, &unserialize))
	assert.Equal(t, "Flc", unserialize.Name)
}
