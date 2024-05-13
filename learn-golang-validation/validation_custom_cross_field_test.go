package learn_golang_validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func validationUsername(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()

	if !ok {
		panic("not ok")
	}

	value2 := field.Field().String()

	//for this example case, value is value from other field from parameter, value2 is value from field itself
	fmt.Println("value", value.String(), "value2", value2)

	valueToLower := strings.ToLower(value.String())
	value2ToLower := strings.ToLower(value2)

	return valueToLower == value2ToLower
}

func TestValidationCustomCrossField(t *testing.T) {

	if err := validate.RegisterValidation("ignore_case", validationUsername); err != nil {
		t.Error(err)
	}

	type Register struct {
		Username string `validate:"required,ignore_case=Email|ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Password string `validate:"required"`
	}

	request := &Register{
		Username: "0812345677",
		Email:    "test@gmail.com",
		Phone:    "0812345677",
		Password: "password",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err)
	}

}
