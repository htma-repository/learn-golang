package main

import (
	"fmt"
)

type User interface {
	GetUserDetail() (string, string, int)
}

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) GetUserDetail() (string, string, int) {
	return customer.Name, customer.Address, customer.Age
}

func getUserDetail(user User) {
	Name, Address, Age := user.GetUserDetail()
	fmt.Println("Name", Name, "Address", Address, "Age", Age)
}

func main() {

	customerOne := Customer{
		Name:    "Hutama",
		Address: "Tangerang",
		Age:     29,
	}

	getUserDetail(customerOne)

	// Interface{} / any

	var customerTwo any = "Hutama"
	customerTwo = 1

	fmt.Println(customerTwo)

}
