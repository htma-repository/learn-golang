package main

import "fmt"

func main() {
	people := make(map[string]string)
	people["name"] = "Hutama Trirahmanto"
	people["address"] = "Tangerang"

	fmt.Println(people)

	car := map[string]string{
		"type":  "sedan",
		"color": "red",
		"seat":  "4",
	}

	fmt.Println(car)
	fmt.Println(car["type"])
	fmt.Println(len(car))

	delete(car, "seat")

	fmt.Println(car)
}
