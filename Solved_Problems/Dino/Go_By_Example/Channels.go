package main

import (
	"fmt"
)

func main() {
	poraka1 := "Poraka so channel"
	poraka2 := "Poraka bez channel"
	kanal := make(chan string)
	var neKanal string
	go func() {
		kanal <- poraka1
	}()

	go func() {
		neKanal = poraka2
	}()

	fmt.Println(neKanal)
	fmt.Println("--------------------")
	fmt.Println(<-kanal)
	fmt.Println("++++++++++++++++++++")
}
