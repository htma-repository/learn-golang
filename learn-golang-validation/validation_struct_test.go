package learn_golang_validation

import "testing"

type ProductRequest struct {
	Name        string `validate:"required,min=1,max=100"`
	Description string `validate:"required,min=1,max=500"`
	Price       int    `validate:"required,numeric"`
}

func TestXxx(t *testing.T) {

	newProduct := &ProductRequest{
		Name:        "Product-1",
		Description: "Description-1",
		Price:       100000,
	}

	err := validate.Struct(newProduct)

	if err != nil {
		t.Error(err.Error())
	}
}
