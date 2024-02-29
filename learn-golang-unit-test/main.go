package main

import "fmt"

func sumNumber(num1, num2 int) int {
	return num1 + num2
}

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	result := sumNumber(10, 20)
	fmt.Println(result)

	s := sayGoodbye("Tama")
	fmt.Println(s)
}
