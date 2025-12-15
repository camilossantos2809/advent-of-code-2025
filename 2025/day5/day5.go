package day5

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func parseRanges(lines []string) (ranges []Range, codes []int) {
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			strValues := strings.Split(line, "-")
			start, err := strconv.Atoi(strValues[0])
			if err != nil {
				log.Fatalf("Failed to convert start value %s", strValues[0])
			}
			end, err := strconv.Atoi(strValues[1])
			if err != nil {
				log.Fatalf("Failed to convert end value %s", strValues[0])
			}

			ranges = append(ranges, Range{start, end})
		} else {
			code, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Failed to convert end value %s", line)
			}
			codes = append(codes, code)
		}
	}
	return ranges, codes
}

func countFreshIngredients(lines []string) int {
	ranges, codes := parseRanges(lines)
	var freshCount int
	for _, code := range codes {
		for _, _range := range ranges {
			if code >= _range.start && code <= _range.end {
				freshCount++
				break
			}
		}
	}
	return freshCount
}

func Run() {
	lines := helpers.ReadInput("2025/day5/input.txt")

	part1 := countFreshIngredients(lines)
	fmt.Println("Part 1", part1) // 733
}
