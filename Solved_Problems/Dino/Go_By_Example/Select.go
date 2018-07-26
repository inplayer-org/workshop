package main

import "fmt"
import "time"

func main() {
	prvKanal := make(chan int)
	vtorKanal := make(chan int)
	prekini := make(chan bool)

	go func() {
		for true {
			time.Sleep(time.Second * 2)
			prvKanal <- 1
		}
	}()
	go func() {
		for true {
			time.Sleep(time.Second * 3)
			vtorKanal <- 2
		}
	}()
	go func() {
		time.Sleep(time.Second * 10)
		prekini <- true
	}()
	for i := false; !i; {
		select {
		case msg1 := <-prvKanal:
			fmt.Println(msg1)
		case msg2 := <-vtorKanal:
			fmt.Println(msg2)
		case i = <-prekini:
			break
		}
	}
}
