package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	name    T
	address T
}

func (d *Data[T]) AddName(name T) {
	d.name = name
}

func (d *Data[T]) AddAddress(address T) {
	d.address = address
}

func (d *Data[T]) GetAll() (T, T) {
	return d.name, d.address
}

func TestGenericsStruct(t *testing.T) {
	data := Data[string]{}

	data.AddName("Joko")
	data.AddAddress("jakarta")

	name, address := data.GetAll()

	assert.Equal(t, data.name, name)
	assert.Equal(t, data.address, address)
}
