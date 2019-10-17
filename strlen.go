package strlen

func StrLen(str string) int {
	ln := 0
	for _, c := range str {
		ln++
	}
	return ln
}
