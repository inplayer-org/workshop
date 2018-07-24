package main

import "fmt"

func zbirNaNiza(niza ...int) int {
	fmt.Println("Printanje na niza", niza)
	zbir := 0
	for _, vrednost := range niza {
		zbir += vrednost
	}
	return zbir
}

func main() {
	fmt.Println("Zbirot e 3 + 4 + 5 =", zbirNaNiza(3, 4, 5))
	fmt.Println("Zbirot e 2 + 10 + 12 =", zbirNaNiza(2, 10, 12))
	a, b, c, d, e := 1, 2, 3, 4, 5
	fmt.Printf("Zbirot e %d + %d + %d + %d + %d = %d\n", a, b, c, d, e, zbirNaNiza(a, b, c, d, e))
	fmt.Println
}
