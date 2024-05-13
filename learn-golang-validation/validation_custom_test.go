package learn_golang_validation

import (
	"testing"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func validationPassword(field validator.FieldLevel) bool {

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	value := field.Field().String()

	if len(value) > 6 {
		hasMinLen = true
	}

	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func TestValidationCustom(t *testing.T) {
	validate.RegisterValidation("password", validationPassword)

	type User struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,password"`
	}

	request := &User{
		Email:    "test@gmail.com",
		Password: "Password1!",
	}

	err := validate.Struct(request)

	if err != nil {
		t.Error(err.Error())
	}
}
