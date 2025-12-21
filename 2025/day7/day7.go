package day7

import (
	"advent-of-code/helpers"
	"fmt"
	"strings"
)

const startManifold = "S"
const emptySpace = '.'
const splitter = '^'

type Bean struct {
	row, col int
	active   bool
}

func countTachyonBeanSplits(lines []string) int {
	var splitCount int
	var beans []Bean

	startIndex := strings.Index(lines[0], startManifold)
	startBean := Bean{row: 1, col: startIndex, active: true}
	beans = append(beans, startBean)

	for index, line := range lines {
		if index == 0 {
			continue
		}

		for i := range beans {
			if !beans[i].active {
				continue
			}

			if line[beans[i].col] == emptySpace {
				beans[i].row = index
			} else if line[beans[i].col] == splitter {
				beans[i].active = false
				leftBean := Bean{active: true, row: index, col: beans[i].col - 1}
				rightBean := Bean{active: true, row: index, col: beans[i].col + 1}

				var leftPositionOccupied bool
				var rightPositionOccupied bool
				for _, search := range beans {
					if search.active && leftBean.col == search.col {
						leftPositionOccupied = true
					}
					if search.active && rightBean.col == search.col {
						rightPositionOccupied = true
					}
				}

				if !leftPositionOccupied && leftBean.col >= 0 {
					beans = append(beans, leftBean)
				}
				if !rightPositionOccupied && rightBean.col < len(line) {
					beans = append(beans, rightBean)
				}

				splitCount++
			}
		}
	}

	return splitCount
}

func Run() {
	lines := helpers.ReadInput("2025/day7/input.txt")

	part1 := countTachyonBeanSplits(lines)
	fmt.Println("Part 1", part1) // 1587
}
