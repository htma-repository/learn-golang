package learn_golang_validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationError(t *testing.T) {

	product := ProductRequest{
		Name:        "product-1",
		Description: "",
		Price:       0,
	}

	err := validate.Struct(product)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			t.Errorf("error field: %v, error tag: %v, error message: %v", fieldError.Field(), fieldError.Tag(), fieldError.Error())
		}
	}

}
