package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeInference(t *testing.T) {
	assert.Equal(t, 40, Max(40, 10))
	assert.Equal(t, int64(300), Max(int64(100), int64(300)))
}
