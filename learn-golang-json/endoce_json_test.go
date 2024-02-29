package learn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Prices struct {
	Lite, Regular, Pro int
}

type Smartphone struct {
	Brand        string
	RAM, Storage int
	Features     []string
	Prices       []Prices
}

func JsonEncode(data any) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {

	JsonEncode("Hutama")
	JsonEncode(99)
	JsonEncode(true)
	JsonEncode([]string{"Hutama", "Trirahmanto"})
	JsonEncode(map[string]any{
		"name":    "Hutama",
		"address": "Tangerang",
	})
	JsonEncode(Smartphone{
		Brand:    "Samsung",
		RAM:      8,
		Storage:  128,
		Features: []string{"IP68", "Zoom50x", "120hzScreen"},
		Prices: []Prices{
			{
				Lite:    2000000,
				Regular: 3000000,
				Pro:     4500000,
			},
			{
				Lite:    3000000,
				Regular: 4000000,
				Pro:     5000000,
			},
		},
	})
	JsonEncode([]Smartphone{
		{
			Brand:   "Samsung",
			RAM:     8,
			Storage: 128,
		},
		{
			Brand:   "Sony",
			RAM:     8,
			Storage: 128,
		},
		{
			Brand:   "Apple",
			RAM:     8,
			Storage: 128,
		},
	})

}
