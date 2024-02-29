package learn_golang_json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {

	product := Product{
		Id:          1,
		Name:        "Product-1",
		Description: "Description-1",
		ImageURL:    "https://image.com/image.png",
	}

	b, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//go:embed product.json
var product []byte

func TestJsonTagDecode(t *testing.T) {

	productData := &Product{}

	err := json.Unmarshal(product, productData)

	if err != nil {
		panic(err)
	}

	fmt.Println(productData)
	fmt.Println(productData.Name)
}
