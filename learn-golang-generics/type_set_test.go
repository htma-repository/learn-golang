package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number](num1, num2 T) T {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

type String interface {
	string
}

func Word[T String](word1, word2 T) bool {
	if word1 == word2 {
		return true
	} else {
		return false
	}
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, 10, Max[int](10, 5))
	assert.Equal(t, int64(100), Max[int64](100, 50))
	assert.Equal(t, true, Word[string]("hutama", "hutama"))
}
