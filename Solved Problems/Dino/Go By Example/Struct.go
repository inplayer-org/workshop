package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	covek1 := person{"Ana", 20}
	fmt.Println(covek1)
	covek2 := person{"Marw", 25}
	fmt.Println(covek2)
	covekPointer := &covek2
	covekPointer.name = "Mark"
	fmt.Println(covek2, covekPointer)

}
