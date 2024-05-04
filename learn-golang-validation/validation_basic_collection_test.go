package learn_golang_validation

import "testing"

type Manager struct {
	Name    string    `validate:"required"`
	Age     string    `validate:"required"`
	Address []Address `validate:"required,dive"`
	Hobbies []string  `validate:"required,dive,required,min=3"`
}

func TestBasicCollection(t *testing.T) {
	employee := &Manager{
		Name: "a",
		Age:  "a",
		Address: []Address{
			{
				City:     "a",
				Province: "a",
			},
			{
				City:     "a",
				Province: "a",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"A",
		},
	}

	err := validate.Struct(employee)

	if err != nil {
		t.Error(err.Error())
	}

}
