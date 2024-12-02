package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2() {

	var reports [][]int

	p1_input := readFileContents("day2_input.txt")
	for i := range p1_input {
		fields := strings.Fields(p1_input[i])

		var report []int

		for _, f := range fields {
			s_to_i, _ := strconv.Atoi(f)
			report = append(report, s_to_i)
		}

		reports = append(reports, report)
	}

	fmt.Printf("Day 2 - Red-Nosed Reports, p1 result: %d \n", day2Part1(reports))

	fmt.Printf("Day 2 - Red-Nosed Reports, p2 result: %d \n", day2Part2(reports))

}

/*
Challenge:
The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems
can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report
only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
*/
func day2Part1(reports [][]int) int {

	count := 0
	for _, report := range reports {
		if reportIsSafe(report) {
			count++
		}
	}

	return count
}

func reportIsSafe(report []int) bool {

	i := 0
	trend := 0
	for i < len(report)-1 {

		change := report[i+1] - report[i]
		diff := diff(report[i+1], report[i])

		if diff < 1 || diff > 3 {
			return false
		}

		// if decreasing and already trending up
		if change < 0 && trend > 0 {
			return false
		}
		// if change is increasing and already trending down
		if change > 0 && trend < 0 {
			return false
		}
		trend = change

		i++
	}

	return true
}

/*
Challenge:
The engineers are surprised by the low number of safe reports until they realize they forgot to tell
you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a
single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would
make it safe, the report instead counts as safe.
*/
func day2Part2(reports [][]int) int {

	count := 0
	for _, report := range reports {
		if reportIsSafeWithDamper(report) {
			count++
		}
	}

	return count
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func removeIndex(s []int, index int) []int {

	newSlice := []int{}

	for i, v := range s {
		if i != index {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

// apply problem dampener
func reportIsSafeWithDamper(report []int) bool {

	if reportIsSafe(report) {
		return true
	}

	i := 0
	for i <= len(report)-1 {
		newReport := removeIndex(report, i)
		if reportIsSafe(newReport) {
			return true
		}
		i++
	}

	return false
}
