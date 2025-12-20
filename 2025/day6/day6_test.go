package day6

import "testing"

func TestPart1(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
			"*   +   *   +  ",
		}
		result := calcGrandTotal(lines)
		// 33210 + 490 + 4243455 + 401
		expected := 4277556
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
