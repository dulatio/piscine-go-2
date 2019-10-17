package basicatoi

func BasicAtoi(str string) int {
	ok := false
	x := 0
	for _, c := range str {
		if c != 48 {
			ok = true
		}
		if ok == false {
			continue
		}
		cnt := 0
		for i := '1'; i <= c; i++ {
			cnt++
		}
		x += cnt
		x *= 10
	}
	x /= 10
	return x
}
