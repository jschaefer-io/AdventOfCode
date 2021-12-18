package day18

type Trace struct {
    parent *Compound
    isLeft bool
}

type ExplodeTrace []Trace

func (e ExplodeTrace) AddToFirst(n int, left bool) {
    var num *Number
    for i := len(e) - 1; i >= 0; i-- {
        c := e[i]
        if left && !e[i].isLeft {
            t, ok := c.parent.left.(*Number)
            if ok {
                num = t
            } else {
                num = c.parent.left.(*Compound).findFirstNumber(false)
            }
        } else if !left && e[i].isLeft {
            t, ok := c.parent.right.(*Number)
            if ok {
                num = t
            } else {
                num = c.parent.right.(*Compound).findFirstNumber(true)
            }
        }
        if num != nil {
            break
        }
    }
    if num != nil {
        num.value += n
    }
}
