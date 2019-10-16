package printcombn

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

func check(p int) bool {
	for {
		if p == 0 {
			break
		}
		if p/10 != 0 && p%10 <= ((p/10)%10) {
			return false
		}
		p /= 10
	}
	return true
}
func ok(p int) bool {
	for {
		if p == 0 {
			break
		}
		if p/10 != 0 && p%10 != ((p/10)%10)+1 {
			return true
		}
		p /= 10
	}
	return false
}
func PrintCombN(n int) {
	mx_ln := 9
	for i := 2; i <= n; i++ {
		mx_ln *= 10
	}
	for i := mx_ln / 9 / 10; i < mx_ln/9; i++ {
		if check(i) == true {
			z01.PrintRune('0')
			GIVE(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')

		}
	}
	for i := mx_ln / 9; i <= mx_ln; i++ {
		if check(i) == true {
			GIVE(i)
			if ok(i) == true {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
