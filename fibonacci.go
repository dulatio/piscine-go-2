package fibonacci

func Fibonacci(x int) int {
	if x < 1 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}
