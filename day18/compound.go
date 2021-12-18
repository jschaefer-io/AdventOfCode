package day18

import "fmt"

type Compound struct {
    left  Formula
    right Formula
}

func (c Compound) String() string {
    return fmt.Sprintf("[%s,%s]", c.left.String(), c.right.String())
}

func (c *Compound) add(f Formula) Formula {
    return &Compound{
        left:  c.Copy(),
        right: f.Copy(),
    }
}

func (c *Compound) explode(depth int, parents ExplodeTrace) bool {
    var left, right bool
    if depth == 4 {
        l := c.left.(*Number)
        r := c.right.(*Number)
        parents.AddToFirst(l.value, true)
        parents.AddToFirst(r.value, false)
        p := parents[len(parents)-1]
        if p.isLeft {
            p.parent.left = &Number{0}
        } else {
            p.parent.right = &Number{0}
        }
        return true
    } else {
        left = c.left.explode(depth+1, append(parents, Trace{c, true}))
        if !left {
            right = c.right.explode(depth+1, append(parents, Trace{c, false}))
        }
    }
    return left || right
}

func (c *Compound) findFirstNumber(left bool) *Number {
    if left {
        n, ok := c.left.(*Number)
        if ok {
            return n
        }
        return c.left.(*Compound).findFirstNumber(true)
    } else {
        n, ok := c.right.(*Number)
        if ok {
            return n
        }
        return c.right.(*Compound).findFirstNumber(false)
    }
}

func (c *Compound) split() (bool, Formula) {
    splitLeft, newLeft := c.left.split()
    if splitLeft {
        c.left = newLeft
    }

    splitRight := false
    if !splitLeft {
        var newRight Formula
        splitRight, newRight = c.right.split()
        if splitRight {
            c.right = newRight
        }
    }
    return splitLeft || splitRight, c
}

func (c *Compound) Magnitude() int {
    return 3*c.left.Magnitude() + 2*c.right.Magnitude()
}

func (c *Compound) Copy() Formula {
    return &Compound{
        c.left.Copy(),
        c.right.Copy(),
    }
}
