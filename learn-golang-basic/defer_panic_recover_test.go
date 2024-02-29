package main

import (
	"fmt"
	"testing"
)

func finishedFunc() {
	fmt.Println("Finished")
	errMessage := recover()
	fmt.Println("Error", errMessage)
}

func blockedUser(user string) bool {
	blockedList := []string{"Alfin", "Riu", "Ayub"}
	var result bool

	for _, name := range blockedList {
		if user == name {
			result = true
			break
		} else {
			result = false
		}
	}

	return result
}

func newUser(name string, isError bool) string {
	defer finishedFunc()
	if blockedUser(name) && !isError {
		return "Your blocked!"
	} else if isError {
		panic("Error Occurred")
	} else {
		return "Welcome " + name
	}
}

func TestDeferPanicRecover(t *testing.T) {
	fmt.Println(newUser("Hutama", false))
	fmt.Println(newUser("Riu", true))
	fmt.Println("Func Finished")
}
