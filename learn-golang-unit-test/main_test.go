package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Manual without assert (Not Recommended)
// func TestSumNumber(t *testing.T) {
// 	result := sumNumber(10, 10)
// 	if result != 20 {
// 		t.Fatal("Result Must be 20")
// 	}
// }

// func TestSayGoodbye(t *testing.T) {
// 	result := sayGoodbye("Hutama")
// 	if result != "Goodbye Hutama" {
// 		t.Fatal("Result must be Goodbye Hutama")
// 	}
// }

const successMessage = "Success test"

// With assert (Recommended)
func TestSumNumberAssert(t *testing.T) {
	result := sumNumber(10, 30)
	assert.Equal(t, 40, result, "Result must be 40")
	assert.NotEqual(t, 50, result, "Result must be 40")
	fmt.Println(successMessage)
}

func TestSayGoodbyeAssert(t *testing.T) {
	result := sayGoodbye("Hutama")
	assert.Equal(t, "Goodbye Hutama", result, "Result must be 'Goodbye Hutama'")
	fmt.Println(successMessage)
}

func TestSumNumberRequire(t *testing.T) {
	result := sumNumber(20, 30)
	require.Equal(t, 50, result, "Result must be 50")
	fmt.Println(successMessage)
}

func TestSayGoodbyeRequire(t *testing.T) {
	result := sayGoodbye("Joko")
	require.Equal(t, "Goodbye Joko", result, "Result must be 'Goodbye Joko'")
	fmt.Println(successMessage)
}
