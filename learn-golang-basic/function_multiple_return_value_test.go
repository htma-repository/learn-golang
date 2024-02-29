package main

import (
	"fmt"
	"testing"
)

func multipleValue(num1 int, num2 int) (string, int, int, int) {
	result := num1 * num2
	return "Multiply", num1, num2, result
}

func TestFuncMultiReturnVal(t *testing.T) {

	strings, firstNum, secondNum, result := multipleValue(2, 10)
	fmt.Println(strings, firstNum, secondNum, result)
	_, firstNumber, secondNumber, resultValue := multipleValue(3, 10)
	fmt.Println(firstNumber, secondNumber, resultValue)

}
