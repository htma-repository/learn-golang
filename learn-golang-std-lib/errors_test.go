package learn_golang_std_lib

import (
	"errors"
	"fmt"
	"testing"
)

var (
	validationError = errors.New("Validation error")
	notFoundError   = errors.New("Not found error")
)

func validationUser(userId string) error {
	if userId == "" {
		return validationError
	}
	if userId != "id" {
		return notFoundError
	}

	//success
	return nil
}

func TestErrors(t *testing.T) {

	err := validationUser("")

	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Error occurred :", err.Error())
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Error occurred :", err.Error())
		} else {
			fmt.Println("Unknown error :", err.Error())
		}
	} else {
		fmt.Println("Success")
	}

	fmt.Println("End")

}
