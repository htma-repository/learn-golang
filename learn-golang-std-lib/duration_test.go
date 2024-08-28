package learn_golang_std_lib

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	var durationSecond time.Duration = 10 * time.Second
	var durationMilliSecond time.Duration = 10 * time.Millisecond
	var durationNow time.Duration = durationSecond - durationMilliSecond

	fmt.Println(durationSecond)
	fmt.Println(durationMilliSecond)
	fmt.Println(durationNow)
}
