package main

import "fmt"

func main() {

	bookPrice := 70000

	switch bookPrice {
	case 50000:
		fmt.Println("Book price is 50,000")
	case 100000:
		fmt.Println("Book price is 100,000")
	default:
		fmt.Println("Book price is unknown")
	}

	switch {
	case bookPrice < 50000:
		fmt.Println("Book is cheap")
	case bookPrice >= 50000 && bookPrice <= 100000:
		fmt.Println("Book is affordable")
	default:
		fmt.Println("Book is very expensive")
	}
}
