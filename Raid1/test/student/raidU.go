package student

import "github.com/01-edu/z01"

func RaidU(x int, y int) {
	a := 'A'
	b := 'C'
	c := 'C'
	d := 'A'
	wall := 'B'
	wall2 := 'B'
	matrix := make([][]rune, y)
	for i := 0; i < y; i++ {
		matrix[i] = make([]rune, x)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			matrix[i][j] = ' '
		}
	}

	matrix[y-1][x-1] = d

	matrix[y-1][0] = c

	matrix[0][x-1] = b

	matrix[0][0] = a
	for i := 1; i < x-1; i++ {
		matrix[0][i], matrix[y-1][i] = wall, wall
	}
	for i := 1; i < y-1; i++ {
		matrix[i][0], matrix[i][x-1] = wall2, wall2
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			z01.PrintRune(matrix[i][j])
		}
		z01.PrintRune('\n')
	}
}
