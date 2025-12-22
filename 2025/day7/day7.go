package day7

import (
	"advent-of-code/helpers"
	"fmt"
	"strings"
)

const startManifold = "S"
const emptySpace = '.'
const splitter = '^'

type Beam struct {
	row, col int
	active   bool
}

func countTachyonBeamSplits(lines []string) int {
	var splitCount int
	var beams []Beam

	startIndex := strings.Index(lines[0], startManifold)
	startBeam := Beam{row: 1, col: startIndex, active: true}
	beams = append(beams, startBeam)

	for index, line := range lines {
		if index == 0 {
			continue
		}

		for i := range beams {
			if !beams[i].active {
				continue
			}

			if line[beams[i].col] == emptySpace {
				beams[i].row = index
			} else if line[beams[i].col] == splitter {
				beams[i].active = false
				leftBeam := Beam{active: true, row: index, col: beams[i].col - 1}
				rightBeam := Beam{active: true, row: index, col: beams[i].col + 1}

				var leftPositionOccupied bool
				var rightPositionOccupied bool
				for _, search := range beams {
					if search.active && leftBeam.col == search.col {
						leftPositionOccupied = true
					}
					if search.active && rightBeam.col == search.col {
						rightPositionOccupied = true
					}
				}

				if !leftPositionOccupied && leftBeam.col >= 0 {
					beams = append(beams, leftBeam)
				}
				if !rightPositionOccupied && rightBeam.col < len(line) {
					beams = append(beams, rightBeam)
				}

				splitCount++
			}
		}
	}

	return splitCount
}

func countTachyonBeamTimelines(lines []string) int {
	counts := make(map[int]int)

	startIndex := strings.Index(lines[0], startManifold)
	counts[startIndex] = 1

	for index, line := range lines {
		if index == 0 {

			continue
		}
		newCounts := make(map[int]int)
		for col, count := range counts {
			if line[col] == splitter {
				newCounts[col-1] += count
				newCounts[col+1] += count
			} else {
				newCounts[col] += count
			}
		}
		counts = newCounts
	}

	var total int
	for _, count := range counts {
		total += count
	}

	return total
}

func Run() {
	lines := helpers.ReadInput("2025/day7/input.txt")

	part1 := countTachyonBeamSplits(lines)
	fmt.Println("Part 1", part1) // 1587

	part2 := countTachyonBeamTimelines(lines)
	fmt.Println("Part 2", part2) // 5748679033029
}
