package main

import (
	"fmt"
	"log"
	"math"
)

func day6() {
	day6_part2()
	// quadratic(1, t, d)
}

/*
(py3.9) dougb $ go run day6_solution.go
time 46689866, distance 358105418071080, lower 9674510, upper 37015356, total 27340847
total 27340847
*/
func day6_part2() {

	//input := [][]int64{{7, 9}, {15, 40}, {30, 200}}
	input := [][]int64{{46689866, 358105418071080}}

	total := int64(1)
	for _, pair := range input {
		lower, upper := quadratic(1, pair[0], pair[1])
		fmt.Printf("range 0...%d, lower bound %d, upper bound %d, total combinations %d\n", pair[0], lower, upper, upper-lower+1)
		total *= upper - lower + 1
	}
	fmt.Printf("total %d\n", total)
}

/*
The problem can be solved via the quadratic equation:

	    distance = (time - charge) * charge
		d = (t - c) * c
		d = c(t - c)
		d = ct - c^2
		0 = c^2 - ct + d

		or, given inputs of t = 7 and d = 9:

		x^2 - 7x + 9

This function returns the x-intercepts of the graph, where y=0.
*/
func quadratic(a int64, b int64, c int64) (int64, int64) {

	fa := float64(a)
	fb := float64(b)
	fc := float64(c + 1)

	sqrt := math.Sqrt(math.Pow(fb, 2) - (4 * fa * fc))

	x1 := (fb - sqrt) / 2.0
	x2 := (fb + sqrt) / 2.0

	return int64(math.Ceil(x1)), int64(math.Floor(x2))
}

/*
press is the time charging the boat
time is the total time the race lasts
*/
func calcDistance(press int64, time int64) int64 {

	distance := (time - press) * press

	return distance
}

func day6_part1() {

	input := [][]int64{{7, 9}, {15, 40}, {30, 200}}
	//input2 := [][]int64{{46, 358}, {68, 1054}, {98, 1807}, {66, 1080}}

	result := int64(1)
	for _, pair := range input {
		result *= f(pair[0], pair[1])
	}
	log.Println(result)
}

func f(time int64, d int64) int64 {

	var ways int64
	for charge := int64(0); charge <= time; charge++ {

		distance := calcDistance(charge, time)

		if distance > d {
			ways += 1

			fmt.Printf("charge %d ms, for distance %d\n",
				charge,
				distance,
			)
		}
	}
	return ways
}
