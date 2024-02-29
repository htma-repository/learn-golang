package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"name", "age"})
	_ = writer.Write([]string{"John", "30"})
	_ = writer.Write([]string{"Jane", "25"})

	writer.Flush()
}
