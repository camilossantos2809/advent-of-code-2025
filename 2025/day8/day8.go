package day8

import (
	"advent-of-code/helpers"
	"fmt"
)

func multiply3LargestCircuits(lines []string) int {
	return len(lines)
}

func Run() {
	lines := helpers.ReadInput("2025/day8/input.txt")

	part1 := multiply3LargestCircuits(lines)
	fmt.Println("Part 1", part1)
}
