package main

import (
	"bufio"
	"fmt"
	"log"
	"math"

	"os"
	"strconv"
)

const startDial = 50

type CalculateDialResult struct {
	current  int
	rotation int
	diff     int
}

func calculateDial(current int, rotation string) CalculateDialResult {
	rotationInt, err := strconv.Atoi(rotation[1:])
	if err != nil {
		log.Fatal(err)
	}

	var diff int
	if rotation[0] == 'L' {
		diff = current - rotationInt
	} else {
		diff = current + rotationInt
	}

	return CalculateDialResult{
		current:  ((diff)%100 + 100) % 100,
		rotation: rotationInt,
		diff:     diff,
	}
}

func countZerosAtTheEndOfRotation(rotations []string) (int, int) {
	current := startDial
	zeroCount := 0
	for _, rotation := range rotations {
		result := calculateDial(current, rotation)
		current = result.current
		if result.current == 0 {
			zeroCount++
		}
	}
	return current, zeroCount
}

func countZerosForAnyClick(rotations []string) int {
	current := startDial
	zeroCount := 0
	for _, rotation := range rotations {
		result := calculateDial(current, rotation)
		current = result.current

		absDiff := math.Abs(float64(result.rotation))
		if absDiff != 0 {
			division := absDiff / 100
			if division > 1 {
				zeroCount += int(division)
			} else if result.diff < 0 || result.diff > 99 {
				zeroCount++
			}
			fmt.Println(rotation, zeroCount, division)
		}

	}

	return zeroCount
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
	// _, zeroCount := countZerosAtTheEndOfRotation(fileLines)
	// fmt.Println("Part 1", zeroCount)

	zeroCount2 := countZerosForAnyClick(fileLines)
	fmt.Println("Part 2", zeroCount2)
}
