package day07

import (
	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

const (
	start    rune = 'S'
	splitter rune = '^'
	done     rune = 'X'
	beam     rune = '|'
)

func init() {
	puzzle.Register(7, 1, Part1)
	puzzle.Register(7, 2, Part2)
}

func Part1(lines []string) (int, error) {
	grid := kit.AsGrid(lines)

	// Find start point
	sy, sx := 0, 0
	for i, val := range grid[0] {
		if val == start {
			sx = i
			break
		}
	}

	return hitSplitter(grid, sy, sx), nil
}

func Part2(lines []string) (int, error) {
	return 0, nil
}

func hitSplitter(grid [][]rune, sy, sx int) int {
	// Follow beam downwards until end of field or splitter
	// Do not follow if cell has already been "beamed"
	for sy < len(grid)-1 && grid[sy][sx] != splitter && grid[sy][sx] != beam {
		grid[sy][sx] = beam
		sy++
	}

	// Splitter hit!
	splittersHit := 0
	if grid[sy][sx] == splitter {
		grid[sy][sx] = done
		splittersHit++

		// Split beam left
		splittersHit += hitSplitter(grid, sy, sx-1)
		// Split beam right
		splittersHit += hitSplitter(grid, sy, sx+1)
	}

	return splittersHit
}
