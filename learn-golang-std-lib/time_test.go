package learn_golang_std_lib

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())
	fmt.Println(now.Unix())

	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(date)
	fmt.Println(date.Local())

	layout := "2006-01-02 15:04:05"

	timeParsed, err := time.Parse(layout, "1994-12-19 00:00:00")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(timeParsed)
	}

	fmt.Println(timeParsed.Year())
	fmt.Println(timeParsed.Month())
	fmt.Println(timeParsed.Day())
	fmt.Println(timeParsed.Hour())
	fmt.Println(timeParsed.Second())

}
