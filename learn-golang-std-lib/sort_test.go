package learn_golang_std_lib

import (
	"fmt"
	"sort"
	"testing"
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

func TestSort(t *testing.T) {

	user := []User{
		{"Hutama", 29},
		{"Tama", 25},
		{"Tri", 27},
		{"Rahmanto", 26},
	}

	sort.Sort(UserSlice(user))

	fmt.Println(user)

}
