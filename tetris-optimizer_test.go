package main

import (
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

var testCasesBigTetros = []testCaseBad{
	{
		name:     "bad 00",
		input:    bad00,
		expected: false,
	},

	{
		name:     "bad 01",
		input:    bad01,
		expected: false,
	},
	{
		name:     "bad 02",
		input:    bad02,
		expected: false,
	},
	{
		name:     "bad 03",
		input:    bad03,
		expected: false,
	},
	{
		name:     "bad 04",
		input:    bad04,
		expected: false,
	},
	{
		name:     "bad format",
		input:    bad05,
		expected: false,
	},
	{
		name:     "good 00",
		input:    good00,
		expected: true,
	},
	{
		name:     "good 01",
		input:    good01,
		expected: true,
	},
	{
		name:     "good 02",
		input:    good02,
		expected: true,
	},
	{
		name:     "good 03",
		input:    good03,
		expected: true,
	},
	{
		name:     "good hard",
		input:    goodHard,
		expected: true,
	},
}

var testCasesGood = []testCaseGood{
	{
		name:     "good 00",
		input:    makeTetros(getBigTetros(good00)),
		expected: 0,
	},
	{
		name:     "good 01",
		input:    makeTetros(getBigTetros(good01)),
		expected: 9,
	},
	{
		name:     "good 02",
		input:    makeTetros(getBigTetros(good02)),
		expected: 4,
	},
	{
		name:     "good 03",
		input:    makeTetros(getBigTetros(good03)),
		expected: 5,
	},
	// The last case takes ~80 seconds to complete
	/* 	{
		name:     "good hard",
		input:    makeTetros(getBigTetros(goodHard)),
		expected: 1,
	}, */
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
