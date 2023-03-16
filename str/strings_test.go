package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	assert.True(t, Is("abc", "abc"))
	assert.False(t, Is("abcc", "abc"))
	assert.True(t, Is("ab*", "abc"))
	assert.True(t, Is("ab*", "ab"))
	assert.True(t, Is("ab/*", "ab/cc"))
	assert.True(t, Is("ab/*", "ab/"))
	assert.True(t, Is("*", "ab"))
	assert.True(t, Is("*", ""))
	assert.True(t, Is("ab/*", "ab/cc/dd"))
	assert.True(t, Is("*dd/", "ab/cc/dd/"))
	assert.False(t, Is("*dd/d", "dd/"))
}

func TestInArray(t *testing.T) {
	assert.True(t, InArray("1", []string{"1", "2"}))
	assert.True(t, InArray("2", []string{"1", "2"}))
	assert.False(t, InArray("3", []string{"1", "2"}))
	assert.False(t, InArray("12", []string{"1", "2"}))
}

func TestMd5(t *testing.T) {
	assert.Equal(t, "900150983cd24fb0d6963f7d28e17f72", Md5("abc"))
}

func TestStrpos(t *testing.T) {
	assert.Equal(t, 0, Strpos("aabbcc", "a"))
	assert.Equal(t, 2, Strpos("aabbcc", "b"))
	assert.Equal(t, -1, Strpos("aabbcc", "d"))
}

func TestStrrpos(t *testing.T) {
	assert.Equal(t, 1, Strrpos("aabbcc", "a"))
	assert.Equal(t, 3, Strrpos("aabbcc", "b"))
	assert.Equal(t, -1, Strrpos("aabbcc", "d"))
}

func TestStrrev(t *testing.T) {
	assert.Equal(t, "cba", Strrev("abc"))
}

func TestStrtr(t *testing.T) {
	assert.Equal(t, "bbbbcc", Strtr("aabbcc", "a", "b"))
}

// todo: to be fixed
// func TestStrtrArray(t *testing.T) {
// 	assert.Equal(t, "ddffcc", StrtrArray("aabbcc", map[string]string{"a": "d", "b": "f"}))
// 	assert.Equal(t, "ffffcc", StrtrArray("aabbcc", map[string]string{"a": "b", "b": "f"}))
// }

func TestStrShuffle(t *testing.T) {
	assert.True(t, InArray(StrShuffle("abc"), []string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}

func TestRandomString(t *testing.T) {
	r1, r2 := RandomString(10), RandomString(10)

	assert.Equal(t, 10, len(r1))
	assert.Equal(t, 10, len(r2))
	assert.NotEqual(t, r1, r2)
}
