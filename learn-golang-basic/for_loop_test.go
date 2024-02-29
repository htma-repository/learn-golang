package main

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	counter := 0
	for counter <= 10 {
		fmt.Println("Counter", counter)
		counter++
	}
}
