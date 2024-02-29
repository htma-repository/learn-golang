package main

import "fmt"

func main() {
	var allName = [...]string{"Hutama", "Trirahmanto", "Ajim", "Bayu", "Gerdy"}
	friendName := allName[2:]
	myName := allName[:2]
	allNameSlice := allName[:]

	// change slice data, array data allName changed too
	friendName[0] = "Miko"

	fmt.Println(allName)
	fmt.Println(friendName)
	fmt.Println(myName)
	fmt.Println(allNameSlice)
	fmt.Println(len(friendName))
	fmt.Println(cap(friendName))

	newFriendName := append(friendName, "Nicho")

	fmt.Println(newFriendName)
	fmt.Println(allName)

	newSlice := make([]string, 3, 5)
	newSlice[0] = "Hutama"
	newSlice[1] = "Trirahmanto"
	newSlice[2] = "Uut"
	// newSlice[3] = "Tri" // error, must use append because size of slice is 3 & cap is 5
	newSlice2 := append(newSlice, "Tri")

	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Tama"

	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	fromSlice := allName[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
