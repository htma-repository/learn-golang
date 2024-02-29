package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) []uint64 {
	// your code goes here
	var result []uint64

	for i := 0; i <= n; i++ {
		result = append(result, uint64(math.Pow(float64(2), float64(i))))
	}

	return result
}

func main() {
	fmt.Println(math.Floor(5.7))
	fmt.Println(math.Ceil(5.7))
	fmt.Println(math.Round(5.7))
	fmt.Println(math.Max(12.3, 12.1))
	fmt.Println(math.Min(12.3, 12.1))
	fmt.Println(math.Pow(12.3, 12.1))
}
