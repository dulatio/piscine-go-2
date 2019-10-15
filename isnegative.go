package isnegative

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T\n')
	} else {
		z01.PrintRune('F\n')
	}
}
