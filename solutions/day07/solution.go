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

type point struct {
	x, y int
}

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

	// Count splitters hit for each row
	splitterMap := make(map[int][]int)
	hitSplitter(grid, splitterMap, sy, sx)

	// Add row's splitter counts
	hits := 0
	for _, v := range splitterMap {
		hits += len(v)
	}

	return hits, nil
}

func Part2(lines []string) (int, error) {
	grid := kit.AsGrid(lines)

	// Find start point
	sy, sx := 0, 0
	for i, val := range grid[0] {
		if val == start {
			sx = i
			break
		}
	}

	// Cache
	cache := make(map[point]int)
	totalPaths := findPaths(grid, cache, sy, sx)

	return totalPaths, nil
}

func hitSplitter(grid [][]rune, splitterMap map[int][]int, sy, sx int) {
	// Follow beam downwards until end of field or splitter
	// Do not follow if cell has already been "beamed"
	for sy < len(grid)-1 && grid[sy][sx] != splitter && grid[sy][sx] != beam {
		grid[sy][sx] = beam
		sy++
	}

	// Splitter hit!
	if grid[sy][sx] == splitter {
		grid[sy][sx] = done
		if _, ok := splitterMap[sy]; ok {
			splitterMap[sy] = append(splitterMap[sy], sx)
		} else {
			splitterMap[sy] = []int{sy}
		}

		// Split beam left
		hitSplitter(grid, splitterMap, sy, sx-1)
		// Split beam right
		hitSplitter(grid, splitterMap, sy, sx+1)
	}
}

func findPaths(grid [][]rune, cache map[point]int, y, x int) int {
	// Point already processed?
	if val, ok := cache[point{x, y}]; ok {
		return val
	}

	// Follow beam downwards until abyss
	currY := y
	for currY < len(grid) {
		// Out of bounds?
		if x < 0 || x >= len(grid[0]) {
			return 1
		}

		// Hit a splitter?
		cell := grid[currY][x]
		if cell == splitter {
			// Go left
			leftCount := findPaths(grid, cache, currY, x-1)
			// Go right
			rightCount := findPaths(grid, cache, currY, x+1)

			// Sum up paths found
			total := leftCount + rightCount

			// Set point as processed
			cache[point{x, y}] = total

			return total
		}

		currY++
	}

	return 1
}
