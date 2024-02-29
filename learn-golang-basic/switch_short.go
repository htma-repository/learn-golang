package main

import "fmt"

func main() {
	switch name := "Hutama"; name {
	case "Hutama":
		fmt.Println("Hello " + name)
	case "Trirahmanto":
		fmt.Println("Hello " + name)
	}
}
