package main

import (
	"fmt"
	"testing"
)

func TestFuncAsValue(t *testing.T) {
	multiply := multiplyAll
	numList := []int{10, 10, 10, 10}

	fmt.Println(multiply(numList...))
}
