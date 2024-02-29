package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// these method func follow sort interface contract

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {

	user := []User{
		{"Hutama", 29},
		{"Tama", 25},
		{"Rahmanto", 26},
	}

	sort.Sort(UserSlice(user))

	fmt.Println(user)

}
