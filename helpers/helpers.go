package helpers

import (
	"bufio"
	"log"
	"os"
)

// Read input from a file and return a slice of strings
//
// Example:
//
//	readInput("2025/day1/input.txt")
func ReadInput(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}
