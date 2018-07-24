package main

import "fmt"

func main() {
	s1 := make([]string, 3)
	fmt.Println("Empty : ", s1)
	s1[1] = "b"
	fmt.Println("Get na pozicija 1 :", s1[1])
	s1 = append(s1, "w")
	fmt.Println("Append : ", s1)
	fmt.Println("Length : ", len(s1))
	s1 = append(s1, "er")
	fmt.Println("Append : ", s1)
	fmt.Println("Length : ", len(s1))
	fmt.Println("Slice [3:5]", s1[3:5])

	s2 := make([]string, len(s1))
	copy(s2, s1)
	fmt.Println("Copy :", s2)

}
