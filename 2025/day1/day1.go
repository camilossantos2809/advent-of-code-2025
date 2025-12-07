package day1

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
	"math"

	"strconv"
)

const startDial = 50

type CalculateDialResult struct {
	newPosition int
	rotation    int
	diff        int
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
		newPosition: ((diff)%100 + 100) % 100,
		rotation:    rotationInt,
		diff:        diff,
	}
}

func countZerosAtTheEndOfRotation(rotations []string) (int, int) {
	current := startDial
	zeroCount := 0
	for _, rotation := range rotations {
		result := calculateDial(current, rotation)
		current = result.newPosition
		if result.newPosition == 0 {
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
		var rotationZeroCount float64
		if result.diff > current {
			rotationZeroCount = math.Floor(float64(result.diff)/100) - math.Floor(float64(current)/100)
		} else {
			rotationZeroCount = math.Floor(float64((result.diff-1))/100) - math.Floor(float64((current-1))/100)
		}
		rotationZeroCountInt := int(math.Abs(rotationZeroCount))
		zeroCount += rotationZeroCountInt
		current = result.newPosition
	}

	return zeroCount
}

func Run() {
	fileLines := helpers.ReadInput("2025/day1/input.txt")

	_, zeroCount := countZerosAtTheEndOfRotation(fileLines)
	fmt.Println("Part 1", zeroCount)

	zeroCount2 := countZerosForAnyClick(fileLines)
	fmt.Println("Part 2", zeroCount2)
}
