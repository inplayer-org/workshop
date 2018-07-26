package main

import "fmt"

func pecati(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, "=", i)
	}
}

func main() {
	pecati("blocked")
	go pecati("j")

	go func(zbor string) {
		for j := 0; j < 10; j++ {
			fmt.Println(zbor, "=", j)
		}
	}("vtoro")
	fmt.Scanln()
}
