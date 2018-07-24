package main

import "fmt"
import "strconv"

func main() {
	broevi := []int{2, 3, 4}
	suma := 0
	for _, broj := range broevi {
		suma += broj
	}
	fmt.Println(suma)
	for indeks, broj := range broevi {
		if broj == 3 {
			fmt.Println("Brojot ", broj, " ima indeks ", indeks)
		}
	}

	mapa := map[string]string{"key1": "vrednost1", "key2": "vrednost2"}
	for kluch, vrednost := range mapa {
		fmt.Printf("%s -> %s \n", kluch, vrednost)
	}

	for indeks, token := range "abcd" {
		fmt.Println(indeks, token, strconv.QuoteRune(token))
	}

}
