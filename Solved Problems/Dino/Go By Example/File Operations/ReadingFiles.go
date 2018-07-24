package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, e := ioutil.ReadFile("Variables.go")
	if e != nil {
		panic(e)
	}
	fmt.Print(string(f))
}
