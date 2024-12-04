package main

import (
	"fmt"
)

func day4() {

	input := readFileContents("day4_input.txt")

	fmt.Printf("Day 4 - Mull It Over, p1 result: %d \n", day4Part1(input))
	fmt.Printf("Day 4 - Mull It Over, p2 result: %d \n", day4Part2(input))

}

/*
Challenge:
*/
func day4Part1(input []string) int32 {

	var total int32

	for rowNum, row := range input {
		for colNum, char := range row {
			if string(char) == "X" {
				count := isXmas(rowNum, colNum, input)
				if count > 0 {
					total += count
				}
			}
		}
	}

	return total
}

func isXmas(row int, col int, input []string) int32 {

	var count int32

	if isXmasForward(row, col, input) {
		count++
	}
	if isXmasBackward(row, col, input) {
		count++
	}
	if isXmasUpward(row, col, input) {
		count++
	}
	if isXmasUpwardRight(row, col, input) {
		count++
	}
	if isXmasUpwardLeft(row, col, input) {
		count++
	}
	if isXmasDownward(row, col, input) {
		count++
	}
	if isXmasDownwardRight(row, col, input) {
		count++
	}
	if isXmasDownwardLeft(row, col, input) {
		count++
	}
	return count
}

func isXmasForward(row int, col int, input []string) bool {

	if col+3 >= len(input[row]) {
		return false
	}

	if string(input[row][col+1]) == "M" && string(input[row][col+2]) == "A" && string(input[row][col+3]) == "S" {
		return true
	}

	return false
}

func isXmasBackward(row int, col int, input []string) bool {

	if col-3 < 0 {
		return false
	}

	if string(input[row][col-1]) == "M" && string(input[row][col-2]) == "A" && string(input[row][col-3]) == "S" {
		return true
	}

	return false
}

func isXmasUpward(row int, col int, input []string) bool {

	if row-3 < 0 {
		return false
	}

	if string(input[row-1][col]) == "M" && string(input[row-2][col]) == "A" && string(input[row-3][col]) == "S" {
		return true
	}

	return false
}

func isXmasDownward(row int, col int, input []string) bool {

	if row+3 >= len(input) {
		return false
	}

	if string(input[row+1][col]) == "M" && string(input[row+2][col]) == "A" && string(input[row+3][col]) == "S" {
		return true
	}

	return false
}

func isXmasUpwardRight(row int, col int, input []string) bool {

	if row-3 < 0 || col+3 >= len(input[0]) {
		return false
	}

	if string(input[row-1][col+1]) == "M" && string(input[row-2][col+2]) == "A" && string(input[row-3][col+3]) == "S" {
		return true
	}

	return false
}

func isXmasUpwardLeft(row int, col int, input []string) bool {

	if row-3 < 0 || col-3 < 0 {
		return false
	}

	if string(input[row-1][col-1]) == "M" && string(input[row-2][col-2]) == "A" && string(input[row-3][col-3]) == "S" {
		return true
	}

	return false
}

func isXmasDownwardRight(row int, col int, input []string) bool {

	if row+3 >= len(input) || col+3 >= len(input[row+3]) {
		return false
	}

	if string(input[row+1][col+1]) == "M" && string(input[row+2][col+2]) == "A" && string(input[row+3][col+3]) == "S" {
		return true
	}

	return false
}

func isXmasDownwardLeft(row int, col int, input []string) bool {

	if row+3 >= len(input) || col-3 < 0 {
		return false
	}

	if string(input[row+1][col-1]) == "M" && string(input[row+2][col-2]) == "A" && string(input[row+3][col-3]) == "S" {
		return true
	}

	return false
}

/*
Challenge:
*/
func day4Part2(input []string) int32 {

	var total int32

	for rowNum, row := range input {
		for colNum, char := range row {
			if string(char) == "A" {
				count := isXmas2(rowNum, colNum, input)
				if count > 0 {
					total += count
				}
			}
		}
	}

	return total
}

func isXmas2(row int, col int, input []string) int32 {

	var count int32

	// if the above left is within bounds
	if row-1 < 0 || col-1 < 0 {
		return count
	}

	// if the above right is within bounds
	if row-1 < 0 || col+1 >= len(input[0]) {
		return count
	}

	// if the below left is within bounds
	if row+1 >= len(input) || col-1 < 0 {
		return count
	}

	// if the below right is within bounds
	if row+1 >= len(input) || col+1 >= len(input[0]) {
		return count
	}

	// all wtithin bounds
	// start above left == M
	if string(input[row-1][col-1]) == "M" && string(input[row+1][col+1]) == "S" ||
		// opposite, above left == S
		string(input[row-1][col-1]) == "S" && string(input[row+1][col+1]) == "M" {

		// above right == M
		if string(input[row-1][col+1]) == "M" && string(input[row+1][col-1]) == "S" ||
			// oppositem, above right == S
			string(input[row-1][col+1]) == "S" && string(input[row+1][col-1]) == "M" {
			return 1
		}
	}

	return 0
}
