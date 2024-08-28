package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestTypeDeclaration(t *testing.T) {
	type Price float32

	const price Price = 10000
	const newPrice = 50000
	const priceNewPrice = Price(newPrice)

	fmt.Println(price)
	fmt.Println(newPrice)
	fmt.Println(priceNewPrice)
}
