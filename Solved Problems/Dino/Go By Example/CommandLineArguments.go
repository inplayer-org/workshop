package main

import (
	"fmt"
	"os"
)

func recoverProgram() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("GRESHKA VO INDEKSOT !!!")
			//os.Exit(1)
		}
	}()
	funkcija()
}

func funkcija() {
	argumenti := os.Args
	fmt.Println("Argumenti so + komandna linija", argumenti)
	fmt.Println("Argumenti bez komandna linija", argumenti[1:])
	fmt.Println("Argument na pozicija 3", argumenti[2])
}

func main() {
	recoverProgram()
	fmt.Println("Print ova")
}
