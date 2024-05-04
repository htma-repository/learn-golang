package learn_golang_validation

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func TestValidationVariable(t *testing.T) {
	name := "Tama"

	err := validate.Var(name, "required")
	if err != nil {
		t.Error(err.Error())
	}

	err = validate.VarCtx(context.Background(), name, "required")
	if err != nil {
		t.Error(err.Error())
	}
}
