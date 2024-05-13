package learn_golang_validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type Users struct {
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	Age     uint   `validate:"required"`
	Address string `validate:"required"`
}

func ValidationUser(value validator.StructLevel) {

	structUser := value.Current().Interface().(Users)

	if structUser.Name == structUser.Email {
		// success
	} else {
		value.ReportError(structUser.Name, "Name", "Name", "name", "")
	}
}

func TestValidationStructLevel(t *testing.T) {

	validate.RegisterStructValidation(ValidationUser, Users{})

	request := &Users{
		Name:    "John@gmail.com",
		Email:   "John@gmail.com",
		Age:     30,
		Address: "USA",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err)
	}
}
