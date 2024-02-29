package learn_golang_goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)

	thread := runtime.GOMAXPROCS(-1)
	fmt.Println(thread)

	goRoutine := runtime.NumGoroutine()
	fmt.Println(goRoutine)

}
