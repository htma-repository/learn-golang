package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func SayHello(name string) {
	fmt.Println("Hello ", name)
}
func TestTimer(t *testing.T) {

	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())

	// delay the execution
	newTimer := <-timer.C
	SayHello("Hutama")
	fmt.Println(newTimer)

}

func TestAfter(t *testing.T) {
	// More simpler way than NewTimer
	chanTimer := time.After(3 * time.Second)
	fmt.Println(time.Now())

	// delay the execution
	newTimer := <-chanTimer
	SayHello("Hutama")
	fmt.Println(newTimer)

}

func TestAfterFunc(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(1)

	// More simpler way than After
	time.AfterFunc(3*time.Second, func() {
		defer group.Done()
		fmt.Println("delay time", time.Now())
		SayHello("Hutama")
	})

	fmt.Println("current time", time.Now())
	group.Wait()
}
