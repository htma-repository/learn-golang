package main

import (
	"fmt"
	"learn-golang-basic/helper"
	"testing"
)

func TestAsteriskOperator(t *testing.T) {
	var user1 helper.NewUser = helper.NewUser{FirstName: "Hutama", LastName: "Trirahmanto", Age: 29}
	var user2 *helper.NewUser = &user1 // references from user 1 (pass by references)
	var user3 *helper.NewUser = &user1

	user2.Age = 25

	fmt.Println(user1) // age change to 25
	fmt.Println(user2) // age change to 25
	fmt.Println(user3) // age change to 25

	// All Pointer change to this data include data used with NewUser struct
	*user2 = helper.NewUser{FirstName: "Joko", LastName: "Prasetyo", Age: 30}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

	// Double pointer
	var user4 **helper.NewUser = &user2

	fmt.Println(user4)
}
