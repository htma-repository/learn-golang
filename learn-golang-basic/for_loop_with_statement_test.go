package main

import (
	"fmt"
	"testing"
)

func TestForLoopStatement(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println("Counter", i)
	}
}
