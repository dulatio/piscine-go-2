package basicatoi

package main

import (
	"fmt"
)

func get(x int) int {
	switch x {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	case 4:
		return 10000
	case 5:
		return 100000
	case 6:
		return 1000000
	case 7:
		return 10000000
	case 8:
		return 100000000
	case 9:
		return 1000000000
	case 10:
		return 10000000000
	case 11:
		return 100000000000
	case 12:
		return 1000000000000
	case 13:
		return 10000000000000
	case 14:
		return 100000000000000
	case 15:
		return 1000000000000000
	case 16:
		return 10000000000000000
	case 17:
		return 100000000000000000
	case 18:
		return 1000000000000000000

	}
	return -1
}
func BasicAtoi(str string) int {
	ok := false
	x := 0
	ln := 0
	for _, c := range str {
		if c != 0 {
			ln++
		}
	}
	cur_ln := 0
	for _, c := range str {
		cur_ln++
		if c != 48 {
			ok = true
		}
		if ok == false {
			continue
		}
		var cnt int
		for i := '1'; i <= c; i++ {
			cnt++
		}

		if cur_ln == 1 {
			x = cnt * get(ln-cur_ln)
		} else {
			x = x + cnt*get(ln-cur_ln)
		}

	}
	return x
}

