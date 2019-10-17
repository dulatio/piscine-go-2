package basicatoi

func BasicAtoi(str string) int {
	ok := false
	x := 0
	for _, c := range str {
		if c != '0' {
			ok = true
		}
		if ok == false {
			continue
		}
		x += int(c - '0')
		x *= 10
	}
	x /= 10
	return x
}
