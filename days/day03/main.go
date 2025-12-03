package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
)

func main() {
	lines := kit.ReadFile("inputs/day03.txt")

	// Part 1
	startOne := time.Now()
	resultOne := Part1(lines)
	durationOne := time.Since(startOne)
	fmt.Println("Part 1:", resultOne, "in", durationOne)

	// Part 2
	startTwo := time.Now()
	resultTwo := Part2(lines)
	durationTwo := time.Since(startTwo)
	fmt.Println("Part 2:", resultTwo, "in", durationTwo)
}

func Part1(lines []string) int {
	totalOutput := 0

	for _, bank := range lines {
		bankArr := strings.Split(bank, "")

		// Find highest first number
		var first, firstFoundAt int = -1, -1
		for index, digitStr := range bankArr {
			digit, _ := strconv.Atoi(digitStr)

			if digit > first && index < len(bank)-1 {
				first = digit
				firstFoundAt = index
			}
		}

		// Find highest second number
		var second, secondFoundAt int = -1, -1
		for i := firstFoundAt + 1; i < len(bankArr); i++ {
			digit, _ := strconv.Atoi(bankArr[i])

			if digit > second {
				second = digit
				secondFoundAt = i
			}
		}

		// Add constructed number to total output joltage
		highestNumber, _ := strconv.Atoi(bankArr[firstFoundAt] + bankArr[secondFoundAt])
		totalOutput += highestNumber
	}

	return totalOutput
}

func Part2(lines []string) int {
	totalOutput := 0

	for _, bank := range lines {
		bankArr := strings.Split(bank, "")
		bankJoltages := [12]int{}

		lastIndex := 0
		for n := range 12 {
			highest := 0

			// Find highest n'th number
			for i := lastIndex; i < len(bankArr)-11+n; i++ {
				digit, _ := strconv.Atoi(bankArr[i])
				if digit > highest {
					highest = digit
					lastIndex = i + 1
				}
			}

			// Save highest found number
			bankJoltages[n] = highest
		}

		// Convert int array to string
		var sb strings.Builder
		for _, num := range bankJoltages {
			sb.WriteString(strconv.Itoa(num))
		}

		// Convert constructed "string-number" to int and add up
		bankJoltage, _ := strconv.Atoi(sb.String())
		totalOutput += bankJoltage
	}

	return totalOutput
}
