package day06

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

func init() {
	puzzle.Register(6, 1, Part1)
	puzzle.Register(6, 2, Part2)
}

func Part1(lines []string) (int, error) {
	re := regexp.MustCompile(`\s+`)
	cols := len(re.Split(lines[0], -1))

	// Init 2D problems array
	problems := make([][]string, cols)
	for i := range problems {
		problems[i] = make([]string, len(lines))
	}

	// Fill problems array
	for rowIndex, row := range lines {
		for colIndex, cell := range re.Split(strings.TrimSpace(row), -1) {
			problems[colIndex][rowIndex] = cell
		}
	}

	result := 0
	for _, problem := range problems {
		result += solveProblem(problem)
	}

	return result, nil
}

func Part2(lines []string) (int, error) {
	return 0, nil
}

func solveProblem(problem []string) int {
	solution := 0
	operator := problem[len(problem)-1]

	for index, operandStr := range problem {
		if index == len(problem)-1 {
			break
		}

		operand, _ := strconv.Atoi(operandStr)
		if index == 0 {
			solution = operand
			continue
		}

		if operator == "+" {
			solution += operand
			continue
		}

		if operator == "*" {
			solution *= operand
		}
	}

	return solution
}
