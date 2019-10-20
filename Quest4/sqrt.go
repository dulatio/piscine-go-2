package sqrt

func Sqrt(x int) int {
	for i := 0; i*i <= x; i++ {
		if i*i == x {
			return i
		}
	}
	return 0
}
