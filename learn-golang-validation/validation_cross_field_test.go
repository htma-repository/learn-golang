package learn_golang_validation

import "testing"

type SignUp struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=8,max=32"`
	ConfirmPassword string `validate:"required,min=8,max=32,eqfield=Password"`
}

func TestStructCrossField(t *testing.T) {
	signupUser := &SignUp{
		Email:           "test@gmail.com",
		Password:        "Test123!",
		ConfirmPassword: "Test123!",
	}

	err := validate.Struct(signupUser)

	if err != nil {
		t.Error(err.Error())
	}

}
