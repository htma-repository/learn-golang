package learn_golang_std_lib

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	text := "Hutama Trirahmanto"

	encoded := base64.StdEncoding.EncodeToString([]byte(text))

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}

}
