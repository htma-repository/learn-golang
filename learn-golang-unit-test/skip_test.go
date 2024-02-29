package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cant run test on linux")
	}

	result := sumNumber(10, 30)
	assert.Equal(t, 40, result, "Result must be 40")
}
