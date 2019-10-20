package iterativefactorial

func IterativeFactorial(x int) int {
	if x < 0 {
		return 0
	}
	ans := 1
	for i := 2; i <= x; i++ {
		ans *= i
		if ans < 0 {
			return 0
		}
	}
	return ans
}
