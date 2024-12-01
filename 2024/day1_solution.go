package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	filePath = "day1_input.txt"
)

func day1() {

	var list1, list2 []int

	p1_input := readFileContents("day1_input.txt")
	for i := range p1_input {
		s := strings.Fields(p1_input[i])

		s_to_i, _ := strconv.Atoi(s[0])
		list1 = append(list1, s_to_i)

		s_to_i, _ = strconv.Atoi(s[1])
		list2 = append(list2, s_to_i)
	}

	fmt.Printf("Day 1 - Historian Hysteria, p1 result: %d \n", day1Part1(list1, list2))

	fmt.Printf("Day 1 - Historian Hysteria, p2 result: %d \n", day1Part2(list1, list2))
}

/*
Challenge:
Pair up the smallest number in the left list with the smallest number in the right list, then the
second-smallest left number with the second-smallest right number, and so on.

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those
distances. For example, if you pair up a 3 from the left list with a 7 from the right list, the
distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.
*/
func day1Part1(list1 []int, list2 []int) int {

	sort.Ints(list1)
	sort.Ints(list2)

	var totalDistance int

	for i := range list1 {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance *= -1
		}
		totalDistance += distance
	}
	return totalDistance
}

/*
Challenge:
This time, you'll need to figure out exactly how often each number from the left list appears in the
right list. Calculate a total similarity score by adding up each number in the left list after
multiplying it by the number of times that number appears in the right list.
*/
func day1Part2(list1 []int, list2 []int) int {

	sort.Ints(list1)
	sort.Ints(list2)

	list2Counts := map[int]int{}

	// build a lookup map of counts
	for _, v := range list2 {
		list2Counts[v]++
	}

	var totalSimilarity int

	for _, v := range list1 {
		totalSimilarity += v * list2Counts[v]
	}
	return totalSimilarity
}
