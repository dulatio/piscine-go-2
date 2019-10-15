package printnbr

import "github.com/01-edu/z01"

func PrintNbr(x int) {
	for {
		if x == 0 {
			break
		}
		i := '0' + x%10
		z01.PrintRune(i)
		x /= 10

	}
	z01.PrintRune('\n')
}
