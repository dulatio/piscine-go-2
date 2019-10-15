package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for w := '0'; w <= '9'; w++ {
					if i == k && j < w || i < k {
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
	}
	z01.PrintRune('\n')
}
