package main

import (
	"bufio"
	"os"
)

func main() {
	day1()
}

// return file contents
func readFileContents(filePath string) []string {
	readFile, _ := os.Open(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var inputlines []string

	// load up all the data into an array
	for fileScanner.Scan() {
		linestring := fileScanner.Text()
		inputlines = append(inputlines, linestring)
	}
	return inputlines
}
