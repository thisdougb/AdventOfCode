package main

/*
Got 539219 which is wrong. Bug fixing...
- added unittests on a sample of the real data, found an index offset bug

Got 324159 which is also wrong...

Finally...

2024/04/10 16:54:04 Symbols: 738
2024/04/10 16:54:04 Parts: 1073
2024/04/10 16:54:04 Sum: 540212

the problem was duplicate parts are allowed in the schematic, I was squashing duplicates.

and simplified the logic.
*/
import (
	"log"
	"strconv"
)

type Symbol struct {
	char      int
	row       int
	col       int
	labels    int
	gears     []int
	gearRatio int
}

func isDigit(char int) bool {

	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

// given a single array location return the string of digits passing through it
func getLabelContaining(input *[]string, row int, col int) string {

	// make things a little more readable
	rowString := (*input)[row]

	labelStart := 0
	labelEnd := 0

	// first go back to find the start of the string
	for i := col; i >= 0; i-- {
		if !isDigit(int(rowString[i])) {
			break
		}
		labelStart = i
	}

	// now walk forward to find the word boundary
	for i := labelStart; i < len(rowString); i++ {
		if !isDigit(int(rowString[i])) {
			break
		}
		labelEnd = i
	}

	// return a slice so +1 the upper bound
	return rowString[labelStart : labelEnd+1]
}

func isLabel(input *[]string, row int, col int) (string, bool) {

	if row < 0 || col < 0 {
		return "", false
	}
	if row > len((*input))-1 {
		return "", false
	}
	if col > len((*input)[row]) {
		return "", false
	}

	if isDigit(int((*input)[row][col])) {
		return getLabelContaining(input, row, col), true
	}

	return "", false
}

func day3() {
	//log.Println(day3Part1(readFileContents("test.txt")))
	log.Println("Sum:", day3Part1(readFileContents("day3_input.txt")))
}

func day3Part1(input []string) int {

	totalParts := []int{}
	totalGears := 0

	// get all symbols in the input, with their row,col location
	symbols := getAllSymbols(input)
	for _, s := range symbols {

		// I found this the most readable way to walk round the
		// adjacent char positions. Let the called function handle
		// bounds checking.

		partNumbers := make(map[string]bool)
		// above left
		if label, result := isLabel(&input, s.row-1, s.col-1); result {
			partNumbers[label] = true
			s.labels++
		}

		// above
		if label, result := isLabel(&input, s.row-1, s.col); result {
			partNumbers[label] = true
			s.labels++
		}
		// above right
		if label, result := isLabel(&input, s.row-1, s.col+1); result {
			partNumbers[label] = true
			s.labels++
		}
		// remove duplicate labels this symbol is connected to
		for k, _ := range partNumbers {
			n, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			totalParts = append(totalParts, n)
			s.gears = append(s.gears, n)
		}

		partNumbers = make(map[string]bool)
		// below right
		if label, result := isLabel(&input, s.row+1, s.col+1); result {
			partNumbers[label] = true
			s.labels++
		}
		// below
		if label, result := isLabel(&input, s.row+1, s.col); result {
			partNumbers[label] = true
			s.labels++
		}
		// below left
		if label, result := isLabel(&input, s.row+1, s.col-1); result {
			partNumbers[label] = true
			s.labels++
		}
		// remove duplicate labels this symbol is connected to
		for k, _ := range partNumbers {

			n, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			totalParts = append(totalParts, n)
			s.gears = append(s.gears, n)
		}

		partNumbers = make(map[string]bool)
		// left
		if label, result := isLabel(&input, s.row, s.col-1); result {
			partNumbers[label] = true
			s.labels++
		}
		// right
		if label, result := isLabel(&input, s.row, s.col+1); result {
			partNumbers[label] = true
			s.labels++
		}
		// remove duplicate labels this symbol is connected to
		for k, _ := range partNumbers {

			n, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			totalParts = append(totalParts, n)
			s.gears = append(s.gears, n)
		}

		if s.labels == 0 {
			log.Println("WARNING: No labels", s)
		}

		if s.char == '*' && len(s.gears) == 2 {
			totalGears += s.gears[0] * s.gears[1]
		}
	}

	log.Println("Symbols:", len(symbols))
	log.Println("Parts:", len(totalParts))
	log.Println("Total Gears:", totalGears)

	total := 0
	for _, p := range totalParts {
		total += p
	}

	return total
}

func getAllSymbols(input []string) []Symbol {

	symbols := []Symbol{}

	for row, line := range input {
		var rowsymbols int
		for col, char := range line {
			if !isDigit(int(char)) && char != '.' {
				symbols = append(symbols, Symbol{int(char), row, col, 0, []int{}, 0})
				rowsymbols++
			}
		}
		//	fmt.Printf("%02d %s\n", rowsymbols, line)
	}

	return symbols
}
