package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	bookPrice := 50000

	if bookPrice < 50000 {
		fmt.Println("Book is cheap")
	} else if bookPrice >= 50000 && bookPrice <= 100000 {
		fmt.Println("Book is affordable")
	} else {
		fmt.Println("Book is very expensive")
	}
}
