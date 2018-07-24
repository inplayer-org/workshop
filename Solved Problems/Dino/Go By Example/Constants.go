package main

import "fmt"
import "math"

const zbor = "Ova e String"

func main() {
	const celBroj = 10000
	const izmenetBroj = 3e20 / celBroj

	fmt.Println(zbor)
	fmt.Println(celBroj)
	fmt.Println((celBroj + 0.5)) // Stanuva float iako e definiran kako int
	fmt.Println(math.Sin(celBroj))
	fmt.Println(izmenetBroj)
}
