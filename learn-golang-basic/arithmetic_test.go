package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Arithmetic struct {
	Name   string
	Result int
}

func TestArithmetic(t *testing.T) {
	num1 := 100
	num2 := 350

	result := []Arithmetic{
		{
			Name:   "+",
			Result: num1 + num2,
		},
		{
			Name:   "-",
			Result: num1 - num2,
		},
		{
			Name:   "*",
			Result: num1 * num2,
		},
		{
			Name:   "/",
			Result: num1 / num2,
		},
		{
			Name:   "%",
			Result: num1 % num2,
		},
	}

	for _, test := range result {
		t.Run(test.Name, func(t *testing.T) {
			red := test.Result
			assert.Equal(t, test.Result, red)
		})

	}

}
