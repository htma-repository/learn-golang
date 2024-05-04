package learn_golang_generics

import (
	"fmt"
	"testing"
)

// inline type constrains
func Payment[T interface{ int | int64 | float64 }](amount, price T) T {
	total := amount * price
	return total
}

func TestPayment(t *testing.T) {
	result := Payment(10, 100000)
	fmt.Println(result)
}

// inline type parameter

type Product struct {
	Name  string
	Price int
}

func Payments[T []U, U Product](products T) U {
	var result U
	for _, product := range products {
		result = product
	}
	return result
}

func (p *Product) PricePtr() *int {
	return &p.Price
}

func TestPayments(t *testing.T) {

	products := []Product{
		{
			Name:  "Book",
			Price: 200000,
		},
		{
			Name:  "Pencil",
			Price: 50000,
		},
	}

	result := Payments(products)

	fmt.Println(result)

}
