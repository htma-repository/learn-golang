package main

import (
	"learn-golang-basic/helper"
)

func main() {
	var hutama helper.User
	hutama.FirstName = "Hutama"
	hutama.LastName = "Trirahmanto"
	hutama.Address = "Jl.elok 1"
	hutama.City = "Tangerang"
	hutama.PhoneNumber = "081213839577"
	hutama.Age = 29

	hutama.GetUserData()

	albert := helper.User{
		FirstName:   "Albert",
		LastName:    "Einstein",
		City:        "New York",
		Address:     "New York Street",
		PhoneNumber: "09123812384",
		Age:         29,
	}

	graham := helper.User{
		FirstName:   "Graham",
		LastName:    "Coxon",
		City:        "London",
		Address:     "London Street",
		PhoneNumber: "09123812384",
		Age:         50,
	}

	albert.GetUserData()
	graham.GetUserData()

}
