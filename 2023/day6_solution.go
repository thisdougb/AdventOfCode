package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
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
		fmt.Printf("time %d, distance %d, lower %d, upper %d, total %d\n", pair[0], pair[1], lower, upper, upper-lower+1)
		total *= upper - lower + 1
	}
	fmt.Printf("total %d\n", total)
}

/*
x^2 -7x + 9
*/
func quadratic(a int64, b int64, c int64) (int64, int64) {

	fa := float64(a)
	fb := float64(b)
	fc := float64(c + 1)

	b2 := math.Pow(fb, 2)
	n := b2 - (4 * fa * fc)
	sqrt := math.Sqrt(float64(n))

	x1 := (fb - sqrt) / float64(2)
	x2 := (fb + sqrt) / float64(2)

	return int64(math.Ceil(x1)), int64(math.Floor(x2))
}

/*
press is the time charing the boat
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
