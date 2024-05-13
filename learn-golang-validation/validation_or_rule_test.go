package learn_golang_validation

import "testing"

func TestOrRuleValidation(t *testing.T) {

	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := &Login{
		Username: "819038192038",
		Password: "password",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err)
	}

}
