package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {

	var testCases = []struct {
		description string
		lines       []string
		expect      int
	}{
		{
			description: "example given",
			lines: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expect: 4361,
		},
		{
			description: "number touched twice by symbol",
			lines: []string{
				".467.114..",
				"...*......",
			},
			expect: 467,
		},

		{
			description: "same number twice",
			lines: []string{
				"467..114..",
				"...*......",
				"467..114..",
			},
			expect: 934,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expect, day3Part1(tc.lines), tc.description)
	}
}
func TestIsLabel(t *testing.T) {

	lines := []string{
		"467..114..",
		"...*...#..",
		"..35..633.",
		"....$70...",
		"190.......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
		"...4.598..",
	}

	var testCases = []struct {
		description string
		row         int
		col         int
		expect      bool
		expectLabel string
	}{
		{
			description: "row too small",
			expect:      false,
			row:         -1,
			col:         5,
		},
		{
			description: "row too large",
			expect:      false,
			row:         11,
			col:         5,
		},
		{
			description: "col too small",
			expect:      false,
			row:         1,
			col:         -1,
		},
		{
			description: "col too large",
			expect:      false,
			row:         1,
			col:         11,
		},
		{
			description: "left most label",
			expect:      true,
			row:         0,
			col:         2,
			expectLabel: "467",
		},
		{
			description: "single digit label",
			expect:      true,
			row:         10,
			col:         3,
			expectLabel: "4",
		},
	}

	for _, tc := range testCases {
		label, result := isLabel(&lines, tc.row, tc.col)
		assert.Equal(t, tc.expect, result, tc.description)
		assert.Equal(t, tc.expectLabel, label, tc.description)
	}
}
func TestFindSymbol(t *testing.T) {

	var testCases = []struct {
		description string
		lines       []string
		expect      []Symbol
	}{
		{
			description: "sample data",
			lines: []string{
				"467..114..",
				"...*...#..",
				"..35..633.",
				"....$70...",
				"190.......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expect: []Symbol{
				{'*', 1, 3, 0, []int{}, 0},
				{'#', 1, 7, 0, []int{}, 0},
				{'$', 3, 4, 0, []int{}, 0},
				{'+', 5, 5, 0, []int{}, 0},
				{'$', 8, 3, 0, []int{}, 0},
				{'*', 8, 5, 0, []int{}, 0},
			},
		},
	}

	for _, tc := range testCases {
		symbols := getAllSymbols(tc.lines)
		assert.Equal(t, tc.expect, symbols, tc.description)
	}
}

func TestGetLabelContaining(t *testing.T) {

	var testCases = []struct {
		description string
		line        []string
		row         int
		col         int
		expect      string
	}{
		{
			description: "left boundary",
			line:        []string{"467..114.."},
			row:         0,
			col:         0,
			expect:      "467",
		},
		{
			description: "near left boundary",
			line:        []string{".467.114.."},
			row:         0,
			col:         1,
			expect:      "467",
		},
		{
			description: "middle",
			line:        []string{"467..114.."},
			row:         0,
			col:         6,
			expect:      "114",
		},
		{
			description: "right boundary",
			line:        []string{"..467..889"},
			row:         0,
			col:         9,
			expect:      "889",
		},
		{
			description: "near right boundary",
			line:        []string{"..467.889."},
			row:         0,
			col:         8,
			expect:      "889",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expect, getLabelContaining(&tc.line, tc.row, tc.col), tc.description)
	}
}
