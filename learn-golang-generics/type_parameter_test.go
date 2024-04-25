package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Sum[T int | float64](num1, num2 T) T {
	return num1 + num2
}

// TestFuncGenerics is a test function that tests the Sum function with different types of arguments.
// The function takes no parameters and does not return any values.
func TestFuncGenerics(t *testing.T) {
	result := Sum(10, 20)
	result2 := Sum(10.1, 20.4)

	assert.Equal(t, result, result)
	assert.NotEqual(t, result, result2)
}
