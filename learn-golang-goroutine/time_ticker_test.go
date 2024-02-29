package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}

}

func TestTick(t *testing.T) {
	tickChan := time.Tick(1 * time.Second)

	for tick := range tickChan {
		fmt.Println(tick)
	}

}
