package main

import (
	"fmt"
	"testing"
)

func factorialRecursive(number int) int {
	if number == 10 {
		return 10
	}
	return number * factorialRecursive(number+1)
}

func TestFuncRecursive(t *testing.T) {
	fmt.Println(factorialRecursive(1))
}
