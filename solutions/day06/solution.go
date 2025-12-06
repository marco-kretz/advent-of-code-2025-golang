package day06

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
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
	grid := kit.AsGrid(lines)
	problems := []problem{}
	currentProblem := problem{}

	inCol := true
	// Iterate cols from right to left
	for x := len(grid[0]) - 1; x >= 0; x-- {
		// Skip the empty cols
		if !inCol && x > 0 {
			inCol = true
			continue
		}

		// Iterate rows top to bottom
		builder := strings.Builder{}
		for y := range grid {
			value := grid[y][x]

			// Hit end of column -> save the complete "col-number"
			if y == len(grid)-1 {
				num, _ := strconv.Atoi(builder.String())
				currentProblem.nums = append(currentProblem.nums, num)

				// Hit the operater -> end of "problem block" reached
				if !unicode.IsSpace(value) {
					currentProblem.op = string(value)
					problems = append(problems, currentProblem)
					currentProblem = problem{}
					inCol = false // Indicate to skip next (empty) col
					break
				}

				continue
			}

			// Hit a number
			if !unicode.IsSpace(value) {
				builder.WriteString(string(value))
			}
		}
	}

	// Solve problems and add up result
	result := 0
	for _, problem := range problems {
		result += problem.solve()
	}

	return result, nil
}

// Convert a "problem column" to a problem struct
// Only used in Part1
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
