package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
)

const DialSize = 100
const StartPosition = 50

func main() {
	lines := kit.ReadFile("inputs/day01.txt")

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

	return totalZeros
}

func Part2(lines []string) int {
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

	return totalZeros
}
