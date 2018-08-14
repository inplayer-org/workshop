package main

import (
	"fmt"
	"log"
)

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {
	for current := start; current < limit; current++ {
		zbir1 := 1 + current
		for i := current / 2; i > 1; i-- {
			if current%i == 0 {
				if current/i > i {
					break
				}
				if current/i == i {
					zbir1 += i
					break
				} else {
					zbir1 += i + current/i
				}
			}
		}
		if zbir1 < current {
			continue
		}
		for reverse := zbir1; reverse > current; reverse-- {
			zbir2 := 1 + reverse
			for j := reverse / 2; j > 1; j-- {
				if reverse%j == 0 {
					if reverse/j > j {
						break
					}
					if reverse/j == j {
						zbir2 += j
						break
					} else {
						zbir2 += j + reverse/j
					}
				}
			}
			//log.Println("zbir1-reverse", zbir1-reverse, "zbir2-current", current-zbir2)

			if zbir2 == zbir1 {
				log.Println("current =", current, "zbir1 =", zbir1, "reverse =", reverse, "zbir2 =", zbir2)
				//log.Println("VLEGUVA ZA current,reverse", current, reverse)
				//return []int{current, reverse}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(Buddy(5700, 6000))
}
