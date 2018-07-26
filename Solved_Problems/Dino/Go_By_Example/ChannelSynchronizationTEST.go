package main

import (
	"fmt"
)

func raboti(rabotnik chan int, done chan<- bool) {
	fmt.Println("Vleguva")
	for covek := range rabotnik {
		fmt.Printf("Worker n.%d \t", covek)
		//time.Sleep(time.Second)

	}
	done <- true

}

func main() {

	rabotnik := make(chan int, 100)
	done := make(chan bool, 1)
	go raboti(rabotnik, done)
	for i := 1; i <= 100; i++ {
		//dozvola <- true
		rabotnik <- i
	}
	//dozvola <- true
	close(rabotnik)

	<-done
}
