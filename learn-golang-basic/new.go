package main

import (
	"fmt"
	"learn-golang-basic/helper"
)

func main() {
	var user4 *helper.NewUser = new(helper.NewUser) // create new pointer like operator "&" but only return empty data
	var user5 *helper.NewUser = user4

	user5.FirstName = "Hutama"
	user5.LastName = "Trirahmanto"

	fmt.Println(user4) // firstname & lastname changes
	fmt.Println(user5) // firstname & lastname changes
}
