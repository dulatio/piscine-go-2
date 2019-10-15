package printnbr

import "github.com/01-edu/z01"

func GIVE(y int) {
	if y == 0 {
		return
	}
	w := '0'
	for j := 1; j <= y%10; j++ {
		w++
	}
	GIVE(y / 10)
	z01.PrintRune(w)
	return
}
func PrintNbr(x int) {
	if x == 0 {
		z01.PrintRune('0')
	}

	if x < 0 {
		z01.PrintRune('-')
		x *= -1
	}
	GIVE(x)
}
