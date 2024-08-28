package learn_golang_basic

import (
	"fmt"
	"learn-golang-basic/helper"
	"testing"
)

func TestPointerMethod(t *testing.T) {
	human1 := helper.Human{Name: "Tama", Address: "Cipondoh"}
	human1.SayHello()

	fmt.Println(human1)
}
