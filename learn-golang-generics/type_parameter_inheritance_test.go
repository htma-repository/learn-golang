package learn_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Users interface {
	GetEmail() string
}

func GetEmail[T Users](param T) string {
	return param.GetEmail()
}

// create customer interface
type Customer interface {
	GetEmail() string
	GetCustomerName() string
}

// create customers struct that implement customer interface
type Customers struct {
	name  string
	email string
}

func (customer *Customers) GetEmail() string {
	return customer.email
}

func (customer *Customers) GetCustomerName() string {
	return customer.name
}

// create admin interface
type Admin interface {
	GetEmail() string
	GetAdminName() string
}

// create admins struct that implement admin interface
type Admins struct {
	name  string
	email string
}

func (admin *Admins) GetEmail() string {
	return admin.email
}

func (admin *Admins) GetAdminName() string {
	return admin.name
}

func TestTypeParameterInheritance(t *testing.T) {
	assert.Equal(t, "customer@gmail.com", GetEmail[Customer](&Customers{email: "customer@gmail.com"}))
	assert.NotEqual(t, "customer@gmail.com", GetEmail[Admin](&Admins{email: "admin@gmail.com"}))
}
