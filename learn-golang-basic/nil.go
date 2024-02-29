package main

import (
	"fmt"
)

func getUserName(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	data := getUserName("")

	if data == nil {
		fmt.Println("Data nil")
	} else {
		fmt.Println("Data available", data)
	}
}
