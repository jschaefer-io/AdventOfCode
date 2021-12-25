package day25

import (
	"github.com/jschaefer-io/aoc2021/orchestration"
	"strconv"
	"strings"
)

func Solve(data string, result *orchestration.Result) error {
	floor := Floor{
		cucumbers: make([]*Cucumber, 0),
		field:     make(map[Position]*Cucumber),
	}
	for y, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		floor.height++
		floor.width = 0
		for x, c := range line {
			floor.width++
			if c == '.' {
				continue
			}
			pos := Position{x, y}
			cu := &Cucumber{
				down:     c == 'v',
				Position: pos,
			}
			floor.cucumbers = append(floor.cucumbers, cu)
			floor.field[pos] = cu
		}
	}

	stepCound := 0
	for {
		stepCound++
		if !floor.Step() {
			break
		}
	}
	result.AddResult(strconv.Itoa(stepCound))
	return nil
}

func init() {
	orchestration.MainDispatcher.AddSolver("Day25", orchestration.NewSolver(Solve))
}
