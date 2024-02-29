package main

import (
	"fmt"
	"time"
)

func main() {
	var durationSecond time.Duration = time.Duration(10) * time.Second
	var durationMilliSecond time.Duration = time.Duration(10) * time.Millisecond
	var durationNow time.Duration = durationSecond - durationMilliSecond

	fmt.Println(durationSecond)
	fmt.Println(durationMilliSecond)
	fmt.Println(durationNow)
}
