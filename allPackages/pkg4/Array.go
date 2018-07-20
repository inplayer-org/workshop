package soberiniza

func SoberiNizaa(a, b int) [][]int {
	twoD := [][]int{}
	for i := 0; i < a; i++ {
		temp := []int{}
		for j := 0; j < b; j++ {
			temp = append(temp, i+j)
		}
		twoD = append(twoD, temp)
	}
	return twoD
}
