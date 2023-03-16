package ints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInArray(t *testing.T) {
	assert.True(t, InArray(1, []int{1, 2, 3}))
	assert.False(t, InArray(4, []int{1, 2, 3}))
}

func TestItoa(t *testing.T) {
	assert.Equal(t, "1", Itoa(1))
	assert.Equal(t, "2", Itoa(2))
}

func TestSplit(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Split("1,2,3", ","))
}

func TestRandom(t *testing.T) {
	r1, r2 := Random(1, 1000), Random(1, 1000)
	println(r1, r2)

	assert.True(t, r1 >= 1 && r1 <= 1000)
	assert.True(t, r2 >= 1 && r2 <= 1000)
	assert.NotEqual(t, r1, r2)
}

func TestRandomString(t *testing.T) {
	s1, s2 := RandomString(10), RandomString(10)
	println(s1, s2)

	assert.Equal(t, 10, len(s1))
	assert.Equal(t, 10, len(s2))
	assert.NotEqual(t, s1, s2)
}

func TestRange(t *testing.T) {
	assert.Equal(t, []int{1}, Range(1, 1))
	assert.Equal(t, []int{1, 2, 3}, Range(1, 3))
	assert.Equal(t, []int{1, 2, 3, 4}, Range(1, 4))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2, 3))
	assert.Equal(t, 1, Min(1, 2, 3, 4))
	assert.Equal(t, 1, Min(1, 2, 3, 4, 5))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max(1, 2, 3))
	assert.Equal(t, 4, Max(1, 2, 3, 4))
	assert.Equal(t, 5, Max(1, 2, 3, 4, 5))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 6, Sum(1, 2, 3))
	assert.Equal(t, 10, Sum(1, 2, 3, 4))
	assert.Equal(t, 15, Sum(1, 2, 3, 4, 5))
}

func TestBetween(t *testing.T) {
	assert.True(t, Between(1, 1, 3))
	assert.True(t, Between(2, 1, 3))
	assert.True(t, Between(3, 1, 3))
	assert.False(t, Between(4, 1, 3))
}
