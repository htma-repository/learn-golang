package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

type GameData struct {
	Name         string
	releasedYear string
}

func (g *GameData) GetGameData(channel chan<- string) {
	channel <- g.Name
	channel <- g.releasedYear

	time.Sleep(2 * time.Second)
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	gameData1 := GameData{Name: "Pokemon", releasedYear: "2000"}
	gameData2 := GameData{Name: "Zelda", releasedYear: "2017"}

	go gameData1.GetGameData(channel1)
	go gameData2.GetGameData(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel-1 data ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel-2 data ", data)
			counter++
		default:
			fmt.Println("waiting data...")
		}

		if counter == 4 {
			break
		}
	}

}
