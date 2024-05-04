package learn_golang_validation

import "testing"

func TestParameter(t *testing.T) {

	password := "Rahmanto999@"

	err := validate.Var(password, "required,containsany=1234567890!@#$%^&*,min=6,max=32")
	if err != nil {
		t.Error(err.Error())
	}
}
