package main

import (
	"fmt"
	"testing"
)

type ValidationCheck struct {
	Message string
}

func (vc *ValidationCheck) Error() string {
	return vc.Message
}

type NotFoundCheck struct {
	Message string
}

func (nf *NotFoundCheck) Error() string {
	return nf.Message
}

func saveData(id string) error {
	if id == "" {
		return &ValidationCheck{Message: "Error, validation error"}
	}
	if id != "id" {
		return &NotFoundCheck{Message: "Error, not found"}
	}

	return nil
}

func TestErrCustom(t *testing.T) {
	result := saveData("")

	if result != nil {

		switch err := result.(type) {
		case *ValidationCheck:
			fmt.Println("Error validation :", err.Message)
		case *NotFoundCheck:
			fmt.Println("Error not found :", err.Message)
		default:
			fmt.Println("Error unknown :", err.Error())
		}

		// if validationErr, isError := err.(*ValidationCheck); isError {
		// 	fmt.Println("Error validation :", validationErr.Message)
		// } else if notFoundErr, isError := err.(*NotFoundCheck); isError {
		// 	fmt.Println("Error not found :", notFoundErr.Message)
		// } else {
		// 	fmt.Println("Unknown error :", err.Error())
		// }
	}
}
