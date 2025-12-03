package day01

import (
	"math"
	"strconv"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

const DialSize = 100
const StartPosition = 50

func init() {
	puzzle.Register(1, 1, Part1)
	puzzle.Register(1, 2, Part2)
}

func Part1(lines []string) (int, error) {
	currentPosition := StartPosition
	totalZeros := 0

	for _, line := range lines {
		direction := line[:1]
		distance, _ := strconv.Atoi(line[1:])
		segDistance := distance % DialSize

		if direction == "R" {
			currentPosition += segDistance
		}
		if direction == "L" {
			currentPosition -= segDistance
		}

		currentPosition = kit.EuclideanModulo(currentPosition, DialSize)

		if currentPosition == 0 {
			totalZeros++
		}
	}

	return totalZeros, nil
}

func Part2(lines []string) (int, error) {
	currentPosition := StartPosition
	totalZeros := 0

	for _, line := range lines {
		direction := line[:1]
		distance, _ := strconv.Atoi(line[1:])

		// Count "full cycles"
		fullCycles := math.Floor(float64(distance) / DialSize)
		totalZeros += int(fullCycles)

		segDistance := distance % DialSize
		oldPosition := currentPosition
		if direction == "R" {
			currentPosition += segDistance
		}
		if direction == "L" {
			currentPosition -= segDistance
		}

		if (currentPosition < 0 && oldPosition != 0) ||
			currentPosition >= DialSize ||
			currentPosition == 0 {
			totalZeros++
		}

		currentPosition = kit.EuclideanModulo(currentPosition, DialSize)
	}

	return totalZeros, nil
}
