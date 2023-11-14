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
