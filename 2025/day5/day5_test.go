package day5

import "testing"

func TestPart1(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
			"",
			"1",
			"5",
			"8",
			"11",
			"17",
			"32",
		}
		result := countFreshIngredients(lines)
		expected := 3
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
