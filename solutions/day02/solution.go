package day02

import (
	"strconv"
	"strings"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

const DialSize = 100
const StartPosition = 50

func init() {
	puzzle.Register(2, 1, Part1)
	puzzle.Register(2, 2, Part2)
}

func Part1(lines []string) (int, error) {
	firstLine := lines[0]
	ranges := strings.Split(firstLine, ",")
	result := 0

	for _, currentRange := range ranges {
		rangeData := strings.Split(currentRange, "-")
		firstId, _ := strconv.Atoi(rangeData[0])
		lastId, _ := strconv.Atoi(rangeData[1])

		for i := firstId; i <= lastId; i++ {
			if !validateId(i) {
				result += i
			}
		}
	}

	return result, nil
}

func Part2(lines []string) (int, error) {
	firstLine := lines[0]
	ranges := strings.Split(firstLine, ",")
	result := 0

	for _, currentRange := range ranges {
		rangeData := strings.Split(currentRange, "-")
		firstId, _ := strconv.Atoi(rangeData[0])
		lastId, _ := strconv.Atoi(rangeData[1])

		for i := firstId; i <= lastId; i++ {
			if !validateIdAdvanced(i) {
				result += i
			}
		}
	}

	return result, nil
}

// Check if a given integer consists of the same pattern exactly twice
func validateId(id int) bool {
	asStr := strconv.Itoa(id)
	if len(asStr)%2 != 0 {
		return true
	}

	middle := len(asStr) / 2
	firstHalf := asStr[:middle]
	secondHalf := asStr[middle:]

	return firstHalf != secondHalf
}

// Check if a given integer consists of repeating patterns (twice or more)
// This could also be used for Part 1.
func validateIdAdvanced(id int) bool {
	idAsStr := strconv.Itoa(id)
	concated := idAsStr + idAsStr

	return !strings.Contains(concated[1:len(concated)-1], idAsStr)
}
