package main

import "fmt"

func main() {
	prk1, prk2, prk3 := "poraka1", "poraka2", "poraka3"
	kanal := make(chan string, 3)
	kanal <- prk1
	kanal <- prk2
	kanal <- prk3
	close(kanal)
	for msg := range kanal {
		fmt.Println("msg =", msg)

	}
}
