package day8

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

func (p *Point) NewPoint(x, y, z string) {
	numx, err := strconv.Atoi(x)
	if err != nil {
		log.Fatalf("Failed to convert x")
	}
	numy, err := strconv.Atoi(y)
	if err != nil {
		log.Fatalf("Failed to convert y")
	}
	numz, err := strconv.Atoi(z)
	if err != nil {
		log.Fatalf("Failed to convert z")
	}

	p.x = numx
	p.y = numy
	p.z = numz
}

type Connection struct {
	firstPoint, secondPoint, distance int
}

func (c *Connection) newConnection(firstPoint, secondPoint Point, firstI, secondI int) {
	c.firstPoint = firstI
	c.secondPoint = secondI

	dx := secondPoint.x - firstPoint.x
	dy := secondPoint.y - firstPoint.y
	dz := secondPoint.z - firstPoint.z
	c.distance = dx*dx + dy*dy + dz*dz
}

type UnionFind struct {
	parents []int
	sizes   []int
}

func (u *UnionFind) init(length int) {
	if len(u.parents) == 0 {
		u.parents = make([]int, length)
	}
	if len(u.sizes) == 0 {
		u.sizes = make([]int, length)
	}
	for i := 0; i < length; i++ {
		u.parents[i] = i
		u.sizes[i] = 1
	}
}

func (u *UnionFind) find(index int) int {
	for u.parents[index] != index {
		index = u.parents[index]
	}
	return index
}

func (u *UnionFind) union(first, second int) {
	firstRoot := u.find(first)
	secondRoot := u.find(second)

	if firstRoot == secondRoot {
		return
	}

	u.parents[secondRoot] = firstRoot
	u.sizes[firstRoot] += u.sizes[secondRoot]
}

func multiply3LargestCircuits(lines []string, limit int) int {
	var points []Point
	for _, line := range lines {
		split := strings.Split(line, ",")
		point := Point{}
		point.NewPoint(split[0], split[1], split[2])
		points = append(points, point)
	}

	var connections []Connection
	for i, parentPoint := range points {
		for j := i + 1; j < len(points); j++ {
			connection := Connection{}
			connection.newConnection(parentPoint, points[j], i, j)
			connections = append(connections, connection)
		}
	}
	slices.SortFunc(connections, func(a, b Connection) int {
		return a.distance - b.distance
	})

	firstConn := connections[:limit]
	union := UnionFind{}
	union.init(len(points))
	for _, connection := range firstConn {
		union.union(connection.firstPoint, connection.secondPoint)
	}

	var circuitSizes []int
	for i, parent := range union.parents {
		if parent == i {
			circuitSizes = append(circuitSizes, union.sizes[i])
		}
	}
	slices.SortFunc(circuitSizes, func(a, b int) int {
		return b - a
	})

	var latest int
	var largest []int
	result := 1
	for _, circuit := range circuitSizes {
		if circuit == latest {
			continue
		}
		largest = append(largest, circuit)
		latest = circuit
		result *= circuit

		if len(largest) == 3 {
			break
		}
	}

	return result
}

func Run() {
	lines := helpers.ReadInput("2025/day8/input.txt")

	part1 := multiply3LargestCircuits(lines, 1000)
	fmt.Println("Part 1", part1) // 97384
}
