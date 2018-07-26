package main

import "fmt"

func main() {
	mapa := make(map[string]int)
	mapa["k1"] = 1
	mapa["k2"] = 2
	fmt.Println(mapa)

	v1, vrakja := mapa["k2"]
	fmt.Println("value = ", v1, " - a kluchot ", vrakja)

	delete(mapa, "k2")

	fmt.Println("Delete key : [k2]")
	fmt.Println(mapa)
	v1, vrakja = mapa["k2"]
	fmt.Println("value = ", v1, " - a kluchot ", vrakja)

}
