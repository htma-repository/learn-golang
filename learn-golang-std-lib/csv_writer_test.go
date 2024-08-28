package learn_golang_std_lib

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCSVWriter(t *testing.T) {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"name", "age"})
	_ = writer.Write([]string{"John", "30"})
	_ = writer.Write([]string{"Jane", "25"})

	writer.Flush()
}
