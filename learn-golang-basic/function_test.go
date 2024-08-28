package learn_golang_basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func greeting() {
	fmt.Println("Hello Hutama Trirahmanto")
}

func sumValue(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func multiplyValue(num1, num2 int) int {
	return num1 * num2
}
func TestFunc(t *testing.T) {

	greeting()
	fmt.Println(sumValue(5, 10))
	fmt.Println(multiplyValue(5, 10))

	assert.Equal(t, 10, sumValue(5, 5))
}
