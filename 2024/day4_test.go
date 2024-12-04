package main

import (
	"testing"
)

var testInput = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestDay4Part1(t *testing.T) {

	if result := day4Part1(testInput); result != 18 {
		t.Errorf("got %v", result)
	}
}

func TestIsXmasForward(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input:       []string{"XMAS"},
			row:         0,
			col:         0,
			result:      true,
		},
		{
			description: "end of the string",
			input:       []string{"XSWEFAXMAS"},
			row:         0,
			col:         6,
			result:      true,
		},
		{
			description: "off end of the string",
			input:       []string{"XSWEFASXMA"},
			row:         0,
			col:         7,
			result:      false,
		},
		{
			description: "incomplete word",
			input:       []string{"XSWEFAXMAM"},
			row:         0,
			col:         6,
			result:      false,
		},
	}

	for _, tc := range testCases {

		result := isXmasForward(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasBackward(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input:       []string{"SAMXASNEKS"},
			row:         0,
			col:         3,
			result:      true,
		},
		{
			description: "end of the string",
			input:       []string{"SWQNDSSAMX"},
			row:         0,
			col:         9,
			result:      true,
		},
		{
			description: "off the start of the string",
			input:       []string{"AMXASSDQW"},
			row:         0,
			col:         3,
			result:      false,
		},
	}

	for _, tc := range testCases {

		result := isXmasBackward(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasUpward(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input: []string{
				"S",
				"A",
				"M",
				"X",
			},
			row:    3,
			col:    0,
			result: true,
		},
		{
			description: "off start of the string",
			input: []string{
				"A",
				"M",
				"X",
			},
			row:    2,
			col:    0,
			result: false,
		},
	}

	for _, tc := range testCases {

		result := isXmasUpward(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasDownward(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input: []string{
				"X",
				"M",
				"A",
				"S",
			},
			row:    0,
			col:    0,
			result: true,
		},
		{
			description: "off end of input",
			input: []string{
				"X",
				"M",
				"A",
			},
			row:    0,
			col:    0,
			result: false,
		},
	}

	for _, tc := range testCases {

		result := isXmasDownward(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasDownwardRight(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input: []string{
				"X",
				"AM",
				"AAA",
				"SSSS",
			},
			row:    0,
			col:    0,
			result: true,
		},
		{
			description: "off end of the input",
			input: []string{
				"X",
				"AM",
				"AAA",
			},
			row:    0,
			col:    0,
			result: false,
		},
	}

	for _, tc := range testCases {

		result := isXmasDownwardRight(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasDownwardLeft(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input: []string{
				"ASSX",
				"AMM",
				"AA",
				"S",
			},
			row:    0,
			col:    3,
			result: true,
		},
		{
			description: "off end of the input",
			input: []string{
				"ASSX",
				"MAM",
				"AA",
			},
			row:    0,
			col:    3,
			result: false,
		},
	}

	for _, tc := range testCases {

		result := isXmasDownwardLeft(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasUpwardRight(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input: []string{
				"AAAS",
				"AAA",
				"AM",
				"X",
			},
			row:    3,
			col:    0,
			result: true,
		},
		{
			description: "off end of the input",
			input: []string{
				"AAA",
				"AM",
				"X",
			},
			row:    2,
			col:    0,
			result: false,
		},
	}

	for _, tc := range testCases {

		result := isXmasUpwardRight(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmasUpwardLeft(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      bool
	}{
		{
			description: "start of the string",
			input: []string{
				"S",
				"AAA",
				"MMM",
				"XXXX",
			},
			row:    3,
			col:    3,
			result: true,
		},
		{
			description: "off end of the input",
			input: []string{
				"AA",
				"MMM",
				"XXX",
			},
			row:    2,
			col:    2,
			result: false,
		},
	}

	for _, tc := range testCases {

		result := isXmasUpwardLeft(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v", result)
		}
	}
}

func TestIsXmas(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		description string
		input       []string
		row         int
		col         int
		result      int32
	}{
		{
			description: "start of the string",
			input: []string{
				"M.S",
				".A.",
				"M.S",
			},
			row:    1,
			col:    1,
			result: 1,
		},
		{
			description: "start of the string",
			input: []string{
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
			},
			row:    1,
			col:    7,
			result: 1,
		},
	}

	for _, tc := range testCases {

		result := isXmas2(tc.row, tc.col, tc.input)

		if result != tc.result {
			t.Errorf("got %v %s", result, tc.description)
		}
	}
}

// func TestIsXmas2(t *testing.T) {

// 	result := day4Part2(testInput)

// 	if result != 9 {
// 		t.Errorf("got %v", result)
// 	}
// }
