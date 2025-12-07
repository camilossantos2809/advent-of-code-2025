package day1

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
			if result.newPosition != _case.expected {
				test.Errorf("Expected %d, result %d", _case.expected, result.newPosition)
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
	test.Run("Deve contar corretamente quando começar perto do zero e rotacionar muito", func(test *testing.T) {
		// Starting at 50, rotating R150 should cross 0 twice:
		// 50 -> 100 (crosses 0 once) -> 200 (crosses 0 again, lands at 0)
		zeroCount := countZerosForAnyClick([]string{"R150"})
		if zeroCount != 2 {
			test.Errorf("zeroCount = %d; expected 2", zeroCount)
		}
	})
	test.Run("Rotação pequena que não passa pelo zero", func(test *testing.T) {
		// Starting at 50, R30 -> ends at 80, never crosses 0
		zeroCount := countZerosForAnyClick([]string{"R30"})
		if zeroCount != 0 {
			test.Errorf("zeroCount = %d; expected 0", zeroCount)
		}
	})
	test.Run("Começando em 99 e rotacionando R1", func(test *testing.T) {
		// 99 + 1 = 100 -> crosses once, lands at 0
		// Need to set current to 99 first
		zeroCount := countZerosForAnyClick([]string{"R49", "R1"}) // 50->99, then 99->0
		if zeroCount != 1 {
			test.Errorf("zeroCount = %d; expected 1", zeroCount)
		}
	})
	test.Run("Rotação exata de 100 a partir de posição 0", func(test *testing.T) {
		zeroCount := countZerosForAnyClick([]string{"L50", "R100"})
		if zeroCount != 2 {
			test.Errorf("zeroCount = %d; expected 2", zeroCount)
		}
	})
	test.Run("Rotação de 99 que não cruza zero", func(test *testing.T) {
		// Starting at 1, R99 -> ends at 0, crosses once
		zeroCount := countZerosForAnyClick([]string{"L49", "R99"}) // 50->1, then 1->0 (crosses at 100)
		if zeroCount != 1 {
			test.Errorf("zeroCount = %d; expected 1", zeroCount)
		}
	})
	test.Run("Rotação pequena para esquerda que cruza", func(test *testing.T) {
		// Starting at 0, L1 -> crosses once, ends at 99
		zeroCount := countZerosForAnyClick([]string{"L50", "L1"}) // 50->0, then 0->99
		if zeroCount != 1 {
			test.Errorf("zeroCount = %d; expected 1", zeroCount)
		}
	})
	test.Run("Começando em 99 e rotacionando muito para direita", func(test *testing.T) {
		// Starting at 99, R201 goes to position 300
		// Crosses: 100 (1st), 200 (2nd), 300 (3rd) = 3 times
		zeroCount := countZerosForAnyClick([]string{"R49", "R201"}) // 50->99, then 99->300
		if zeroCount != 3 {
			test.Errorf("zeroCount = %d; expected 3", zeroCount)
		}
	})
	test.Run("Começando em 1 e rotacionando L2", func(test *testing.T) {
		// Starting at 1, L2 goes to -1 (wraps to 99)
		// Crosses 0 once going backward
		zeroCount := countZerosForAnyClick([]string{"L49", "L2"}) // 50->1, then 1->99
		if zeroCount != 1 {
			test.Errorf("zeroCount = %d; expected 1", zeroCount)
		}
	})
	test.Run("Debug: position 99, R201", func(test *testing.T) {
		// 99 -> 300 should cross at 100, 200, 300 = 3 times
		zeroCount := countZerosForAnyClick([]string{"R49", "R201"})
		if zeroCount != 3 {
			test.Errorf("zeroCount = %d; expected 3", zeroCount)
		}
	})
	test.Run("Começando próximo ao zero indo para esquerda com rotação grande", func(test *testing.T) {
		// Starting at 5, L205 goes to -200
		// Crosses: 0 (1st), -100/0 (2nd), -200/0 (3rd) = 3 times
		zeroCount := countZerosForAnyClick([]string{"L45", "L205"}) // 50->5, then 5->-200
		if zeroCount != 3 {
			test.Errorf("zeroCount = %d; expected 3", zeroCount)
		}
	})
}
