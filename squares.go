package main

type square [][]rune

// newSquare creates an empty square with a given side length
func newSquare(side int) (sq square) {
	for i := 0; i < side; i++ {
		row := []rune{}
		for j := 0; j < side; j++ {
			row = append(row, '.')
		}
		sq = append(sq, row)
	}
	return
}

// fillSquare places a tetronomino into the square at a given place ([2]int)
func fillSquare(t tetro, plc place, sq square) square {
	for i, row := range t {
		for j, r := range row {
			if r != '.' {
				sq[i+plc[0]][j+plc[1]] = r
			}
		}
	}
	return sq
}

// clearSquare removes a tetronomino from the square by index
func clearSquare(t int, sq square) square {
	for i, row := range sq {
		for j, r := range row {
			if r == rune(int('A')+t) {
				sq[i][j] = '.'
			}
		}
	}
	return sq
}
