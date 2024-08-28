package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestBreakContinue(t *testing.T) {
	for i := 1; i < 10; i++ {
		if i == 7 {
			break
		}

		fmt.Println("Counter", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Genap", i)
	}
}
