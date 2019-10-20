package recursivefactorial

func RecursiveFactorial(x int) int {
	if x < 0 {
		return 0
	}
	if x == 0 {
		return 1
	}
	return x * RecursiveFactorial(x-1)
}
