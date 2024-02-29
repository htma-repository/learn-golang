package learn_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJsonEncoder(t *testing.T) {

	file, err := os.Create("product_encoder.json")

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)

	product := Product{
		Id:          2,
		Name:        "Product-2",
		Description: "Description-2",
		ImageURL:    "https://image.com/image.png",
	}

	err = encoder.Encode(product)

	if err != nil {
		panic(err)
	}

}
