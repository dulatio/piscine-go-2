package basicatoi

func BasicAtoi(str string) int {
	ok := false
	x := 0
	ln := 0
	for _, c := range str {
		if c == c {
			ln++
		}
	}
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
		if c != rune(str[ln-1]) {
			x *= 10
		}
	}
	return x
}
