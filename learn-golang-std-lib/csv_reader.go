package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvData := "name,age\nJohn,30\nJane,25\nBob,40"

	reader := csv.NewReader(strings.NewReader(csvData))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else {
			fmt.Println(record)
		}
	}
}
