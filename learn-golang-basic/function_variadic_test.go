package learn_golang_basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func multiplyAll(numbers ...int) int {
	result := 1

	for _, number := range numbers {
		result *= number
	}

	return result
}

func TestFuncVariadic(t *testing.T) {
	fmt.Println(multiplyAll(12, 12, 12, 12))
	fmt.Println(multiplyAll(23, 14, 56, 68))

	numList := []int{10, 10, 10, 10}

	fmt.Println(multiplyAll(numList...))

	assert.Equal(t, 10000, multiplyAll(numList...))
}
