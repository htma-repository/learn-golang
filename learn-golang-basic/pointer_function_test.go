package learn_golang_basic

import (
	"fmt"
	"learn-golang-basic/helper"
	"testing"
)

func changeFullName(user *helper.NewUser, firstName string, lastName string) {
	user.FirstName = firstName
	user.LastName = lastName
}

func TestPointerFunc(t *testing.T) {
	newUser1 := helper.NewUser{}
	changeFullName(&newUser1, "Hutama", "Trirahmanto")

	fmt.Println(newUser1)

}
