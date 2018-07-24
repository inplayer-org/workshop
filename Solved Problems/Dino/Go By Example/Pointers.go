package main

import "fmt"

func zeroValue(broj int) {
	broj = 0
}

func zeroPointer(broj *int) {
	*broj = 0
}

func stringValue(zbor string) {
	zbor = "qwerty"
}

func stringPointer(zbor *string) {
	*zbor = "qwerty"
}

func main() {
	a := 5

	zeroValue(a)
	fmt.Println(a)
	zeroPointer(&a)
	fmt.Println(a)

	b := "abcd"

	stringValue(b)
	fmt.Println(b)
	stringPointer(&b)
	fmt.Println(b)

}
