package learn_golang_json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJsonEncode(t *testing.T) {
	product := map[string]any{
		"id":          1,
		"name":        "Product-1",
		"description": "description-1",
		"price":       100000,
	}

	b, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//go:embed product.json
var productMap []byte

func TestMapJsonDecode(t *testing.T) {

	var productData map[string]any

	err := json.Unmarshal(productMap, &productData)

	if err != nil {
		panic(err)
	}

	fmt.Println(productData)
	fmt.Println(productData["name"])
	fmt.Println(productData["description"])
	fmt.Println(productData["image_url"])
}
