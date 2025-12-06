package day06

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

type problem struct {
	op   string
	nums []int
}

func init() {
	puzzle.Register(6, 1, Part1)
	puzzle.Register(6, 2, Part2)
}

func Part1(lines []string) (int, error) {
	re := regexp.MustCompile(`\s+`)
	cols := len(re.Split(lines[0], -1))

	// Init 2D problems array
	problemsArr := make([][]string, cols)
	for i := range problemsArr {
		problemsArr[i] = make([]string, len(lines))
	}

	// Fill problems array
	for rowIndex, row := range lines {
		for colIndex, cell := range re.Split(strings.TrimSpace(row), -1) {
			problemsArr[colIndex][rowIndex] = cell
		}
	}

	result := 0
	for _, column := range problemsArr {
		p := toProblem(column)
		result += p.solve()
	}

	return result, nil
}

func Part2(lines []string) (int, error) {
	return 0, nil
}

// Convert a "problem column" to a problem struct
func toProblem(arr []string) problem {
	op := arr[len(arr)-1]
	nums := make([]int, len(arr)-1)

	for i := range nums {
		value, _ := strconv.Atoi(arr[i])
		nums[i] = value
	}

	return problem{
		op:   op,
		nums: nums,
	}
}

// Return a problem's solution
func (p problem) solve() int {
	solution := 0

	for i, value := range p.nums {
		if i == 0 {
			solution = value
			continue
		}

		if p.op == "+" {
			solution += value
			continue
		}

		if p.op == "*" {
			solution *= value
		}
	}

	return solution
}
