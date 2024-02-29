package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

type userData struct {
	FirstName, LastName, Address string
}

func (u *userData) GetUserData(channel chan<- string) {
	channel <- u.FirstName
	channel <- u.LastName
	channel <- u.Address
}

func TestBufferChannel(t *testing.T) {
	userChannel := make(chan string, 3)
	defer close(userChannel)
	newUser := userData{FirstName: "Hutama", LastName: "Trirahmanto", Address: "Cipondoh"}

	go newUser.GetUserData(userChannel)

	fmt.Println(cap(userChannel))
	fmt.Println(len(userChannel))
	fmt.Println(<-userChannel)
	fmt.Println(<-userChannel)
	fmt.Println(<-userChannel)

	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Hutama"
		channel <- "Trirahmanto"
		channel <- "Cipondoh"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

}
