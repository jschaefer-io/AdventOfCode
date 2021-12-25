package day25

import (
	"github.com/jschaefer-io/aoc2021/orchestration"
	"strconv"
	"strings"
)

type Cucumber struct {
	down     bool
	Position Position
}

func (c *Cucumber) NextPosition() Position {
	if c.down {
		return Position{c.Position.X, c.Position.Y + 1}
	} else {
		return Position{c.Position.X + 1, c.Position.Y}
	}
}

type Position struct {
	X int
	Y int
}

type Floor struct {
	width     int
	height    int
	cucumbers []*Cucumber
	field     map[Position]*Cucumber
}

func (f Floor) String() string {
	var sb strings.Builder
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			c, ok := f.field[Position{x, y}]
			if ok {
				if c.down {
					sb.WriteString("v")
				} else {
					sb.WriteString(">")
				}
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (f *Floor) updatePositions(down bool) bool {
	updates := make([][2]Position, 0)
	for _, cuc := range f.cucumbers {
		if cuc.down != down {
			continue
		}
		nextPos := f.FitPosition(cuc.NextPosition())
		if _, ok := f.field[nextPos]; !ok {
			updates = append(updates, [2]Position{cuc.Position, nextPos})
		}
	}
	for _, update := range updates {
		cuc := f.field[update[0]]
		cuc.Position = update[1]
		delete(f.field, update[0])
		f.field[update[1]] = cuc
	}
	return len(updates) > 0
}

func (f *Floor) Step() bool {
	east := f.updatePositions(false)
	south := f.updatePositions(true)
	return east || south
}

func (f *Floor) FitPosition(pos Position) Position {
	return Position{
		X: pos.X % f.width,
		Y: pos.Y % f.height,
	}
}

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
