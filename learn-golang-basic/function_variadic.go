package main

import "fmt"

func multiplyAll(numbers ...int) int {
	result := 1

	for _, number := range numbers {
		result *= number
	}

	return result
}

func main() {
	fmt.Println(multiplyAll(12, 12, 12, 12))
	fmt.Println(multiplyAll(23, 14, 56, 68))

	numList := []int{10, 10, 10, 10}

	fmt.Println(multiplyAll(numList...))
}
