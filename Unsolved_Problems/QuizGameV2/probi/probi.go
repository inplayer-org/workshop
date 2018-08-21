package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	fmt.Println(IsLetter("Alex"))    // true
	fmt.Println(IsLetter("1dfew23")) // false

	a := "bla"
	switch a {
	case "bla":
		fmt.Println("bla00")
	default:
		fmt.Println("fedghejd")
	}
	fmt.Println(strconv.ParseInt("12", 0, 0))
}
