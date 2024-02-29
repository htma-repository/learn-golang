package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`h([a-z])a`)

	fmt.Println(regex.MatchString("hutama"))
	fmt.Println(regex.MatchString("huyama"))
	fmt.Println(regex.MatchString("huTama"))

	fmt.Println(regex.FindAllString("htama hutama huyama huTAma", 5))
}
