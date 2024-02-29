package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Without Cancel
func CreateCounter() chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		counter := 1
		for {
			channel <- counter
			counter++
		}
	}()
	return channel
}

// With Cancel
func CreateCounter2(ctx context.Context) chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				channel <- counter
				counter++
			}
		}
	}()
	return channel
}

func TestContextCancel(t *testing.T) {
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)

	fmt.Println("Goroutine running, ", runtime.NumGoroutine())

	// result := CreateCounter()
	result2 := CreateCounter2(ctxCancel)
	for value := range result2 {
		fmt.Println("Counter ", value)
		if value == 10 {
			break
		}
	}

	cancel()
	time.Sleep(3 * time.Second)

	fmt.Println("Goroutine running, ", runtime.NumGoroutine())

}
