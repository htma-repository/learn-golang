package helper

import "fmt"

func TypeAssert() any {
	return false
}

type NewUser struct {
	FirstName, LastName string
	Age                 int
}

type Filter func(string) string

type Human struct {
	Name    string
	Address string
}

// For method highly recommended always use pointer for minimize memory consumption
func (human *Human) SayHello() {
	human.Name = "Hello, " + human.Name
	human.Address = "Your Address is " + human.Address
}

type Users interface {
	GetUserDetail() (string, string, string)
}

type Customer struct {
	FirstName string
	LastName  string
	Address   string
}

func (customer Customer) GetUserDetail() (string, string, string) {
	return customer.FirstName, customer.LastName, customer.Address
}

type User struct {
	FirstName, LastName, City, Address, PhoneNumber string
	Age                                             int
}

func (user User) GetUserData() {
	fmt.Println(user.FirstName, user.LastName, user.Age, "Address", user.Address, "City", user.City, "PhoneNumber", user.PhoneNumber)
}
