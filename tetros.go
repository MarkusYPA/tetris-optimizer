package main

type tetro []string

// width returns the width of a tetro
func (t tetro) width() int {
	if t == nil {
		return 0
	}
	return len(t[0])
}

// height returns the height of a tetro
func (t tetro) height() int {
	return len(t)
}

// getBigTetros returns a slice of 4x4 squares from a string
func getBigTetros(s string) [][]string {
	bigTetros := [][]string{{""}}
	sqIndex := 0
	sqRow := 0

	for i, r := range s {

		// start new square at two line changes
		if i > 1 && r == '\n' && s[i-1] == '\n' {
			bigTetros = append(bigTetros, []string{""})
			sqRow = 0
			sqIndex++
			continue
		}

		// start new row or add to current
		if r == '\n' {
			if sqRow != 3 {
				bigTetros[sqIndex] = append(bigTetros[sqIndex], "")
				sqRow++
			}
		} else {
			if r == '.' {
				bigTetros[sqIndex][sqRow] += string(r)
			} else {
				// Put letters in instead of #
				bigTetros[sqIndex][sqRow] += string(rune('A' + sqIndex))
			}
		}

	}
	return bigTetros
}

// emptyRow checks if a string is only made up of '.':s
func emptyRow(s string) bool {
	others := 0
	for _, r := range s {
		if r != '.' && r != '\n' {
			others++
		}
	}
	return others == 0
}

// makeTetros creates a slice of different sized rectangle tetros from 4x4 squares
func makeTetros(bts [][]string) []tetro {
	tetros := []tetro{}

	for _, bt := range bts {
		// list of rows with content
		rowsToKeep := []int{}
		for j, row := range bt {
			if !emptyRow(row) {
				rowsToKeep = append(rowsToKeep, j)
			}
		}

		// list of columns with content
		colsToKeep := []int{}
		cols := []int{0, 0, 0, 0}
		for _, row := range bt {
			for k, r := range row {
				if r != '.' && r != '\n' {
					cols[k]++
				}
			}
		}
		for j, col := range cols {
			if col != 0 {
				colsToKeep = append(colsToKeep, j)
			}
		}

		newTetro := make(tetro, len(rowsToKeep))
		for i, row := range rowsToKeep {
			for _, num := range colsToKeep {
				newTetro[i] += string(bt[row][num])
			}
		}
		tetros = append(tetros, newTetro)
	}
	return tetros
}

// checkBigTetros returns false if any of the tetroniminoes are invalid
func checkBigTetros(bts [][]string) bool {

	// check if all tetronominoes have 4 cells
	for _, t := range bts {
		nonEmpties := 0
		for _, row := range t {
			for _, r := range row {
				if r != '.' {
					nonEmpties++
				}
			}
		}
		if nonEmpties != 4 {
			return false
		}
	}

	// counting adjoining cells
	for i, t := range bts {

		adjoiningDatas := []int{}
		for j, row := range t {
			for k, r := range row {
				if r != '.' {
					adjoiningData := 0

					if j > 0 && bts[i][j-1][k] != '.' { // if the cell below isn't empty
						adjoiningData++
					}
					if j < len(t)-1 && bts[i][j+1][k] != '.' { // above
						adjoiningData++
					}
					if k > 0 && bts[i][j][k-1] != '.' { // left
						adjoiningData++
					}
					if k < len(row)-1 && bts[i][j][k+1] != '.' { // right
						adjoiningData++
					}

					adjoiningDatas = append(adjoiningDatas, adjoiningData)
				}
			}
		}

		ones := 0
		for _, ad := range adjoiningDatas {
			if ad == 0 {
				// all cells must have orthogonal neighbours
				return false
			}
			if ad == 1 {
				ones++
			}
		}

		// max three cells are allowed only 1 neighbour
		if ones > 3 {
			return false
		}
	}
	return true
}
