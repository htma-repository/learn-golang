package learn_golang_json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

//go:embed smartphone.json
var smartphone []byte

func TestDecodeJson(t *testing.T) {
	// jsonData := `{"Brand": "Google", "RAM": 8, "Storage": 256}`

	smartphoneData := &Smartphone{}

	err := json.Unmarshal(smartphone, smartphoneData)

	if err != nil {
		panic(err)
	}

	fmt.Println(smartphoneData)
	fmt.Println(smartphoneData.Features)
	fmt.Println(smartphoneData.Prices[0].Pro)
}
