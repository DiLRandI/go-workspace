package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a, b := 2, 3
	expected := 5
	result := Add(a, b)
	assert.Equal(t, expected, result)
}

func TestSub(t *testing.T) {
	a, b := 2, 3
	expected := -1
	result := Sub(a, b)
	assert.Equal(t, expected, result)
}

func TestMul(t *testing.T) {
	a, b := 2, 3
	expected := 6
	result := Mul(a, b)
	assert.Equal(t, expected, result)
}
