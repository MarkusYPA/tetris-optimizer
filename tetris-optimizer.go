package main

import (
	"fmt"
	"io"
	"os"
)

type place [2]int

// getPlacements returns possible coordinates for the top left corner of the teronomino in the square
func getPlacements(t tetro, side int) []place {
	places := []place{}
	for i := 0; i < side+1-t.width(); i++ {
		for j := 0; j < side+1-t.height(); j++ {
			places = append(places, place{j, i})
		}
	}
	return places
}

// placeIsLegal tells if there's room in the square for the tetronomino at a given place
func placeIsLegal(t tetro, plc place, sq square) bool {
	for i, row := range t {
		for j, r := range row {
			// if both square and tetro have content at this coordinate point
			if sq[i+plc[0]][j+plc[1]] != '.' && r != '.' {
				return false
			}
		}
	}
	return true
}

// placeTetros puts tetronominos into as small a square as possible
func placeTetros(side int, tetros []tetro) square {
	square := newSquare(side)

	tetroI := 0        // Index of current tetromino
	var curTet tetro   // Current tetromino
	var places []place // Slice of possible coordinates of a tetromino's upper left corner inside a square
	placeI := 0        // Index of currrent coordinate

	allPlacements := [][]place{}
	for _, te := range tetros {
		allPlacements = append(allPlacements, getPlacements(te, side))
	}

	// Slice of coordinates in current solution, used with placeI
	foundPlaces := make([]int, len(tetros))

	// loop util a solution is found
	for {
		// solution complete, exit loop
		if tetroI > len(tetros)-1 {
			break
		}

		curTet = tetros[tetroI]
		places = allPlacements[tetroI]

		// all places failed for first tetro, expanding square
		if tetroI == 0 && placeI > len(places)-1 {
			side++

			allPlacements = [][]place{}
			for _, te := range tetros {
				allPlacements = append(allPlacements, getPlacements(te, side))
			}

			square = newSquare(side)
			placeI = 0
			tetroI = 0
			continue
		}

		// current tetro can't fit, go back to previous tetro
		if placeI > len(places)-1 {
			tetroI--
			placeI = foundPlaces[tetroI] + 1
			foundPlaces[tetroI] = 0
			square = clearSquare(tetroI, square)
			continue
		}

		// place tetro into square or increment to next place
		if placeIsLegal(curTet, places[placeI], square) {
			square = fillSquare(curTet, places[placeI], square)
			foundPlaces[tetroI] = placeI
			tetroI++
			placeI = 0
			continue
		} else {
			placeI++
			continue
		}
	}

	return square
}

// main reads an iput file and prints found tetronominos in as small
// a square as possible
func main() {
	// open and read the file from the first argument
	tetrosFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer tetrosFile.Close()

	bytes, err := io.ReadAll(tetrosFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bigTetros := getBigTetros(string(bytes))
	if !checkBigTetros(bigTetros) {
		fmt.Println("ERROR")
		return
	}
	tetros := makeTetros(bigTetros) // The given tetrominos in smallest possible rectangles

	// start with a square the size of the largest tetro
	side := 0
	for _, t := range tetros {
		if t.width() > side {
			side = t.width()
		}
		if t.height() > side {
			side = t.height()
		}
	}

	square := placeTetros(side, tetros)

	for _, row := range square {
		fmt.Println(row)
	}
}
