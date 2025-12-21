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
		result := countTachyonBeanSplits(lines)
		expected := 21
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}

func TestBeansConvergemParaMesmoSplitter(test *testing.T) {
	test.Run("Beans de splits diferentes convergem", func(test *testing.T) {
		lines := []string{
			"...S...",
			".......",
			"...^...",
			".......",
			"..^.^..",
			".......",
			"...^...",
		}
		result := countTachyonBeanSplits(lines)
		expected := 4

		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
