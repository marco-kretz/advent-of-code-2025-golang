package day04

import (
	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

const roll string = "@"

var directions [][]int

func init() {
	directions = [][]int{
		{-1, 0},  // Top
		{-1, 1},  // Top Right
		{0, 1},   // Right
		{1, 1},   // Bottom Right
		{1, 0},   // Bottom
		{1, -1},  // Bottom left
		{0, -1},  // Left
		{-1, -1}, // Top Left
	}

	puzzle.Register(4, 1, Part1)
	puzzle.Register(4, 2, Part2)
}

func Part1(lines []string) (int, error) {
	// Build grid
	grid := kit.AsGrid(lines)

	result := 0
	for i := range grid {
		for j := range grid[i] {
			// Skip empty cells
			if grid[i][j] != roll {
				continue
			}

			adjacents := 0
			for _, d := range directions {
				// Calculate adjacent position
				y := i + d[0]
				x := j + d[1]

				// Checking y-axis bounds
				if y < 0 || y >= len(grid) {
					continue
				}

				// Checking x-axis bounds
				if x < 0 || x >= len(grid[i]) {
					continue
				}

				// Adjacent is a paper roll
				if string(grid[y][x]) == roll {
					adjacents++

					// Stop if we already have too many
					if adjacents > 3 {
						break
					}
				}
			}

			if adjacents < 4 {
				result++
			}
		}
	}

	return result, nil

}

func Part2(lines []string) (int, error) {
	return 0, nil
}
