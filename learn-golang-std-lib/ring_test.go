package learn_golang_std_lib

import (
	"container/ring"
	"fmt"
	"strconv"
	"testing"
)

func TestRing(t *testing.T) {
	// Circular List Data Structure
	var dataRing = ring.New(10)

	for i := 0; i <= dataRing.Len(); i++ {
		dataRing.Value = "Data Ring " + strconv.Itoa(i)
		dataRing = dataRing.Next()
	}

	dataRing.Do(func(data any) {
		fmt.Println(data)
	})

}
