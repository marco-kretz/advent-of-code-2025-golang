package day05

import (
	"strconv"
	"strings"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

func init() {
	puzzle.Register(5, 1, Part1)
	puzzle.Register(5, 2, Part2)
}

func Part1(lines []string) (int, error) {
	var ranges [][2]int
	var fresh []int

	hitEmpty := false
	for _, value := range lines {
		// Found the empty line
		if value == "" {
			hitEmpty = true
			continue
		}

		// Line with a "fresh" ID range
		if !hitEmpty {
			idRange := strings.Split(value, "-")
			if len(idRange) == 2 {
				start, _ := strconv.Atoi(idRange[0])
				end, _ := strconv.Atoi(idRange[1])
				ranges = append(ranges, [2]int{start, end})
			}
			continue
		}

		// Line with an "available" ID
		id, _ := strconv.Atoi(value)
		if inFresh(ranges, id) {
			fresh = append(fresh, id)
		}
	}

	return len(fresh), nil
}

func Part2(lines []string) (int, error) {

	return 0, nil
}

func inFresh(ranges [][2]int, value int) bool {
	for _, r := range ranges {
		if value >= r[0] && value <= r[1] {
			return true
		}
	}

	return false
}
