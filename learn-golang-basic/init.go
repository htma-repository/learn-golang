package main

import (
	"fmt"
	"learn-golang-basic/database"
	_ "learn-golang-basic/internal"
)

func main() {
	fmt.Println(database.GetDbConnect())
}
