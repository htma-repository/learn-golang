package learn_golang_std_lib

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {

	firstName := "Hutama"
	lastName := "Trirahmanto"
	age := 25

	fmt.Printf("Hello %s %s, age %d\n", firstName, lastName, age)
	fmt.Println(firstName, lastName, age)
}
