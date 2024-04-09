package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	filePath = "input.txt"
)

/*
Challenge:

On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

- 0 is byte value 48, and 9 is 57. So we can detect digits by testing their byte value.
- the first digit (in the combined calibration value) is in the tens
- keep things efficient by working line by line.

$ go solution.go
Day 1: Trebuchet?! - sum of all of the calibration values: 55172
*/
func main() {

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
				f = (int(c) - 48) * 10 // when combing f is in the tens
			}
			// by always setting l to the current digit we automatically capture the last
			if int(c) >= 48 && int(c) <= 57 {
				l = int(c) - 48
			}
		}
		sum = sum + f + l
	}

	fmt.Printf("Day 1: Trebuchet?! - sum of all of the calibration values: %d \n", sum)
}
