package main

import "fmt"

func main() {
	var niza1 [10]int

	fmt.Println("empty : ", niza1)

	niza1[5] = 20
	fmt.Println("set : ", niza1)
	fmt.Println("get :", niza1[5])

	var niza2 [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			niza2[i][j] = i + j
		}
	}
	fmt.Println("2d niza = ", niza2)
}
