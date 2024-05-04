package learn_golang_validation

import "testing"

func TestAliasTag(t *testing.T) {
	validate.RegisterAlias("varchar", "required,min=1,max=100")

	type User struct {
		Name    string `validate:"varchar"`
		Address string `validate:"varchar"`
	}

	newUser := &User{
		Name:    "John",
		Address: "",
	}

	err := validate.Struct(newUser)

	if err != nil {
		t.Error(err.Error())
	}
}
