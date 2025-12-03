package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateDial(current int, rotation string) int {
	rotationInt, err := strconv.Atoi(rotation[1:])
	if err != nil {
		log.Fatal(err)
	}

	if rotation[0] == 'L' {
		return ((current-rotationInt)%100 + 100) % 100
	}
	return ((current+rotationInt)%100 + 100) % 100
}

func countZero(rotations []string) (int, int) {
	current := 50
	zeroCount := 0
	for _, rotation := range rotations {
		current = calculateDial(current, rotation)
		if current == 0 {
			zeroCount++
		}
	}
	return current, zeroCount
}

func main() {
	readFile, err := os.Open("2025/day1/input.txt")
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
	current, zeroCount := countZero(fileLines)
	fmt.Println(current)
	fmt.Println(zeroCount)
}
