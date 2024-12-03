package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func day3() {

	input := readFileContents("day3_input.txt")

	fmt.Printf("Day 3 - Mull It Over, p1 result: %d \n", day3Part1(input))
	fmt.Printf("Day 2 - Mull It Over, p2 result: %d \n", day3Part2(input))

}

/*
Challenge:
It seems like the goal of the program is just to multiply some numbers. It does that with
instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46)
multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the
results of the multiplications?
*/
func day3Part1(input []string) int32 {

	var mul = regexp.MustCompile(`mul\(([0-9]+)\,([0-9]+)\)`)

	var total int32

	for _, l := range input {

		matches := mul.FindAllStringSubmatch(l, -1)

		for _, m := range matches {
			firstNumber, err := strconv.Atoi(string(m[1]))
			if err != nil {
				log.Println(err)
			}
			secondNumber, err := strconv.Atoi(string(m[2]))
			if err != nil {
				log.Println(err)
			}

			total += int32(firstNumber) * int32(secondNumber)

		}
	}

	return total
}

/*
Challenge:
As you scan through the corrupted memory, you notice that some of the conditional statements are
also still intact. If you handle some of the uncorrupted conditional statements in the program, you
might be able to get an even more accurate result.

Only the most recent do() or don't() instruction applies. At the beginning of the program, mul
instructions are enabled.
*/
func day3Part2(input []string) int32 {

	var total int32

	// conditionals may spread across lines
	inputJoined := strings.Join(input, "")

	// first clean the input data, by removing all character sections 'dont()...do()'
	var re = regexp.MustCompile(`don\'t\(\).+?do\(\)`)
	s := re.ReplaceAllString(inputJoined, "")

	total += day3Part1([]string{s})

	return total
}
