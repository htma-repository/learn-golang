package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup) {
	// MUST use Done to not deadlock
	defer group.Done()

	group.Add(1)

	fmt.Println("WaitGroup Run")
	time.Sleep(3 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("Success")
}
