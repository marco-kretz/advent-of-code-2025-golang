package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/marco-kretz/advent-of-code-2025-go/internal/puzzle"
)

func init() {
	puzzle.Register(5, 1, Part1)
	puzzle.Register(5, 2, Part2)
}

func Part1(lines []string) (int, error) {
	// On non 64bit systems we should use int64 explicitly
	var intervals [][2]int
	var fresh []int

	hitEmpty := false
	// Read intervals into slice
	for _, value := range lines {
		// Found the empty line
		if value == "" {
			hitEmpty = true
			continue
		}

		// Line with a "fresh" ID interval
		if !hitEmpty {
			idInterval := strings.Split(value, "-")
			if len(idInterval) == 2 {
				start, _ := strconv.Atoi(idInterval[0])
				end, _ := strconv.Atoi(idInterval[1])
				intervals = append(intervals, [2]int{start, end})
			}
			continue
		}

		// Line with an "available" ID
		id, _ := strconv.Atoi(value)
		if isFresh(intervals, id) {
			fresh = append(fresh, id)
		}
	}

	return len(fresh), nil
}

func Part2(lines []string) (int, error) {
	// On non-64bit systems we should use int64 explicitly!
	var intervals [][2]int

	// Read intervals into slice
	for _, value := range lines {
		// Found the empty line -> done
		if value == "" {
			break
		}

		// Line with a "fresh" ID interval
		idRange := strings.Split(value, "-")
		if len(idRange) == 2 {
			start, _ := strconv.Atoi(idRange[0])
			end, _ := strconv.Atoi(idRange[1])
			intervals = append(intervals, [2]int{start, end})
		}
	}

	// Sort intervals by start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Merge ranges and count the merged interval lengths
	result := 0
	for index, interval := range intervals {
		if interval[1] == -1 {
			continue
		}

		// Interval to add to the final interval list
		cleanedInterval := [2]int{interval[0], interval[1]}

		// Detect overlaps with the next n ranges
		nextToCheck := index + 1
		for nextToCheck < len(intervals) && intervals[nextToCheck][0] <= cleanedInterval[1] {
			cleanedInterval[1] = max(cleanedInterval[1], intervals[nextToCheck][1])
			// Set merged interval to -1, -1 to indicate it has been merged
			intervals[nextToCheck] = [2]int{-1, -1}
			nextToCheck++
		}

		result += cleanedInterval[1] - cleanedInterval[0] + 1
	}

	return result, nil
}

func isFresh(ranges [][2]int, id int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}

	return false
}
