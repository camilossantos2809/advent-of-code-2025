package day7

import "testing"

func TestPart1(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}
		result := countTachyonBeamSplits(lines)
		expected := 21
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}

func TestBeamsConvergemParaMesmoSplitter(test *testing.T) {
	test.Run("Beams de splits diferentes convergem", func(test *testing.T) {
		lines := []string{
			"...S...",
			".......",
			"...^...",
			".......",
			"..^.^..",
			".......",
			"...^...",
		}
		result := countTachyonBeamSplits(lines)
		expected := 4

		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}

func TestPart2(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}
		result := countTachyonBeamTimelines(lines)
		expected := 40
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
