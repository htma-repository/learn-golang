package learn_golang_std_lib

import (
	"fmt"
	"slices"
	"testing"
)

func TestSlices(t *testing.T) {
	dataUser := []string{"Hutama", "Trirahmanto", "Joko"}
	dataNumber := []int{10, 27, 50, 50}

	fmt.Println(slices.Min(dataNumber))
	fmt.Println(slices.Max(dataUser))
	fmt.Println(slices.Contains(dataUser, "Hutama"))
	fmt.Println(slices.Index(dataUser, "Hutama"))
	fmt.Println(slices.Index(dataUser, "Budi"))
}
