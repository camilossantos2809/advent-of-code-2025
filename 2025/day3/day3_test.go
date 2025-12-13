package day3

import "testing"

func TestPart1(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111"}
		result := sumLargestJoltages(lines)
		expected := 357 // 98 + 89 + 78 + 92 = 357
		if result != expected {
			test.Errorf("sumInvalidIDs = %d; expected %d", result, expected)
		}
	})
}

func TestFindMaxNumbers(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		tests := []struct {
			name        string
			line        string
			firstIndex  int
			secondIndex int
			firstValue  int
			secondValue int
			firstStr    string
			secondStr   string
		}{
			{name: "Primeira linha do exemplo", line: "987654321111111", firstIndex: 0, secondIndex: 1, firstValue: 9, secondValue: 8, firstStr: "9", secondStr: "8"},
			{name: "Segunda linha do exemplo", line: "811111111111119", firstIndex: 14, secondIndex: 0, firstValue: 9, secondValue: 8, firstStr: "9", secondStr: "8"},
			{name: "Terceira linha do exemplo", line: "234234234234278", firstIndex: 14, secondIndex: 13, firstValue: 8, secondValue: 7, firstStr: "8", secondStr: "7"},
			{name: "Quarta linha do exemplo", line: "818181911112111", firstIndex: 6, secondIndex: 11, firstValue: 9, secondValue: 2, firstStr: "9", secondStr: "2"},
		}
		for _, tc := range tests {
			first, second := findMaxNumbers(tc.line)
			if first.index != tc.firstIndex {
				test.Errorf("expected %d; result %d", tc.firstIndex, first.index)
			}
			if second.index != tc.secondIndex {
				test.Errorf("expected %d; result %d", tc.secondIndex, second.index)
			}
			if first.value != tc.firstValue {
				test.Errorf("expected %d; result %d", tc.firstValue, first.value)
			}
			if second.value != tc.secondValue {
				test.Errorf("expected %d; result %d", tc.secondValue, second.value)
			}
		}
	})
}

func TestFindMaxKDigits(test *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
	}{
		{name: "Primeira linha do exemplo", line: "987654321111111", expected: "987654321111"},
		{name: "Segunda linha do exemplo", line: "811111111111119", expected: "811111111119"},
		{name: "Terceira linha do exemplo", line: "234234234234278", expected: "434234234278"},
		{name: "Quarta linha do exemplo", line: "818181911112111", expected: "888911112111"},
	}
	for _, tc := range tests {
		result := findMaxKDigits(tc.line, 12)
		if result != tc.expected {
			test.Errorf("expected %s; result %s", tc.expected, result)
		}
	}
}
