package day3

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
	"strconv"
)

type Joltage struct {
	index    int
	value    int
	valueStr string
}

func findMaxNumbers(line string) (first Joltage, second Joltage) {
	var previousFirst Joltage
	lineLength := len(line)
	for charIndex := 0; charIndex < lineLength; charIndex++ {
		digit := int(line[charIndex] - '0')
		if first.value < digit {
			previousFirst = first

			first.value = digit
			first.valueStr = string(line[charIndex])
			first.index = charIndex

			if charIndex == lineLength-1 {
				second = previousFirst
			} else {
				second.index = 0
				second.value = 0
				second.valueStr = ""
			}
			continue
		}

		if second.value < digit {
			second.value = digit
			second.valueStr = string(line[charIndex])
			second.index = charIndex
		}
	}

	return first, second
}

func sumLargestJoltages(lines []string) int {
	var sum int
	for _, line := range lines {
		first, second := findMaxNumbers(line)

		var concatValues string
		if first.index == len(line)-1 {
			concatValues = second.valueStr + first.valueStr
		} else {
			concatValues = first.valueStr + second.valueStr
		}
		finalValue, err := strconv.Atoi(concatValues)
		if err != nil {
			log.Fatal("Failed to convert concatenated numbers")
		}
		sum += finalValue
	}

	return sum
}

func Run() {
	lines := helpers.ReadInput("2025/day3/input.txt")

	part1 := sumLargestJoltages(lines)
	fmt.Println("Part 1", part1)

	// part2 := sumInvalidIDsPart2(lines)
	// fmt.Println("Part 2", part2)
}
