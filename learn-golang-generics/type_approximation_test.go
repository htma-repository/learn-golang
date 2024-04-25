package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Digit int64

type DigitType interface {
	int | int8 | int16 | int32 | ~int64
}

func SumNum[T DigitType](num1, num2 T) T {
	return num1 + num2
}

func TestApproximation(t *testing.T) {
	assert.Equal(t, Digit(30), SumNum[Digit](10, 20))
}
