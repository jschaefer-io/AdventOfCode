package day18

import (
    "fmt"
    "math"
)

type Number struct {
    value int
}

func (n Number) String() string {
    return fmt.Sprintf("%d", n.value)
}

func (n *Number) add(f Formula) Formula {
    return &Compound{
        left:  n.Copy(),
        right: f.Copy(),
    }
}

func (n *Number) explode(depth int, parents ExplodeTrace) bool {
    return false
}

func (n *Number) split() (bool, Formula) {
    if n.value < 10 {
        return false, n
    }
    v := float64(n.value) / 2
    return true, &Compound{
        &Number{int(math.Floor(v))},
        &Number{int(math.Ceil(v))},
    }
}

func (n *Number) Magnitude() int{
    return n.value
}

func (n *Number) Copy() Formula{
    return &Number{n.value}
}
