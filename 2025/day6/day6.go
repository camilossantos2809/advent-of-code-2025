package day6

import (
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

func calcGrandTotal(lines []string) int {
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

func Run() {}
