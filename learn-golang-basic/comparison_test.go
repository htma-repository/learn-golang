package main

import (
	"fmt"
	"testing"
)

func TestComparison(t *testing.T) {
	num1 := 50
	num2 := 60

	isEqual := num1 == num2
	isNotEqual := num1 != num2
	isGreater := num1 > num2
	isGreaterThan := num1 >= num2
	isSmaller := num1 < num2
	isSmallerThan := num1 <= num2

	fmt.Println(isEqual, isNotEqual, isGreater, isGreaterThan, isSmaller, isSmallerThan)
}
