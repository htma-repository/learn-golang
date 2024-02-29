package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterDeadline(ctx context.Context) chan int {
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

func TestContextDeadline(t *testing.T) {
	fmt.Println("Goroutine running, ", runtime.NumGoroutine())
	ctx := context.Background()
	ctxCancel, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	// result := CreateCounter()
	result2 := CreateCounterDeadline(ctxCancel)
	fmt.Println("Goroutine running, ", runtime.NumGoroutine())
	for value := range result2 {
		fmt.Println("Counter ", value)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Goroutine running, ", runtime.NumGoroutine())

}
