package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Is(t *testing.T) {
	assert.True(t, String("abc").Is("ab*"))
}

func TestString_InArray(t *testing.T) {
	assert.True(t, String("abc").InArray([]string{"abc", "def"}))
}

func TestString_Md5(t *testing.T) {
	assert.Equal(t, "900150983cd24fb0d6963f7d28e17f72", String("abc").Md5())
}

func TestString_Sha1(t *testing.T) {
	assert.Equal(t, "a9993e364706816aba3e25717850c26c9cd0d89d", String("abc").Sha1())
}

func TestString_Strpos(t *testing.T) {
	assert.Equal(t, 0, String("aabbcc").Strpos("a"))
}

func TestString_Strrpos(t *testing.T) {
	assert.Equal(t, 1, String("aabbcc").Strrpos("a"))
}

func TestString_Strrev(t *testing.T) {
	assert.Equal(t, "cba", String("abc").Strrev())
}

func TestString_Strtr(t *testing.T) {
	assert.Equal(t, "bbbbcc", String("aabbcc").Strtr("a", "b"))
}

func TestString_StrShuffle(t *testing.T) {
	assert.True(t, String("abc").InArray([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}

func TestString_Atoi(t *testing.T) {
	assert.Equal(t, 1, String("1").Atoi().Val())
	assert.Equal(t, 2, String("2").Atoi().Val())
	assert.Error(t, String("a").Atoi().Err())
	assert.Error(t, String("1a").Atoi().Err())
	assert.True(t, String("1").Atoi().IsOk())
	assert.False(t, String("a").Atoi().IsOk())
	assert.False(t, String("1").Atoi().IsErr())
	assert.True(t, String("a").Atoi().IsErr())
}
