package day6

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Problem struct {
	operation string
	values    []int
	total     int
}

func (p *Problem) calcTotal() {
	p.total = 0
	for _, value := range p.values {
		if p.operation == "+" {
			p.total += value
		} else {
			if p.total == 0 {
				p.total = value
			} else {
				p.total *= value
			}
		}
	}
}

func calcGrandTotalPart1(lines []string) int {
	problems := make(map[int]Problem)
	for _, line := range lines {
		lineSlice := strings.Split(line, " ")
		lineSlice = slices.DeleteFunc(lineSlice, func(item string) bool {
			return item == ""
		})
		for i, item := range lineSlice {
			if len(problems[i].values) == 0 {
				num, err := strconv.Atoi(item)
				if err != nil {
					log.Fatalf("Failed to convert line item %s", item)
				}
				numSlice := make([]int, 1)
				numSlice[0] = num
				problems[i] = Problem{values: numSlice}
			} else if item != "+" && item != "*" {
				num, err := strconv.Atoi(item)
				if err != nil {
					log.Fatalf("Failed to convert line item %s", item)
				}
				problem := problems[i]
				problem.values = append(problems[i].values, num)
				problems[i] = problem
			} else {
				problem := problems[i]
				problem.operation = item
				problems[i] = problem
			}
		}
	}

	var total int
	for _, problem := range problems {
		problem.calcTotal()
		total += problem.total
	}
	return total
}

func calcGrandTotalPart2(lines []string) int {
	var problems []Problem
	var values []int
	for i := len(lines[0]) - 1; i >= 0; i-- {

		if (i+1)%4 == 0 {
			problems = append(problems, Problem{values: values, operation: string(lines[3][i+1])})
			values = make([]int, 0)
			continue
		}
		var builder strings.Builder
		builder.WriteByte(lines[0][i])
		builder.WriteByte(lines[1][i])
		builder.WriteByte(lines[2][i])
		numStr := strings.TrimSpace(builder.String())
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("Failed to convert string %s", numStr)
		}
		values = append(values, num)

		if i == 0 {
			problems = append(problems, Problem{values: values, operation: string(lines[3][i+1])})
			values = make([]int, 0)
		}
	}

	var total int
	for _, problem := range problems {
		problem.calcTotal()
		total += problem.total
	}
	return total
}

func Run() {
	lines := helpers.ReadInput("2025/day6/input.txt")

	part1 := calcGrandTotalPart1(lines)
	fmt.Println("Part 1", part1) // 4771265398012

	part2 := calcGrandTotalPart2(lines)
	fmt.Println("Part 1", part2)
}
