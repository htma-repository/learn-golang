package main

import "fmt"

func main() {
	if gamePrice := 350000; gamePrice > 200000 {
		fmt.Println("Game is expensive")
	} else {
		fmt.Println("Game is affordable")
	}
}
