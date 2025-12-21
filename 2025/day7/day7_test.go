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
			"...S...", // bean inicial em col 3
			".......",
			"...^...", // split cria beans em col 2 e col 4
			".......",
			"..^.^..", // bean esquerda(col 2) encontra ^ em col 2
			// bean direita(col 4) encontra ^ em col 4
			".......",
			"...^...", // AMBOS criam beans que convergem aqui (col 3)
		}

		// Linhas 0-1: bean inicial desce
		// Linha 2: split 1 (col 3)
		// Linha 4: split 2 (col 2) e split 3 (col 4)
		// Linha 6: beans convergem para col 3
		//   - Se ambos processarem este ^, conta 2 vezes
		//   - Deveria contar apenas 1 vez (split 4)

		result := countTachyonBeanSplits(lines)
		expected := 4 // 4 splitters unicos

		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
