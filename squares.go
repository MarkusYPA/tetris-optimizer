package main

type square []string

// newSquare creates an empty square with a given side length
func newSquare(side int) (sq square) {
	for i := 0; i < side; i++ {
		row := ""
		for j := 0; j < side; j++ {
			row += "."
		}
		sq = append(sq, row)
	}
	return
}

// fillSquare places a tetronomino into the square at a given place ([2]int)
func fillSquare(t tetro, plc place, sq square) square {
	sqRunes := squareToRunes(sq)
	for i, row := range t {
		for j, r := range row {
			if r != '.' {
				sqRunes[i+plc[0]][j+plc[1]] = r
			}
		}
	}
	return runesToSquare(sqRunes)
}

// clearSquare removes a tetronomino from the square by index
func clearSquare(t int, sq square) square {
	sqRunes := squareToRunes(sq)

	for i, row := range sqRunes {
		for j, r := range row {
			if r == rune(int('A')+t) {
				sqRunes[i][j] = '.'
			}
		}
	}
	return runesToSquare(sqRunes)
}

// squareToRunes turns a square ([]string) to [][]rune for easier editing
func squareToRunes(sq square) [][]rune {
	sqRunes := [][]rune{}
	for i, row := range sq {
		sqRunes = append(sqRunes, []rune{})
		for _, r := range row {
			sqRunes[i] = append(sqRunes[i], r)
		}
	}
	return sqRunes
}

// runesToSquare turns a [][]rune back to square ([]string)
func runesToSquare(rs [][]rune) square {
	nuSq := square{}
	for i, row := range rs {
		nuSq = append(nuSq, "")
		for _, r := range row {
			nuSq[i] += string(r)
		}
	}
	return nuSq
}
