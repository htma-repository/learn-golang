package main

import (
	"fmt"
	"learn-golang-basic/helper"
)

func main() {
	// default golang is pass by value
	// user1 := NewUser{FirstName: "Hutama", LastName: "Trirahmanto", Age: 29}
	// user2 := user1 // copy from user1 (pass by value)

	// user2.Age = 25

	// fmt.Println(user1) // age dont change
	// fmt.Println(user2) // age change to 25

	// change default from pass by value to pass by references
	var user1 helper.NewUser = helper.NewUser{FirstName: "Hutama", LastName: "Trirahmanto", Age: 29}
	var user2 *helper.NewUser = &user1 // references from user 1 (pass by references)

	user2.Age = 25

	fmt.Println(user1) // age change to 25
	fmt.Println(user2) // age change to 25
}
