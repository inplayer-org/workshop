package main

import "fmt"

func main() {
	dozvola := make(chan bool, 1)
	rabotnici := make(chan int, 500)
	dozvola <- false
	go func() {
		for i := 1; i <= 500; i++ {
			rabotnici <- i
			fmt.Println("Primen", i, "rabotnik")
		}
		<-dozvola
	}()
	dozvola <- true
	<-dozvola
	close(dozvola)
	close(rabotnici)
}
