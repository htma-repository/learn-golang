package learn_golang_validation

import "testing"

type School struct {
	Name string `validate:"required"`
}

type Child struct {
	Name    string            `validate:"required"`
	Age     string            `validate:"required"`
	Address []Address         `validate:"required,dive"`
	Hobbies []string          `validate:"required,dive,required,min=3"`
	Schools map[string]School `validate:"required,dive,keys,required,min=2,endkeys,required"`
	Wallet  map[string]int    `validate:"required,dive,keys,required,min=3,endkeys,required,gt=100"`
}

func TestValidationMap(t *testing.T) {
	request := &Child{
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
		},
		Schools: map[string]School{
			"S": {
				Name: "a",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA":     100000,
			"Mandiri": 50,
			"":        0,
		},
	}

	err := validate.Struct(request)

	if err != nil {
		t.Error(err.Error())
	}

}
