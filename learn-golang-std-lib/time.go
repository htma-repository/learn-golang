package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())
	fmt.Println(now.Unix())

	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(date)
	fmt.Println(date.Local())

	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, "1994-12-19 00:00:00")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(t)
	}

	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Second())

}
