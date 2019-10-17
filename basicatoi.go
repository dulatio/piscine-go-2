package basicatoi

func BasicAtoi(str string) int {
	ok := false
	var x int
	for _, c := range str {
		if c != 48 {
			ok = true
		}
		if ok == false {
			continue
		}
		var cnt int
		for i := '1'; i <= c; i++ {
			cnt++
		}
		x += cnt
		x *= 10
	}
	x /= 10
	return x
}
