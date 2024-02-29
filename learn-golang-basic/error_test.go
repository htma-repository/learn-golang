package main

import (
	"errors"
	"fmt"
	"testing"
)

type UserStruct struct {
	FirstName, LastName string
	Age                 int
}

type UserData struct {
	users []UserStruct
}

func (userData *UserData) SaveUserData(data UserStruct) ([]UserStruct, error) {
	if data.FirstName == "" || data.LastName == "" || data.Age == 0 {
		return nil, errors.New("Error, user data required!")
	} else {
		userData.users = append(userData.users, data)
		return userData.users, nil
	}
}

func TestError(t *testing.T) {

	newUserData1 := UserStruct{FirstName: "Joko", LastName: "Prasetyo", Age: 30}
	newUserData2 := UserStruct{FirstName: "Hutama", LastName: "Trirahmanto", Age: 29}

	userData := UserData{users: []UserStruct{{
		FirstName: "Joko", LastName: "Prasetyo", Age: 30,
	}}}

	userData.SaveUserData(newUserData1)
	userData.SaveUserData(newUserData2)

	fmt.Println(userData.users)

}
