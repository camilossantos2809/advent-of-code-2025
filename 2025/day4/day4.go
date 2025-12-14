package day4

import (
	"advent-of-code/helpers"
	"fmt"
)

type Grid struct {
	lines [][]byte
}

func (grid Grid) newGrid(lines []string) Grid {
	grid.lines = make([][]byte, len(lines))
	for i, line := range lines {
		grid.lines[i] = []byte(line)
	}

	return grid
}

func (grid Grid) isInBounds(row, col int) bool {
	return !(row < 0 || col < 0 || row >= len(grid.lines) || col >= len(grid.lines[0]))
}

func (grid Grid) countNeighborRolls(row, col int) int {
	var count int
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, direction := range directions {
		dRow := row + direction[0]
		dCol := col + direction[1]
		if grid.isInBounds(dRow, dCol) && grid.lines[dRow][dCol] == '@' {
			count++
		}
	}

	return count
}

func (grid Grid) countAccessibleRolls() int {
	var count int
	for row, line := range grid.lines {
		for col, char := range line {
			if char == '@' {
				countNeighbor := grid.countNeighborRolls(row, col)
				if countNeighbor < 4 {
					count++
				}
			}
		}
	}

	return count
}

func Run() {
	lines := helpers.ReadInput("2025/day4/input.txt")

	part1 := Grid{}.newGrid(lines).countAccessibleRolls()
	fmt.Println("Part 1", part1) // 1491

	// part2 := sumJoltagesPart2(lines)
	// fmt.Println("Part 2", part2) // 169349762274117
}
