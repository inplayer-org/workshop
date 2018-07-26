package main

import "fmt"

func wrapper() func() int { // klasa sto return argument ima (func() int)
	brojac := 0
	return func() int { // klasa bez ime za povikuvanje
		brojac++
		return brojac
	}
}

func main() {
	a := wrapper()
	// so edna zagrada a ja prima celata funkcija kako argument (prviot return od wrapper funkcijata)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(wrapper()())
	// so dve zagradi b ja prima vrednosta od returnot vo vtorata funkcija (vo ovoj slucaj int vrednost)
	b := wrapper()()
	fmt.Println(b)
	fmt.Println(b)
}
