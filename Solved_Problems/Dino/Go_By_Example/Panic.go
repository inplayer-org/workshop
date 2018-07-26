package main

import (
	"os"
)

func main() {
	//panic("Panika golema")

	_, err := os.Open("Zdravo")
	if err != nil {
		panic(err)
	}
}
