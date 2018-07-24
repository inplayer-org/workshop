package main

import "fmt"

func brojBukvi() (int, string) {
	return 2, "dva"
}

func main() {

	fmt.Println(brojBukvi())
	brojki, _ := brojBukvi()
	_, bukvi := brojBukvi()
	fmt.Println("Brojot", brojki, "so tekst e", bukvi)

}
