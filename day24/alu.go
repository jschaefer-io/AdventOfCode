package day24

import (
    "strconv"
    "strings"
)

type Argument interface {
    Value(values map[string]int) int
}

type LiteralArgument struct {
    value int
}

func (l LiteralArgument) Value(values map[string]int) int {
    return l.value
}

type RefArgument struct {
    reference string
}

func (r RefArgument) Value(values map[string]int) int {
    return values[r.reference]
}

type Instruction struct {
    name    string
    target  RefArgument
    operand Argument
}

type Alu struct {
    program []Instruction
}

func (a *Alu) Run(numbers string) map[string]int {
    pointer := 0
    values := make(map[string]int)
    for _, inst := range a.program {
        switch inst.name {
        case "inp":
            n, _ := strconv.Atoi(string(numbers[pointer]))
            values[inst.target.reference] = n
            pointer++
        case "add":
            values[inst.target.reference] = inst.target.Value(values) + inst.operand.Value(values)
        case "mul":
            values[inst.target.reference] = inst.target.Value(values) * inst.operand.Value(values)
        case "div":
            values[inst.target.reference] = inst.target.Value(values) / inst.operand.Value(values)
        case "mod":
            values[inst.target.reference] = inst.target.Value(values) % inst.operand.Value(values)
        case "eql":
            if inst.target.Value(values) == inst.operand.Value(values) {
                values[inst.target.reference] = 1
            } else {
                values[inst.target.reference] = 0
            }
        default:
            panic("unknown instruction " + inst.name)
        }
    }
    return values
}

func NewAlu(code string) Alu {
    program := make([]Instruction, 0)
    for _, line := range strings.Split(code, "\n") {
        if len(line) == 0 {
            continue
        }
        groups := strings.Split(line, " ")
        i := Instruction{
            name: groups[0],
            target: RefArgument{
                reference: groups[1],
            },
            operand: nil,
        }
        if len(groups) == 3 {
            n, err := strconv.Atoi(groups[2])
            if err != nil {
                i.operand = RefArgument{
                    reference: groups[2],
                }
            } else {
                i.operand = LiteralArgument{
                    value: n,
                }
            }
        }
        program = append(program, i)
    }

    return Alu{
        program: program,
    }
}
