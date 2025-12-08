package day2

import "testing"

func TestPart1(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
		result := sumInvalidIDs(lines)
		expected := 1227775554
		if result != expected {
			test.Errorf("sumInvalidIDs = %d; expected %d", result, expected)
		}
	})
}

func TestPart2(test *testing.T) {
}
