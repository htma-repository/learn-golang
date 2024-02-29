package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunningCond(cond *sync.Cond, group *sync.WaitGroup, value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("value", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	group := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		go RunningCond(cond, &group, i)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()

}
