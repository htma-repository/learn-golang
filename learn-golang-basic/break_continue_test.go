package main

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
		if i%2 == 1 {
			continue
		}
		fmt.Println("Genap", i)
	}
}
