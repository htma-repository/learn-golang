package main

import "fmt"

func main() {
	type Price float32

	const price Price = 10000
	const newPrice = 50000
	const priceNewPrice = Price(newPrice)

	fmt.Println(price)
	fmt.Println(newPrice)
	fmt.Println(priceNewPrice)
}
