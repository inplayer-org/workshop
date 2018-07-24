package main

import "fmt"

func main() {
	fmt.Println("Prv for :")
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
	fmt.Println("Vtor for : ")
	for i := 1; i < 5; i *= 2 {
		fmt.Println(i)
	}
}
