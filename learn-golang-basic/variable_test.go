package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	name := "Hutama Trirahmanto" // := used only on very first initialization variable
	fmt.Println(name)

	name = "Hutama"
	fmt.Println(name)

	var (
		firstName = "Hutama"
		lastName  = "Trirahmanto"
	)

	fmt.Println(firstName + " " + lastName)

	const number = 100
	const numberStr = 10

	fmt.Println(number + numberStr)
}
