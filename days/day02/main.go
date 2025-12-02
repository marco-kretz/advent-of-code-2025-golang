package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
)

const DialSize = 100
const StartPosition = 50

func main() {
	lines := kit.ReadFile("inputs/day02.txt")

	// Part 1
	startOne := time.Now()
	resultOne := Part1(lines)
	durationOne := time.Since(startOne)
	fmt.Println("Part 1:", resultOne, "in", durationOne)

	startTwo := time.Now()
	resultTwo := Part2(lines)
	durationTwo := time.Since(startTwo)
	fmt.Println("Part 2:", resultTwo, "in", durationTwo)
}

func Part1(lines []string) int {
	firstLine := lines[0]
	ranges := strings.Split(firstLine, ",")
	result := 0

	for _, currentRange := range ranges {
		rangeData := strings.Split(currentRange, "-")
		firstId, _ := strconv.Atoi(rangeData[0])
		lastId, _ := strconv.Atoi(rangeData[1])

		for i := firstId; i <= lastId; i++ {
			isValid := validateId(i)
			if !isValid {
				result += i
			}
		}
	}

	return result
}

func Part2(lines []string) int {
	return 0
}

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
