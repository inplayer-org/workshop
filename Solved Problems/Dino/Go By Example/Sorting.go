package main

import "fmt"
import "sort"

func main() {
	stringovi := []string{"abc", "abbb", "ad"}
	sort.Strings(stringovi)

	broevi := []int{3, 412, 32, 162, 421, 9, 0, 2, 0, 12}
	sort.Ints(broevi)

	fmt.Println(stringovi)
	fmt.Println(broevi)
	fmt.Println(sort.IntsAreSorted(broevi))

}
