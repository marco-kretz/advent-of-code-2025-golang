package day03

import (
	"strconv"
	"strings"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

func init() {
	puzzle.Register(3, 1, Part1)
	puzzle.Register(3, 2, Part2)
}

func Part1(lines []string) (int, error) {
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

	return totalOutput, nil
}

func Part2(lines []string) (int, error) {
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

	return totalOutput, nil
}
