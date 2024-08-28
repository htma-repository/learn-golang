package learn_golang_std_lib

import (
	"fmt"
	"os"
	"testing"
)

func TestOS(t *testing.T) {

	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	name, err := os.Hostname()

	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println("Error", err.Error())
	}

}
