package learn_golang_validation

import "testing"

type Address struct {
	City     string `validate:"required"`
	Province string `validate:"required"`
}

type User struct {
	Name    string  `validate:"required"`
	Age     string  `validate:"required"`
	Address Address `validate:"required"`
}

func TestNestStruct(t *testing.T) {
	user := &User{
		Name: "a",
		Age:  "a",
		Address: Address{
			City:     "a",
			Province: "a",
		},
	}

	err := validate.Struct(user)

	if err != nil {
		t.Error(err.Error())
	}

}
