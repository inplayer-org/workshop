package kata

func Perimeter(n int) int {
	// your code
	if n == 0 {
		return 0
	}
	result := 1
	previous := 1
	current := 1
	for i := 1; i < n; i++ {
		result += current
		temp := current
		current = current + previous
		previous = temp
	}
	return result * 4
}
