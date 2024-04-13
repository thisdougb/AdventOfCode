package main

import (
	"fmt"
	"log"
)

func main() {
	day6_part2()
}

func day6_part2() {

	// 0 < t < 71530
	// where d > 940200-1
	input := [][]int64{{71530, 940200}}
	//input2 := [][]int64{{46689866, 358105418071080}}
	result := int64(1)
	for _, pair := range input {
		result *= f(pair[0], pair[1])
	}
	log.Println(result)
}

/*
find the lower and upper boundary where the charge results in a distance greater
than the target.
*/
func findLowBoundary(t int64, d int64) {

	// L = 0, R = t
	l := 0
	r := t
	while l <= r {

		m := (l+r) / 2

	}
}
func day6_part1() {

	//input := [][]int64{{7, 9}, {15, 40}, {30, 200}}
	input2 := [][]int64{{46, 358}, {68, 1054}, {98, 1807}, {66, 1080}}

	result := int64(1)
	for _, pair := range input2 {
		result *= f(pair[0], pair[1])
	}
	log.Println(result)
}

func f(time int64, d int64) int64 {

	var speed, ways int64
	for charge := int64(0); charge <= time; charge++ {

		timeMoving := time - charge
		distance := timeMoving * speed

		if distance > d {
			ways += 1

			fmt.Printf("charge %d ms, race for %d ms, at speed %d, for distance %d\n",
				charge,
				timeMoving,
				speed,
				distance,
			)
		}

		speed++
	}
	return ways
}
