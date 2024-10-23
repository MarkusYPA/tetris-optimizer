package main

import (
	"log"
	"os"
	"testing"
)

type testCaseBad struct {
	name     string
	input    string
	expected bool
}

type testCaseGood struct {
	name     string
	input    []tetro
	expected int
}

func fileToString(s string) string {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(file)
}

var testCasesBigTetros = []testCaseBad{
	{
		name:     "bad 00",
		input:    fileToString("testcases/bad00.txt"),
		expected: false,
	},

	{
		name:     "bad 01",
		input:    fileToString("testcases/bad01.txt"),
		expected: false,
	},
	{
		name:     "bad 02",
		input:    fileToString("testcases/bad02.txt"),
		expected: false,
	},
	{
		name:     "bad 03",
		input:    fileToString("testcases/bad03.txt"),
		expected: false,
	},
	{
		name:     "bad 04",
		input:    fileToString("testcases/bad04.txt"),
		expected: false,
	},
	{
		name:     "bad format",
		input:    fileToString("testcases/bad05.txt"),
		expected: false,
	},
	{
		name:     "good 00",
		input:    fileToString("testcases/good00.txt"),
		expected: true,
	},
	{
		name:     "good 01",
		input:    fileToString("testcases/good01.txt"),
		expected: true,
	},
	{
		name:     "good 02",
		input:    fileToString("testcases/good02.txt"),
		expected: true,
	},
	{
		name:     "good 03",
		input:    fileToString("testcases/good03.txt"),
		expected: true,
	},
	{
		name:     "good hard",
		input:    fileToString("testcases/goodHard.txt"),
		expected: true,
	},
}

var testCasesGood = []testCaseGood{
	{
		name:     "good 00",
		input:    makeTetros(getBigTetros(fileToString("testcases/good00.txt"))),
		expected: 0,
	},
	{
		name:     "good 01",
		input:    makeTetros(getBigTetros(fileToString("testcases/good01.txt"))),
		expected: 9,
	},
	{
		name:     "good 02",
		input:    makeTetros(getBigTetros(fileToString("testcases/good02.txt"))),
		expected: 4,
	},
	{
		name:     "good 03",
		input:    makeTetros(getBigTetros(fileToString("testcases/good03.txt"))),
		expected: 5,
	},
	// The last case takes ~5 seconds to complete
	{
		name:     "good hard",
		input:    makeTetros(getBigTetros(fileToString("testcases/goodHard.txt"))),
		expected: 1,
	},
}

// countDots returns the numbers of dots (empty cells) in a square
func countDots(sq square) int {
	dots := 0
	for _, row := range sq {
		for _, r := range row {
			if r == '.' {
				dots++
			}
		}
	}
	return dots
}

func TestCheckBigTetros(t *testing.T) {
	for _, tc := range testCasesBigTetros {
		t.Run(tc.name, func(t *testing.T) {
			result := checkBigTetros(getBigTetros(tc.input))
			if tc.expected != result {
				t.Errorf("\n%s Input was \"%s\"\nwant:\n%t\ngot:\n%t", tc.name, tc.input, tc.expected, result)
			}
		})
	}
}

func TestPlaceTetros(t *testing.T) {
	for _, tc := range testCasesGood {
		t.Run(tc.name, func(t *testing.T) {
			result := countDots(placeTetros(sideLen(tc.input), tc.input))
			if tc.expected != result {
				t.Errorf("\n%s Input was \"%s\"\nwant:\n%v\ngot:\n%v", tc.name, tc.input, tc.expected, result)
			}
		})
	}
}
