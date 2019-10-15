package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				for w := j; w <= '9'; w++ {
					if i == k {
						w++
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(w)
					if i != '9' || j != '8' || k != '9' || w != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}

			}
		}
	}
	z01.PrintRune('\n')
}
