package day8

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
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

func multiply3LargestCircuits(lines []string) int {
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

	return len(connections)
}

func Run() {
	lines := helpers.ReadInput("2025/day8/input.txt")

	part1 := multiply3LargestCircuits(lines)
	fmt.Println("Part 1", part1)
}
