package printnbr

import "github.com/01-edu/z01"

func GIVE(y int) {
	w := '0'
	if y == 0 {
		z01.PrintRune(w)
		return
	}
	for j := 1; j <= y%10; j++ {
		w++
	}
	for j := -1; j >= y%10; j-- {
		w++
	}
	if y/10 != 0 {
		GIVE(y / 10)
	}
	z01.PrintRune(w)
	return
}

func PrintNbr(x int) {
	if x < 0 {
		z01.PrintRune('-')
	}
	GIVE(x)
}
