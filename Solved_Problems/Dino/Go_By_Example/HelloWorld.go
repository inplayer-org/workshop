package main

import (
	"fmt"
)

type imeNa struct {
	godini int
}

func smeni(o *imeNa) {
	o.godini = 10
}

func main() {
	prom := &imeNa{2}
	fmt.Println(prom.godini)
	smeni(prom)
	fmt.Println(prom.godini)
}
