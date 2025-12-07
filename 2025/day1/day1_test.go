package main

import "testing"

func TestPart1(test *testing.T) {
	rotations := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	current, zeroCount := countZerosAtTheEndOfRotation(rotations)
	test.Run("Deve retornar 'current' com o valor final 32", func(test *testing.T) {
		if current != 32 {
			test.Errorf("current = %d; expected 32", current)
		}
	})
	test.Run("Deve retornar a contagem de zeros ao final da rotação do dial", func(test *testing.T) {
		if zeroCount != 3 {
			test.Errorf("zeroCount = %d; expected 3", zeroCount)
		}
	})
}

func TestCalculateDial(test *testing.T) {
	cases := []struct {
		initial  int
		rotation string
		expected int
	}{
		{initial: 50, rotation: "L68", expected: 82},
		{initial: 82, rotation: "L30", expected: 52},
		{initial: 52, rotation: "R48", expected: 0},
		{initial: 0, rotation: "L5", expected: 95},
	}
	for _, _case := range cases {
		test.Run("Deve gerar o resultado da rotação correto", func(test *testing.T) {
			result := calculateDial(_case.initial, _case.rotation)
			if result.current != _case.expected {
				test.Errorf("Expected %d, result %d", _case.expected, result.current)
			}
		})
	}
}

func TestPart2(test *testing.T) {

	test.Run("Deve retornar a quantidade de vezes que o dial passa pelo zero", func(test *testing.T) {
		rotations := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		zeroCount := countZerosForAnyClick(rotations)
		if zeroCount != 6 {
			test.Errorf("zeroCount = %d; expected 6", zeroCount)
		}
	})
	test.Run("Deve retornar a quantidade de vezes que o dial passar pelo zero inclusive se mais de uma vez na rotação", func(test *testing.T) {
		zeroCount := countZerosForAnyClick([]string{"R1000"})
		if zeroCount != 10 {
			test.Errorf("zeroCount = %d; expected 10", zeroCount)
		}
	})
	test.Run("Deve retornar a quantidade de vezes que o dial passar pelo zero inclusive se mais de uma vez na rotação para esquerda", func(test *testing.T) {
		zeroCount := countZerosForAnyClick([]string{"L1000"})
		if zeroCount != 10 {
			test.Errorf("zeroCount = %d; expected 10", zeroCount)
		}
	})
	test.Run("Quando houver uma rotação normal e outra que passe pelo zero várias vezes", func(test *testing.T) {
		test.Run("Deve retornar a quantidade de vezes que o dial passar pelo zero", func(test *testing.T) {
			zeroCount := countZerosForAnyClick([]string{"R1000", "L1000", "R100", "L100", "R5", "L5"})
			if zeroCount != 22 {
				test.Errorf("zeroCount = %d; expected 22", zeroCount)
			}
		})
	})
}
