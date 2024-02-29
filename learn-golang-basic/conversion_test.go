package main

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	var fullName string = "Hutama Trirahmanto"
	var indexVal uint8 = fullName[0]
	var indexValStr string = string(indexVal)

	fmt.Println(fullName)
	fmt.Println(indexVal)
	fmt.Println(indexValStr)

	var num32 int32 = 32770
	var num64 int64 = int64(num32)
	var num16 int16 = int16(num32) // overflow

	fmt.Println(num32)
	fmt.Println(num64)
	fmt.Println(num16)

}
