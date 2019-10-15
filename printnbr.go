package printnbr

import "github.com/01-edu/z01"

func PrintNbr(x int) {
	if x == 0 {
		z01.PrintRune('0')
	}

		if x < 0 {
			z01.PrintRune('-')
			x *= -1
		}
    y:=x
    ln := 1
    for{
      if y == 0 {
        break;}
        y /= 10
        ln *= 10;
    }
    for{
      if x == 0 {
      break;
      }
      w := '0'
      for j:= 0;j <= x / ln; j ++ {
        w++;
      }
      z01.PrintRune(w)
      x %= ln;
      ln /= 10;
    }
	}
}
