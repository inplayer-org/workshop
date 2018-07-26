package main

import (
	"fmt"
)

type krug struct {
	ime    string
	tip    string
	radius int
}

type kocka struct {
	ime    string
	tip    string
	strana int
}

type geom interface {
	procitaj()
}

func (kr krug) procitaj() {
	fmt.Println("Ime :", kr.ime, "\tTip :", kr.tip, "\tDolzhina na radius :", kr.radius)
}

func (ko kocka) procitaj() {
	fmt.Println("Ime :", ko.ime, "\tTip :", ko.tip, "\tDolzhina na strana :", ko.strana)
}

func preraboti(elem geom) {
	elem.procitaj()
}

func main() {
	krug1 := krug{"Kr1", "Krug", 10}
	kocka1 := kocka{"Ko1", "Kocka", 5}
	preraboti(krug1)
	preraboti(kocka1)
}
