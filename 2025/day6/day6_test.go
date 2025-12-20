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
		result := calcGrandTotalPart1(lines)
		// 33210 + 490 + 4243455 + 401
		expected := 4277556
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}

func TestPart2(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
			"*   +   *   +  ",
		}
		result := calcGrandTotalPart2(lines)
		// The rightmost problem is 4 + 431 + 623 = 1058
		// The second problem from the right is 175 * 581 * 32 = 3253600
		// The third problem from the right is 8 + 248 + 369 = 625
		// Finally, the leftmost problem is 356 * 24 * 1 = 8544
		// Now, the grand total is 1058 + 3253600 + 625 + 8544 = 3263827.
		expected := 3263827
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})

	test.Run("Deve considerar colunas com quantidades de dígitos diferentes", func(test *testing.T) {
		lines := []string{
			"264 421 92 23",
			"136 212 97 54",
			" 37 416 89 6 ",
			"+   +   +  + ",
		}
		result := calcGrandTotalPart2(lines)
		// 34 + 256 = 290
		// 279 + 998 = 1277
		// 126 + 211 + 424 = 761
		// 467 + 633 + 21 = 1121
		// 290 + 1277 + 761 + 1121 = 3449
		expected := 3449
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
