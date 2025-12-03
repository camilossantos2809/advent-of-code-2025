package main

import (
	"log"
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
	zero_count := 0
	for _, rotation := range rotations {
		current = calculateDial(current, rotation)
		if current == 0 {
			zero_count++
		}
	}
	return current, zero_count
}

func main() {
}
