package strings

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
