package printnbr

import "github.com/01-edu/z01"

func PrintNbr(x int) {

	w := '0'
	for j := 1; j <= x%10; j++ {
		w++
	}

	if x != 0 {
		PrintNbr(x / 10)
	}
	z01.PrintRune(w)
	return
}
