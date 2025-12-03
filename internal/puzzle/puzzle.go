package puzzle

import (
	"fmt"
)

type Solver func(input []string) (int, error)

var Registry = map[string]Solver{}

func Register(day, part int, fn Solver) {
	key := fmt.Sprintf("%02d-%d", day, part)
	Registry[key] = fn
}

func GetSolver(day, part int) Solver {
	key := fmt.Sprintf("%02d-%d", day, part)
	value, ok := Registry[key]
	if !ok {
		panic("No solver found")
	}
	return value
}
