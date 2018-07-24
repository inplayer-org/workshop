package main

import (
	"fmt"
)

func main() {
	rabotnici := make(chan int, 1)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-rabotnici
			if more {
				fmt.Println("Primen rabotnik", j)
			} else {
				fmt.Println("Zavrsheno so primanjeto rabotnici")
				done <- true
				return
			}
		}
	}()
	for i := 0; i < 100; i++ {
		rabotnici <- i
		fmt.Println("Praten rabotnik br", i)
	}
	close(rabotnici)
	<-done
	close(done)
}
