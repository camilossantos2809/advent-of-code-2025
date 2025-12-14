package day4

import "testing"

func TestPart1(test *testing.T) {
	test.Run("Deve retornar o valor correto para o exemplo fornecido no site", func(test *testing.T) {
		lines := []string{
			"..@@.@@@@.",
			"@@@.@.@.@@",
			"@@@@@.@.@@",
			"@.@@@@..@.",
			"@@.@@@@.@@",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"@.@@@.@@@@",
			".@@@@@@@@.",
			"@.@.@@@.@.",
		}
		result := Grid{}.newGrid(lines).countAccessibleRolls()
		expected := 13
		if result != expected {
			test.Errorf("sumInvalidIDs = %d; expected %d", result, expected)
		}
	})
}

func TestGrid_CountNeighborRolls(t *testing.T) {
	tests := []struct {
		name     string
		grid     Grid
		row      int
		col      int
		expected int
	}{
		{
			name: "center position surrounded by 8 rolls",
			grid: Grid{}.newGrid([]string{
				"@@@",
				"@.@",
				"@@@",
			}),
			row:      1,
			col:      1,
			expected: 8,
		},
		{
			name: "center position with no neighbors",
			grid: Grid{}.newGrid([]string{
				"...",
				".@.",
				"...",
			}),
			row:      1,
			col:      1,
			expected: 0,
		},
		{
			name: "top-left corner (3 possible neighbors, 2 are rolls)",
			grid: Grid{}.newGrid([]string{
				"@..",
				".@.",
				"...",
			}),
			row:      0,
			col:      0,
			expected: 1, // only south-east
		},
		{
			name: "top-right corner (3 possible neighbors, all rolls)",
			grid: Grid{}.newGrid([]string{
				"..@",
				"@@@",
				"...",
			}),
			row:      0,
			col:      2,
			expected: 2, // west and south-west
		},
		{
			name: "bottom-left corner",
			grid: Grid{}.newGrid([]string{
				"...",
				"@..",
				"@@.",
			}),
			row:      2,
			col:      0,
			expected: 2, // north and north-east
		},
		{
			name: "bottom-right corner",
			grid: Grid{}.newGrid([]string{
				"...",
				"..@",
				".@@",
			}),
			row:      2,
			col:      2,
			expected: 2, // north and west
		},
		{
			name: "top edge (5 possible neighbors)",
			grid: Grid{}.newGrid([]string{
				".@.",
				"@@@",
				"...",
			}),
			row:      0,
			col:      1,
			expected: 3, // west, east, south-west (south-east seria 4)
		},
		{
			name: "left edge (5 possible neighbors)",
			grid: Grid{}.newGrid([]string{
				"@..",
				"@@.",
				"@..",
			}),
			row:      1,
			col:      0,
			expected: 3, // north, south, north-east
		},
		{
			name: "exactly 3 neighbors (should be accessible)",
			grid: Grid{}.newGrid([]string{
				"@@@",
				"@..",
				"...",
			}),
			row:      0,
			col:      1,
			expected: 3,
		},
		{
			name: "exactly 4 neighbors (should NOT be accessible)",
			grid: Grid{}.newGrid([]string{
				"@@@",
				"@@.",
				"...",
			}),
			row:      0,
			col:      1,
			expected: 4,
		},
		{
			name: "single cell grid",
			grid: Grid{}.newGrid([]string{
				"@",
			}),
			row:      0,
			col:      0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.grid.countNeighborRolls(tt.row, tt.col)
			if result != tt.expected {
				t.Errorf("countNeighborRolls(%d, %d) = %d, expected %d",
					tt.row, tt.col, result, tt.expected)
			}
		})
	}
}

func TestGrid_CountAccessibleRolls(t *testing.T) {
	tests := []struct {
		name     string
		grid     Grid
		expected int
	}{
		{
			name: "example from problem statement",
			grid: Grid{}.newGrid([]string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			}),
			expected: 13,
		},
		{
			name: "empty grid",
			grid: Grid{}.newGrid([]string{
				".....",
				".....",
				".....",
			}),
			expected: 0,
		},
		{
			name: "all isolated rolls (all accessible)",
			grid: Grid{}.newGrid([]string{
				"@.@.@",
				".....",
				"@.@.@",
			}),
			expected: 6,
		},
		{
			name: "single roll (accessible)",
			grid: Grid{}.newGrid([]string{
				"@",
			}),
			expected: 1,
		},
		{
			name: "dense cluster (borders accessible)",
			grid: Grid{}.newGrid([]string{
				"@@@@@",
				"@@@@@",
				"@@@@@",
				"@@@@@",
			}),
			expected: 4,
		},
		{
			name: "2x2 grid all rolls",
			grid: Grid{}.newGrid([]string{
				"@@",
				"@@",
			}),
			expected: 4, // each has 3 neighbors
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.grid.countAccessibleRolls()
			if result != tt.expected {
				t.Errorf("countAccessibleRolls() = %d, expected %d", result, tt.expected)
			}
		})
	}
}
