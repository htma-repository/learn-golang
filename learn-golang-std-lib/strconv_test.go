package learn_golang_std_lib

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {

	resultBool, err := strconv.ParseBool("true") // only convert "true" / "false" to boolean

	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(resultBool, reflect.TypeOf(resultBool))
	}

	resultInt, err2 := strconv.Atoi("1000")

	if err2 != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(resultInt, reflect.TypeOf(resultInt))
	}

	resultStr := strconv.Itoa(1000)
	fmt.Println(resultStr, reflect.TypeOf(resultStr))

	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool, reflect.TypeOf(formatBool))
}
