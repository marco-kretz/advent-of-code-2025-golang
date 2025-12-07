package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/kit"
	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"

	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day01"
	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day02"
	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day03"
	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day04"
	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day05"
	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day06"
	_ "github.com/marco-kretz/advent-of-code-2025-go/solutions/day07"
)

func main() {
	now := time.Now()
	day := flag.Int("d", now.Day(), "Day to run")
	flag.Parse()

	fmt.Printf("🎄 Advent of Code 2025 - Day %02d 🎄\n", *day)

	input := kit.ReadFile(fmt.Sprintf("inputs/day%02d.txt", *day), true)

	if err := runPart(*day, 1, input); err != nil {
		fmt.Printf("Part 1: (Not implemented or error) %v\n", err)
	}

	if err := runPart(*day, 2, input); err != nil {
		fmt.Printf("Part 2: (Not implemented or error) %v\n", err)
	}
}

func runPart(day, part int, input []string) error {
	solver := puzzle.GetSolver(day, part)

	start := time.Now()
	answer, error := solver(input)
	duration := time.Since(start)

	if error != nil {
		return error
	}

	fmt.Printf("🌟 Part %d: %d (%v)\n", part, answer, duration)
	return nil
}
