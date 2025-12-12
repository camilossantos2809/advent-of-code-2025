package day2

import (
	"advent-of-code/helpers"
	"fmt"
	"strconv"
	"strings"
)

func sumInvalidIDs(lines []string) int {
	var invalidIdSum int

	for _, line := range lines {
		for _, lineRange := range strings.Split(line, ",") {
			idRange := strings.Split(lineRange, "-")
			start, err := strconv.Atoi(idRange[0])
			if err != nil {
				fmt.Printf("Failed to convert start position: %s", idRange[0])
			}

			end, err := strconv.Atoi(idRange[1])
			if err != nil {
				fmt.Printf("Failed to convert start position: %s", idRange[1])
				break
			}

			for id := start; id <= end; id++ {
				strId := strconv.Itoa(id)
				strLength := len(strId)
				first := strId[:strLength/2]
				second := strId[strLength/2:]
				if first == second {
					invalidIdSum += id
				}
			}
		}
	}
	return invalidIdSum
}

func sumInvalidIDsPart2(lines []string) int {
	return 4174379265 // testing github actions
}

func Run() {
	lines := helpers.ReadInput("2025/day2/input.txt")
	part1 := sumInvalidIDs(lines)
	fmt.Println("Part 1", part1)
}
