package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterTimeout(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return channel
}

func TestContextTimeout(t *testing.T) {
	fmt.Println("Goroutine running, ", runtime.NumGoroutine())
	ctx := context.Background()
	ctxCancel, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// result := CreateCounter()
	result2 := CreateCounterTimeout(ctxCancel)
	fmt.Println("Goroutine running, ", runtime.NumGoroutine())
	for value := range result2 {
		fmt.Println("Counter ", value)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Goroutine running, ", runtime.NumGoroutine())

}
