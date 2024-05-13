package learn_golang_validation

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/go-playground/validator/v10"
)

var regexNum = regexp.MustCompile("^[0-9]+$")

func validationPIN(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())

	if err != nil {
		panic(err)
	}

	value := field.Field().String()

	if !regexNum.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestValidationCustomParameter(t *testing.T) {

	err := validate.RegisterValidation("pin", validationPIN)

	if err != nil {
		t.Error(err)
	}

	type Payment struct {
		Amount int    `validate:"required,numeric"`
		PIN    string `validate:"required,pin=6"`
	}

	request := &Payment{
		Amount: 100000,
		PIN:    "123456",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err)
	}

}
