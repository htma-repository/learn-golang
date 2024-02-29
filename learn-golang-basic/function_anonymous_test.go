package main

import (
	"fmt"
	"learn-golang-basic/helper"
	"testing"
)

func getHello(name string, filterFunc helper.Filter) string {
	result := filterFunc(name)

	return "Hello " + result
}

func TestFuncAnonymous(t *testing.T) {
	getHelloUser := getHello("Trirahmanto", func(name string) string {
		if name == "Bego" {
			return "..."
		}
		return name
	})

	fmt.Println(getHelloUser)

	// Anonymous function
	spamFunc := func(name string) string {
		if name == "Bego" {
			return "..."
		}

		return name
	}

	fmt.Println(getHello("Bego", spamFunc))
}
