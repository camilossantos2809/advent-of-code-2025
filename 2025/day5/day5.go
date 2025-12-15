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

func parseRange(line string) (Range, error) {
	strValues := strings.Split(line, "-")
	start, err := strconv.Atoi(strValues[0])
	if err != nil {
		return Range{}, fmt.Errorf("failed to convert start value %s", strValues[0])
	}

	end, err := strconv.Atoi(strValues[1])
	if err != nil {
		return Range{}, fmt.Errorf("failed to convert end value %s", strValues[0])
	}

	return Range{start, end}, nil
}

func parseCode(line string) (int, error) {
	code, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("Failed to convert end value %s", line)
	}

	return code, nil
}

func parseRanges(lines []string) ([]Range, []int, error) {
	var ranges []Range
	var codes []int

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			parsed, err := parseRange(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse range")
			}
			ranges = append(ranges, parsed)
		} else {
			code, err := parseCode(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse code %s", line)
			}
			codes = append(codes, code)
		}
	}

	return ranges, codes, nil
}

func countFreshIngredients(lines []string) int {
	ranges, codes, err := parseRanges(lines)
	if err != nil {
		log.Fatal("Failed to parse ranges")
	}

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
