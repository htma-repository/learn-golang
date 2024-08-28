package learn_golang_basic

import (
	"fmt"
	"learn-golang-basic/database"
	_ "learn-golang-basic/internal"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println(database.GetDbConnect())
}
