package learn_golang_validation

import "testing"

func TestValidationTwoVar(t *testing.T) {

	password := "password"
	confirmPassword := "password"

	err := validate.VarWithValue(password, confirmPassword, "eqfield,required")

	if err != nil {
		t.Error(err.Error())
	}
}
