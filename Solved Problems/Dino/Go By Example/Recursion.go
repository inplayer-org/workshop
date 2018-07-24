package main

import (
	"fmt"
)

// Fibonachi niza od n broja so rekurzija

func fibonachiDescending(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return (fibonachiDescending(n-2) + fibonachiDescending(n-1))
}

/*func fibonachiAscending(n int, flag bool) func() int {
	i := 1
	pomal, pogolem := 1, 1
	if flag {
		return func() int {
			return 0
		}
	}
	vrati := fibonachiAscending(n, true)
	return func() int {
		if i == n {
			return (pomal + pogolem)
		}
		temp := pogolem
		pogolem = pogolem + pomal
		pomal = temp
		i++
		return vrati()
	}
} */

func main() {
	fmt.Println(fibonachiDescending(5))
	//fmt.Println(fibonachiAscending(10, false)())
}
