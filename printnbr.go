package printnbr

import "github.com/01-edu/z01"

func PrintNbr(x int) {
  while(x){
    z01.PrintRune(x % 10 + '0')
    x /= 10
  }
	z01.PrintRune('\n')
}
