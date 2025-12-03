package main

import "testing"

func TestPart1(test *testing.T) {
	rotations := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	current, zeroCount := countZero(rotations)
	test.Run("Deve retornar 'current' com o valor final 32", func(test *testing.T) {
		if current != 32 {
			test.Errorf("current = %d; expected 32", current)
		}
	})
	test.Run("Deve retornar a contagem de zeros com o valor 3", func(test *testing.T) {
		if zeroCount != 3 {
			test.Errorf("zero_count = %d; expected 3", zeroCount)
		}
	})
}
