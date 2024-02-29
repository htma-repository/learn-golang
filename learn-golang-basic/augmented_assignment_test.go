package main

import (
	"fmt"
	"testing"
)

func TestAugmentedAssignment(t *testing.T) {
	num1 := 10
	num2 := 20

	num1 += num2
	num2 -= num1
	num1 *= num2
	num2 /= num1

	fmt.Println(num1, num2)
}
