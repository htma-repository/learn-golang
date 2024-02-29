package main

import (
	"fmt"
	"learn-golang-basic/helper"
)

func changeFullName(user *helper.NewUser, firstName string, lastName string) {
	user.FirstName = firstName
	user.LastName = lastName
}

func main() {
	newUser1 := helper.NewUser{}
	changeFullName(&newUser1, "Hutama", "Trirahmanto")

	fmt.Println(newUser1)

}
