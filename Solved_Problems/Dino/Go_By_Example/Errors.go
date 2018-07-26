package main

import "fmt"

import "reflect"

type customError struct {
	promenliva interface{}
	poraka     string
}

func (msg customError) Error() {
	fmt.Println("Promenlivata od tip", msg.promenliva, "ima greshka poradi", msg.poraka)
}

func proveriDaliEBroj(tip interface{}) {
	if tip == reflect.TypeOf(1) {
		fmt.Println("Printanje tip :", tip)
		fmt.Printf("Tipot na podatokot e int\n")
	} else {
		i := customError{tip, "greshen tip na podatok"}
		i.Error()
	}
}

func main() {
	broj := 10
	forma := "Krug"
	if reflect.TypeOf(broj).String() == "int" {
		fmt.Println(reflect.TypeOf(broj), reflect.TypeOf(forma))
	}
	proveriDaliEBroj(reflect.TypeOf(broj))
	proveriDaliEBroj(reflect.TypeOf(forma))

}
