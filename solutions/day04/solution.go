package day04

import (
	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

const (
	roll  rune = '@'
	empty rune = '.'
)

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
	result := 0

	for {
		var toRemove [][2]int

		// Look for removable paper rolls
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == roll && isRemovable(grid, y, x) {
					toRemove = append(toRemove, [2]int{y, x})
				}
			}
		}

		// No more items found to remove
		if len(toRemove) == 0 {
			break
		}

		for _, position := range toRemove {
			grid[position[0]][position[1]] = empty
			result++
		}
	}

	return result, nil
}

// Check number of adjacent paper rolls.
// Return true if number is below 4, else false.
func isRemovable(grid [][]rune, y, x int) bool {
	adjacents := 0
	for _, d := range directions {
		// Calculate adjacent position
		dy, dx := y+d[0], x+d[1]

		// Checking y-axis bounds
		if dy < 0 || dy >= len(grid) {
			continue
		}

		// Checking x-axis bounds
		if dx < 0 || dx >= len(grid[y]) {
			continue
		}

		// Adjacent is a paper roll
		if grid[dy][dx] == roll {
			adjacents++

			// Stop if we already have too many
			if adjacents > 3 {
				return false
			}
		}
	}

	return true
}
