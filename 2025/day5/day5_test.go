package day5

import (
	"slices"
	"testing"
)

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

func TestPart2(test *testing.T) {
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
		result := countFreshIngredientsRanges(lines)
		expected := 14
		if result != expected {
			test.Errorf("result = %d; expected %d", result, expected)
		}
	})
}

func TestMergeRanges(t *testing.T) {
	tests := []struct {
		name     string
		input    []Range
		expected []Range
	}{
		{
			name:     "example from problem",
			input:    []Range{{3, 5}, {10, 14}, {16, 20}, {12, 18}},
			expected: []Range{{3, 5}, {10, 20}},
		},
		{
			name:     "no overlap",
			input:    []Range{{1, 3}, {5, 7}, {9, 11}},
			expected: []Range{{1, 3}, {5, 7}, {9, 11}},
		},
		{
			name:     "complete overlap",
			input:    []Range{{1, 10}, {2, 5}, {3, 8}},
			expected: []Range{{1, 10}},
		},
		{
			name:     "adjacent ranges (touching)",
			input:    []Range{{1, 3}, {4, 6}},
			expected: []Range{{1, 3}, {4, 6}}, // ou merge? Depende da definição
		},
	}
	for _, _case := range tests {
		t.Run(_case.name, func(test *testing.T) {
			result := mergeRanges(_case.input)
			if !slices.Equal(result, _case.expected) {
				t.Errorf("%s: result = %v, expected %v", _case.name, result, _case.expected)
			}
		})
	}
}
