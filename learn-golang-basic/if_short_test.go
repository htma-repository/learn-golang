package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestIfShort(t *testing.T) {
	if gamePrice := 350000; gamePrice > 200000 {
		fmt.Println("Game is expensive")
	} else {
		fmt.Println("Game is affordable")
	}
}
