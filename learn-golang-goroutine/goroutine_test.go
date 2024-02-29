package learn_golang_goroutine

import (
	"testing"
	"time"
)

func TestGoroutineMul(t *testing.T) {

	channel := make(chan int)
	defer close(channel)
	go MultiplyNum(10, 10, channel)

	// result := <-channel
	// fmt.Println("result", result)
	go GetDataChannel(channel)

	time.Sleep(1 * time.Second)
}

func TestGoroutineLoop(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
