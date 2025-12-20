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
			"  264 421 92 23",
			"  136 212 97 54",
			"   37 416 89 6 ",
			"4  87  69 61 4 ",
			"+ +   +   +  + ",
		}
		result := calcGrandTotalPart2(lines)
		// 34 + 2564 = 2598
		// 2791 + 9986 = 12777
		// 1269 + 2116 + 424 = 3809
		// 4677 + 6338 + 21 = 11036
		// 4
		// 2598 + 12777 + 3809 + 11036 + 4 = 30224
		expected := 30224
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}
