package main

import "fmt"

func main() {
	if n := 9; n > 9 {
		fmt.Println("n > 9")
	} else { // else mora da bide vo ista linija so "}" zagradata !!!
		fmt.Println("n <= 9")
	}
}
