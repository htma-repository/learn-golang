package main

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	games := []string{"Zelda BOTW", "Zelda TOTK", "Pokemon Sword", "Pokemon Shield"}

	for _, value := range games {
		fmt.Println(value)
	}

	users := map[string]string{
		"firstName": "Hutama",
		"lastName":  "Trirahmanto",
		"address":   "Tangerang",
		"age":       "29",
		"gender":    "male",
	}

	for index, user := range users {
		fmt.Println(index, user)
	}
}
