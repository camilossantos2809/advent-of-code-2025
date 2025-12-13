package main

import (
	"advent-of-code/2025/day1"
	"advent-of-code/2025/day2"
	"advent-of-code/2025/day3"
	"fmt"
	"os"
)

var puzzles = map[string]func(){
	"2025day1": day1.Run,
	"2025day2": day2.Run,
	"2025day3": day3.Run,
}

// Exemplo de usos pelo terminal:
//
//	go run cmd/main.go 2025 day1
//	go run cmd/main.go 2025 day2
func main() {
	year := os.Args[1]
	day := os.Args[2]

	puzzle, ok := puzzles[year+day]
	if !ok {
		fmt.Println("Puzzle not found")
		return
	}

	puzzle()
}
