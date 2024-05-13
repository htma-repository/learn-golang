package learn_golang_std_lib

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {

	name := "Hutama Trirahmanto"

	fmt.Println(strings.Contains("Hutama", "tama"))
	fmt.Println(strings.Trim("        Hutama Trirahmanto    ", " "))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ReplaceAll(name, "Hutama", "Tama"))
}
