package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	filePath = "day1_input.txt"
)

/*
$ go run solution.go
Day 1 p1: Trebuchet?! - sum of calibration values: 55172
Day 1 p2: Trebuchet?! - sum of calibration values: 54925
*/
func main() {
	fmt.Printf("Day 1 p1: Trebuchet?! - sum of calibration values: %d \n", part1())
	fmt.Printf("Day 1 p2: Trebuchet?! - sum of calibration values: %d \n", part2())
}

/*
Challenge:

On each line, the calibration value can be found by combining the first digit and
the last digit (in that order) to form a single two-digit number.

- 0 is byte value 48, and 9 is 57. So we can detect digits by testing their byte
value.
- the first digit (in the combined calibration value) is in the tens, so
sum += f*10 + l
- keep things efficient by working line by line.

$ go solution.go
Day 1: Trebuchet?! - sum of calibration values: 55172
*/
func part1() int {

	readFile, _ := os.Open(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var sum int

	for fileScanner.Scan() {

		var f, l int

		linestring := fileScanner.Text()

		for _, c := range linestring {

			// set f when the first digit is detected
			if f == 0 && int(c) >= 48 && int(c) <= 57 {
				f = int(c) - 48
			}
			// always set l to current digit we automatically capture the last
			if int(c) >= 48 && int(c) <= 57 {
				l = int(c) - 48
			}
		}
		sum = sum + (f * 10) + l
	}
	return sum
}

/*
Challenge:

Your calculation isn't quite right. It looks like some of the digits are actually
spelled out with letters: one, two, three, four, five, six, seven, eight, and
nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last
digit on each line.

- this suits using an array of strings, the index (int) matching the number
*/
func part2() int {

	// include zero because it is keeps the indexing number-like
	words := []string{"zero", "one", "two",
		"three", "four", "five", "six",
		"seven", "eight", "nine"}

	readFile, _ := os.Open(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var sum int

	for fileScanner.Scan() {

		linestring := fileScanner.Text()

		var f, l int
		for i, c := range linestring {

			var word int

			// detect a number word at this point in the linestring
			for n, m := range words {
				if strings.Index(linestring[i:], m) == 0 {
					word = n
				}
			}

			// set f when the first digit remains undetected
			if f == 0 {
				// if this rune is a number
				if int(c) >= 48 && int(c) <= 57 {
					f = int(c) - 48
				} else if word > 0 {
					f = word
				}
			}

			// always set l to current digit we automatically capture the last
			if int(c) >= 48 && int(c) <= 57 {
				l = int(c) - 48
			} else if word > 0 {
				l = word
			}
		}
		sum = sum + (f * 10) + l
	}
	return sum
}
