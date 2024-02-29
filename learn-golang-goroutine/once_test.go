package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func Counter() {
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	// Only Run func once, next func will ignore to run
	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			// Cant do function with parameter
			once.Do(Counter)
		}()
	}

	group.Wait()
	fmt.Println("Count", counter)
}
