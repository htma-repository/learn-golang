package learn_golang_goroutine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func AddToMap(mapData *sync.Map, keyData string, data int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	mapData.Store(keyData, data)
}

func TestMap(t *testing.T) {

	newMap := sync.Map{}
	group := sync.WaitGroup{}

	for i := 1; i <= 50; i++ {
		keyData := "data " + strconv.Itoa(i)
		go AddToMap(&newMap, keyData, i, &group)
	}

	group.Wait()
	newMap.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

	fmt.Println(newMap.Load("data 10"))

}
