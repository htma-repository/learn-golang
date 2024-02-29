package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var names [2]string
	names[0] = "Hutama"
	names[1] = "Trirahmanto"

	fmt.Println(names)

	var arrNum = [3]int{
		10, 20, 30,
	}

	fmt.Println(arrNum[0] + arrNum[1] + arrNum[2])

	var arrNumDynamic = [...]int{
		10, 20, 30, 40,
	}

	fmt.Println(arrNumDynamic)
	fmt.Println(len(arrNumDynamic))
	fmt.Println(arrNumDynamic[2])

	arrNumDynamic[0] = 15
	fmt.Println(arrNumDynamic)
}
