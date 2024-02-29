package learn_golang_goroutine

import (
	"fmt"
)

// chan<-  > for sending data to channel
func MultiplyNum(num1, num2 int, channel chan<- int) {
	result := num1 * num1
	channel <- result
}

// <-chan  > for receive data from channel
func GetDataChannel(channel <-chan int) {
	result := <-channel
	fmt.Println("result", result)
}

func DisplayNumber(num int) {
	fmt.Println("Number", num)
}

func main() {
	DisplayNumber(19)
}
