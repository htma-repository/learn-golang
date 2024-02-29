package main

import (
	"fmt"
	"testing"
)

func fullNameFunc() (firstName, middleName, lastName string) {
	firstName = "Hutama"
	middleName = "Tri"
	lastName = "Rahmanto"

	return firstName, middleName, lastName
}

func TestFuncReturnNameVal(t *testing.T) {
	a, b, c := fullNameFunc()
	fmt.Println(a, b, c)
}
