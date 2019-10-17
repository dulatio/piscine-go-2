package atoi

func check(s string) bool {

	for _, c := range s {
		if c == ' ' || c < '0' || c > '9' {
			return false
		}
	}
	return true
}
func Atoi(str string) int {
	x := 0
	var this string
	ok := true
	change := false
	for _, c := range str {
		if c != '+' && c != '-' {
			ok = false
		}
		if ok == true {
			if c == '-' {
				change = true
			}
			if c == '+' {
				change = false
			}
		} else {
			this = this + string(c)
		}
	}
	if check(this) == true {
		for _, c := range this {

			cnt := 0
			for i := '1'; i <= c; i++ {
				cnt++
			}
			x = x*10 + cnt
		}
	}
	if change == true {
		x *= -1
	}
	return x
}
