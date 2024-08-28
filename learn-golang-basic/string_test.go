package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("Hello, World!")
	fmt.Println("Hutama Trirahmanto!")
	fmt.Println(len("Hutama"))
	fmt.Println(string("Hutama"[0]))
}
