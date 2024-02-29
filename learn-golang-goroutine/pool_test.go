package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{New: func() any { return 0 }}
	group := sync.WaitGroup{}
	// channel := make(chan int)

	pool.Put(10)
	pool.Put(20)
	pool.Put(30)

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get()

			fmt.Println(data)
			time.Sleep(1 * time.Second)

			pool.Put(data)

		}()
	}

	group.Wait()
	fmt.Println("Success")
}
