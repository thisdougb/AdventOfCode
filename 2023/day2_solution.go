package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	filePath = "day2_input.txt"
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	log.Printf("Day 2 p1: Cube Conundrum - sum of game ids: %d", part1())
	log.Printf("Day 2 p2: Cube Conundrum - sum of the power of sets: %d", part2())
}
func part1() int {

	readFile, _ := os.Open(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var gameIdTotal int

	for fileScanner.Scan() {

		linestring := fileScanner.Text()

		// first get the game id from the linestring
		a := strings.Split(linestring, ":")
		gameId := a[0][len("Game "):]

		var redHigh, greenHigh, blueHigh int

		// now extract the cube sets
		cubeSets := strings.Split(a[1], ";")

		for _, cs := range cubeSets {

			// extract each cub
			cubeDescriptions := strings.Split(cs, ",")

			for _, cubeDescription := range cubeDescriptions {

				trimmedDescription := strings.TrimSpace(cubeDescription)

				// now extract the actual number anbd colour values
				cubeValues := strings.Split(trimmedDescription, " ")

				number, err := strconv.Atoi(cubeValues[0])
				if err != nil {
					log.Println(err)
				}
				switch cubeValues[1] {

				case "red":
					if number > redHigh {
						redHigh = number
					}
				case "green":
					if number > greenHigh {
						greenHigh = number
					}
				case "blue":
					if number > blueHigh {
						blueHigh = number
					}
				}
			}
		}
		if redHigh <= maxRed && greenHigh <= maxGreen && blueHigh <= maxBlue {
			id, _ := strconv.Atoi(gameId)
			gameIdTotal += id
		}
	}
	return gameIdTotal
}

func part2() int {

	readFile, _ := os.Open(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var gameTotal int

	for fileScanner.Scan() {

		linestring := fileScanner.Text()

		// first get the game id from the linestring
		a := strings.Split(linestring, ":")

		// now extract the cube sets
		cubeSets := strings.Split(a[1], ";")

		var redHigh, greenHigh, blueHigh int

		for _, cs := range cubeSets {

			// extract each cub
			cubeDescriptions := strings.Split(cs, ",")

			for _, cubeDescription := range cubeDescriptions {

				trimmedDescription := strings.TrimSpace(cubeDescription)

				// now extract the actual number anbd colour values
				cubeValues := strings.Split(trimmedDescription, " ")

				number, err := strconv.Atoi(cubeValues[0])
				if err != nil {
					log.Println(err)
				}
				switch cubeValues[1] {

				case "red":
					if number > redHigh {
						redHigh = number
					}
				case "green":
					if number > greenHigh {
						greenHigh = number
					}
				case "blue":
					if number > blueHigh {
						blueHigh = number
					}
				}
			}
		}

		gameTotal += redHigh * greenHigh * blueHigh
	}
	return gameTotal
}
