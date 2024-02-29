package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	TestName  string
	Parameter string
	Expected  string
	Message   string
}

func TestAll(t *testing.T) {
	testData := []TestConfig{
		{
			TestName:  "sayGoodbye Hutama",
			Parameter: "Hutama",
			Expected:  "Goodbye Hutama",
			Message:   "Result must be 'Goodbye Hutama'",
		},
		{
			TestName:  "sayGoodbye Trirahmanto",
			Parameter: "Trirahmanto",
			Expected:  "Goodbye Trirahmanto",
			Message:   "Result must be 'Goodbye Trirahmanto'",
		},
	}

	for _, test := range testData {
		t.Run(test.TestName, func(t *testing.T) {
			result := sayGoodbye(test.Parameter)
			assert.Equal(t, test.Expected, result, test.Message)
		})
	}
}
