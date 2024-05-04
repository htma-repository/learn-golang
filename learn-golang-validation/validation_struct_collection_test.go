package learn_golang_validation

import "testing"

type Employee struct {
	Name    string    `validate:"required"`
	Age     string    `validate:"required"`
	Address []Address `validate:"required,dive"`
}

func TestCollection(t *testing.T) {
	employee := &Employee{
		Name: "",
		Age:  "",
		Address: []Address{
			{
				City:     "",
				Province: "",
			},
			{
				City:     "",
				Province: "",
			},
		},
	}

	err := validate.Struct(employee)

	if err != nil {
		t.Error(err.Error())
	}

}
