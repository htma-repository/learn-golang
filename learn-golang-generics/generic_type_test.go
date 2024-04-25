package learn_golang_generics

import (
	"fmt"
	"testing"
)

type Names[T any] []T

func ShowName[U any](names Names[U]) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func TestGenericType(t *testing.T) {
	user := Names[string]{"hutama", "trirahmanto"}
	ShowName(user)
	number := Names[int]{1, 2, 3}
	ShowName(number)
}
