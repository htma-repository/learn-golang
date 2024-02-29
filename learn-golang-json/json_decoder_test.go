package learn_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonDecoder(t *testing.T) {
	file, err := os.Open("product.json")

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)

	product := &Product{}

	err = decoder.Decode(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Name)
	fmt.Println(product.Description)

}
