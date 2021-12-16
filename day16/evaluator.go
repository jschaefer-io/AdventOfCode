package day16

type Evaluator interface {
    Eval() int
    Sum() int
}

type Packet struct {
    Version int
    Type    int
}

func (p Packet) Eval() int {
    return 0
}

func (p Packet) Sum() int {
    return p.Version
}
