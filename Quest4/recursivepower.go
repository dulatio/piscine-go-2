package recursivepower

func RecursivePower(x int, y int) int {
	if y < 0 {
		return 0
	}
	if y == 0 {
		return 1
	}
	return x * RecursivePower(x, y-1)
}
