package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Double linked list Data Structure
	var listData *list.List = list.New()

	listData.PushBack("Hutama")
	listData.PushBack("Trirahmanto")
	listData.PushBack("Tama")
	listData.PushBack("Tri")

	for data := listData.Front(); data != nil; data = data.Next() {
		fmt.Println(data.Value)
	}
}
