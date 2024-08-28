package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestForLoopStatement(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println("Counter", i)
	}
}
