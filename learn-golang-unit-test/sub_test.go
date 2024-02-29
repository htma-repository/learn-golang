package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Sum Test", func(t *testing.T) {
		result := sumNumber(10, 30)
		assert.Equal(t, 40, result, "Result must be 40")
		assert.NotEqual(t, 50, result, "Result must be 40")
	})

	t.Run("SayGoodbye Test", func(t *testing.T) {
		result := sayGoodbye("Hutama")
		assert.Equal(t, "Goodbye Hutama", result, "Result must be 'Goodbye Hutama'")
	})

}
