package iterativepower

func IterativePower(x int, y int) int {
	ans := 1
	if y < 0 {
		return 0
	}
	for i := 1; i <= y; i++ {
		ans *= x
	}
	return x
}
