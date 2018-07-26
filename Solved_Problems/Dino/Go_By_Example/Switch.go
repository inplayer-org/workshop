package main

import "fmt"

func main() {
	OdrediTip := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("Neregistriran tip na promenliva (%T) \n", t)
		}
	}
	OdrediTip(true)
	OdrediTip(15.26)
	OdrediTip("zbor")
	OdrediTip(15)
	OdrediTip(nil)
}
