package day3

import (
	"advent-of-code/helpers"
	"fmt"
	"log"
	"strconv"
)

type Joltage struct {
	index    int
	value    int
	valueStr string
}

func findMaxNumbers(line string) (first, second Joltage) {
	var previousFirst Joltage
	lineLength := len(line)
	for charIndex := 0; charIndex < lineLength; charIndex++ {
		digit := int(line[charIndex] - '0')
		if first.value < digit {
			previousFirst = first
			first = Joltage{value: digit, valueStr: string(line[charIndex]), index: charIndex}
			if charIndex == lineLength-1 {
				second = previousFirst
			} else {
				second = Joltage{}
			}
			continue
		}

		if second.value < digit {
			second = Joltage{value: digit, valueStr: string(line[charIndex]), index: charIndex}
		}
	}

	return
}

func sumLargestJoltages(lines []string) int {
	var sum int
	for _, line := range lines {
		first, second := findMaxNumbers(line)

		var concatValues string
		if first.index == len(line)-1 {
			concatValues = second.valueStr + first.valueStr
		} else {
			concatValues = first.valueStr + second.valueStr
		}
		finalValue, err := strconv.Atoi(concatValues)
		if err != nil {
			log.Fatal("Failed to convert concatenated numbers")
		}
		sum += finalValue
	}

	return sum
}

// OBS: Uso do algoritmo Monotonic Stack
func findMaxKDigits(line string, k int) string {
	toRemove := len(line) - k
	// Apesar de definido o tamanho, quando exceder o tamanho, o go ajusta automaticamente
	stack := make([]byte, 0, k)

	for i := 0; i < len(line); i++ {
		// Quando o novo caractere é maior que o último item da stack, significa que esses valores devem ser substituídos.
		for len(stack) > 0 && stack[len(stack)-1] < line[i] && toRemove > 0 {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, line[i])
	}

	// No final, caso quantidade a remover ainda não esteja zerada, irá remover a quantidade ao final
	return string(stack[:len(stack)-toRemove])
}

func sumJoltagesPart2(lines []string) int {
	var sum int
	for _, line := range lines {
		maxDigits := findMaxKDigits(line, 12)
		maxDigitsNumber, err := strconv.Atoi(maxDigits)
		if err != nil {
			log.Fatal("Failed to convert maxDigits")
		}
		sum += maxDigitsNumber
	}

	return sum
}

func Run() {
	lines := helpers.ReadInput("2025/day3/input.txt")

	part1 := sumLargestJoltages(lines)
	fmt.Println("Part 1", part1) // 17107

	part2 := sumJoltagesPart2(lines)
	fmt.Println("Part 2", part2) // 169349762274117
}
