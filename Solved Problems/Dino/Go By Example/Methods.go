package main

import "fmt"

type pravoagolnik struct {
	visina, shirina int
}

type procitaj interface {
	perimetar() int
	ploshtina() int
}

func (p *pravoagolnik) perimetar() int {
	fmt.Println("Se povikuva funkcijata za perimetar pokazhuvach *p")
	return p.visina*2 + p.shirina*2
}

func (p pravoagolnik) ploshtina() int {
	fmt.Println("Se povikuva funkcijata za ploshtina value *p")
	return p.visina * p.shirina
}

func izmeriPloshtina(g procitaj) int {
	return g.ploshtina()
}

func izmeriPerimetar(g procitaj) int {
	return g.perimetar()
}

func main() {
	pravo := pravoagolnik{5, 10}
	fmt.Println(pravo)
	pokazh := &pravo
	fmt.Println(izmeriPerimetar(&pravo))
	fmt.Println(izmeriPerimetar(pokazh))
	fmt.Println(izmeriPloshtina(&pravo))
	fmt.Println(izmeriPloshtina(pokazh))
}
