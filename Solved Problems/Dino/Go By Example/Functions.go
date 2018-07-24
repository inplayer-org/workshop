package main

import "fmt"

func soberiDvaBroja(a int, b int) int {
	return a + b
}

func odrediMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prv, vtor := 10, 5
	fmt.Printf("Zbirot na %d i %d e %d\n", prv, vtor, soberiDvaBroja(prv, vtor))
	fmt.Println("Pogolemiot broj od tie dva broja e :", odrediMax(prv, vtor))

}
