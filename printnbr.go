package printnbr

import "github.com/01-edu/z01"

func PrintNbr(x int) {
	if x == 0 {
		z01.PrintRune('0')
	}
	for {
		if x < 0 {
			z01.PrintRune('-')
			x *= -1
		}
		if x == 0 {
			break
		}

		i := '0'
		for j := 1; j <= x%10; j++ {
			i++
		}
		z01.PrintRune(i)
		x /= 10
	}
	z01.PrintRune('\n')
}
