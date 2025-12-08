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
			}

			for id := start; id <= end; id++ {
				strId := fmt.Sprint(id)
				strLenght := len(strId)
				first := strId[:strLenght/2]
				second := strId[strLenght/2:]
				if first == second {
					invalidIdSum += id
				}
			}
		}
	}
	return invalidIdSum
}

func Run() {
	lines := helpers.ReadInput("2025/day2/input.txt")
	part1 := sumInvalidIDs(lines)
	fmt.Println("Part 1", part1)
}
