package importall

import "./pkg1"
import "./pkg2"
import "./pkg3"
import "./pkg4"
import "./pkg5"

func Soberi(a, b int) int {
	return soberiDvaBroja.SoberiDvaBroja(a, b)
}

func Mnozhi(a, b int) int {
	return mnozi.MnoziDvaBroja(a, b)
}

func Deli(a, b int) int {
	return dzalepackage.DeliDvaBroja(a, b)
}
func SoberiNizaa(a, b int) [][]int {
	return soberiniza.SoberiNizaa(a, b)
}
func SoberiFloati(a, b float64) float64 {
	return soberiDvaFloati.SoberiDvaFloati(a, b)
}
