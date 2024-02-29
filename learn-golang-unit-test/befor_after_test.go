package main

import (
	"fmt"
	"testing"
)

// This TestMain before & after test only run in this package, other package will not run this TestMain, must create new TestMain if need to add before after test run, and function name MUST TestMain to this function work

func TestMain(m *testing.M) {
	fmt.Println("Before test run")

	m.Run()

	fmt.Println("After test run")
}
