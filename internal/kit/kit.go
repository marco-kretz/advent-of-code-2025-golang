package kit

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Iterate over lines
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func EuclideanModulo(dividend int, divisor int) int {
	return ((dividend % divisor) + divisor) % divisor
}
