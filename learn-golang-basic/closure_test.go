package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {

	// outerFunc := func() func() {
	// 	counter := 0

	// 	innerFunc := func() {
	// 		counter++
	// 		fmt.Println("Inner Func", counter)
	// 	}

	// 	return innerFunc
	// }

	// outerFunc()()
	// outerFunc()()
	// outerFunc()()

	counter := 0

	counterFunc := func() {
		fmt.Println("Counter", counter)
		counter++
	}

	counterFunc()
	counterFunc()
	counterFunc()

	fmt.Println(counter)
}
