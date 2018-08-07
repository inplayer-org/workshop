package main

import (
	"fmt"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(s)
	for _, j := range s {
		fmt.Println(j)
	}
}
