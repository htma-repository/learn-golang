package main

import "fmt"

func greeting() {
	fmt.Println("Hello Hutama Trirahmanto")
}

func sumValue(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func multiplyValue(num1, num2 int) int {
	return num1 * num1
}
func main() {

	greeting()
	fmt.Println(sumValue(5, 10))
	fmt.Println(multiplyValue(5, 10))

}
