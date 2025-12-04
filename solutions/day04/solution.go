package day04

import (
	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

const roll string = "@"
const empty string = "."

var directions [8][2]int

func init() {
	directions = [8][2]int{
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

			if isRemovable(grid, i, j) {
				result++
			}
		}
	}

	return result, nil
}

func Part2(lines []string) (int, error) {
	// Build grid
	grid := kit.AsGrid(lines)

	// Closure for checking if we have any removable rolls left in the grid
	hasRemovable := func(grid [][]string) (bool, int, int) {
		for i := range grid {
			for j := range grid[i] {
				// Skip empty cells
				if grid[i][j] == empty {
					continue
				}

				if isRemovable(grid, i, j) {
					return true, i, j
				}
			}
		}

		return false, -1, -1
	}

	result := 0
	for {
		ok, y, x := hasRemovable(grid)
		if ok {
			removeRoll(grid, y, x)
			result++
		} else {
			break
		}
	}

	return result, nil
}

// Check number of adjacent paper rolls.
// Return true if number is below 4, else false.
func isRemovable(grid [][]string, y, x int) bool {
	adjacents := 0
	for _, d := range directions {
		// Calculate adjacent position
		dy := y + d[0]
		dx := x + d[1]

		// Checking y-axis bounds
		if dy < 0 || dy >= len(grid) {
			continue
		}

		// Checking x-axis bounds
		if dx < 0 || dx >= len(grid[y]) {
			continue
		}

		// Adjacent is a paper roll
		if string(grid[dy][dx]) == roll {
			adjacents++

			// Stop if we already have too many
			if adjacents > 3 {
				return false
			}
		}
	}

	return true
}

// "Remove" a paper roll from the grid.
func removeRoll(grid [][]string, y, x int) {
	if string(grid[y][x]) == roll {
		grid[y][x] = empty
	}
}
