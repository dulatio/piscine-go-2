package printnbr

import "github.com/01-edu/z01"

func PrintNbr(x int) {
	for {
		if x == 0 {
			break
		}

		z01.PrintRune('0' + x%10)
		x /= 10

	}
	z01.PrintRune('\n')
}
