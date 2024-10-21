package main

import (
	"fmt"
	"math"
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

// placeIsLegal tells if the tetronomino fits in a given place
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

// placeTetros puts tetronominoes into as small a square as possible
func placeTetros(side int, tetros []tetro) square {
	square := newSquare(side)

	lengthTet := len(tetros)
	tetroI := 0        // Index of current tetromino
	var curTet tetro   // Current tetromino
	var places []place // Slice of possible coordinates of a tetromino's upper left corner inside a square
	placeI := 0        // Index of current coordinate
	var lengthPlcs int

	allPlacements := [][]place{}
	placementLenghts := make([]int, len(tetros))

	for _, tet := range tetros {
		allPlacements = append(allPlacements, getPlacements(tet, side))
	}
	for i, pls := range allPlacements {
		placementLenghts[i] = len(pls)
	}

	// Slice of coordinates in current solution, used with placeI
	foundPlaces := make([]int, lengthTet)

	// loop util a solution is found
	for {
		// solution complete, exit loop
		if tetroI > lengthTet-1 {
			break
		}

		curTet = tetros[tetroI]
		places = allPlacements[tetroI]
		lengthPlcs = placementLenghts[tetroI]

		// all places failed for first tetro, expanding square
		if tetroI == 0 && placeI > lengthPlcs-1 {
			side++

			allPlacements = [][]place{}
			for _, te := range tetros {
				allPlacements = append(allPlacements, getPlacements(te, side))
			}
			for i, pls := range allPlacements {
				placementLenghts[i] = len(pls)
			}

			square = newSquare(side)
			placeI = 0
			tetroI = 0
			continue
		}

		// current tetro can't fit, go back to previous tetro
		if placeI > lengthPlcs-1 {
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

func sideLen(tets []tetro) int {
	return int(math.Sqrt(float64(len(tets) * 4)))
}

// main reads an input file and prints found tetronominoes in as small
// a square as possible
func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Provide one text file as argument")
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}

	bigTetros := getBigTetros(string(bytes))
	if !checkBigTetros(bigTetros) {
		fmt.Println("ERROR")
		return
	}
	tetros := makeTetros(bigTetros) // The given tetrominoes in smallest possible rectangles

	// Minimum side length is the square root of the number of squares in all tetros
	side := sideLen(tetros)

	square := placeTetros(side, tetros)

	for _, row := range square {
		fmt.Println(row)
	}
}
