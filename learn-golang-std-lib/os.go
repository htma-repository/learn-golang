package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	fmt.Println(args)

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
