package learn_golang_basic

import (
	"fmt"
	"testing"
)

func getUserName(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func TestNil(t *testing.T) {

	data := getUserName("")

	if data == nil {
		fmt.Println("Data nil")
	} else {
		fmt.Println("Data available", data)
	}
}
