package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee[T any] interface {
	AddEmployee(firstName T)
	GetEmployeeName() T
}

func GetFullName[T any](employee Employee[T], name T) T {
	employee.AddEmployee(name)
	return employee.GetEmployeeName()
}

type Developer[T any] struct {
	FirstName T
}

func (d *Developer[T]) AddEmployee(firstName T) {
	d.FirstName = firstName
}

func (d *Developer[T]) GetEmployeeName() T {
	return d.FirstName
}

func TestGenericsInterface(t *testing.T) {
	developer1 := &Developer[string]{}
	// developer1.AddEmployee("Hutama")
	result := GetFullName(developer1, "Hutama")

	assert.Equal(t, result, "Hutama")
	assert.Equal(t, result, developer1.FirstName)
}
