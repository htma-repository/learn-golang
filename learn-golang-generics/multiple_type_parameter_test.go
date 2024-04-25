package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User[X string, Y int] struct {
	name X
	age  Y
}

func GetUser[T string, U int](name T, age U) User[T, U] {
	user := User[T, U]{
		name: name,
		age:  age,
	}
	return user
}

func TestMultipleT(t *testing.T) {
	user := GetUser("joko", 17)

	assert.Equal(t, user.name, "joko")
}
