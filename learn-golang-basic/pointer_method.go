package main

import (
	"fmt"
	"learn-golang-basic/helper"
)

func main() {
	human1 := helper.Human{Name: "Tama", Address: "Cipondoh"}
	human1.SayHello()

	fmt.Println(human1)
}
