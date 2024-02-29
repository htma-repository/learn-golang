package learn_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
)

func addToChannel(channel chan<- string) {
	for i := 0; i < 10; i++ {
		channel <- "Channel " + strconv.Itoa(i)
	}
	// MUST use close channel
	close(channel)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go addToChannel(channel)

	for data := range channel {
		fmt.Println(data)
	}

}
