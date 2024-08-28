package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestSwitchShort(t *testing.T) {
	switch name := "Hutama"; name {
	case "Hutama":
		fmt.Println("Hello " + name)
	case "Trirahmanto":
		fmt.Println("Hello " + name)
	}
}
