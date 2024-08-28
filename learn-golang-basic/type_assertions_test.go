package learn_golang_basic

import (
	"fmt"
	"learn-golang-basic/helper"
	"testing"
)

func TestTypeAssertions(t *testing.T) {
	// result := typeAssert()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	// With Switch
	switch value := helper.TypeAssert(); value.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Bool", value)
	default:
		fmt.Println("Unknown", value)
	}

}
