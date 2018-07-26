package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Imash 5 sekundi za vnes")
	done := make(chan bool)
	go func() {
		var vnes int
		fmt.Scanln(&vnes)
		fmt.Println("Vnesovte", vnes)
		done <- true
	}()
	time.Sleep(time.Second)
	ticking := time.Tick(time.Second)
	go func() {
		for i := 5; i > 0; i-- {
			fmt.Println("Ostanuvaat uste", i, "sekundi !!")
			<-ticking
		}
		done <- false
	}()
	tochnost := <-done
	if tochnost {
		fmt.Println("Korisnikot vnese broj")
	} else {
		fmt.Println("Korisnikot NE vnese broj")
	}

}
