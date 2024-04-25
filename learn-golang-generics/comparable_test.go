package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isEqual[T comparable](val1, val2 T) bool {
	if val1 == val2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	result := isEqual(1, 1)
	result2 := isEqual(10, 20)

	assert.True(t, result)
	assert.False(t, result2)
}
