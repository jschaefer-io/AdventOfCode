package day21

type Dice interface {
    Roll() []int
    Count() int
}

type DeterministicDice struct {
    next  int
    count int
}

func (d *DeterministicDice) Count() int {
    return d.count
}

func (d *DeterministicDice) Roll() []int {
    v := d.next + 1
    value := v*3 + 3
    value %= 100
    d.next += 3 % 100
    d.count += 3
    return []int{value}
}

type DiracDice struct{}

func (d *DiracDice) Count() int {
    return 0
}

func (d *DiracDice) Roll() []int {
    list := make([]int, 0)
    for a := 1; a <= 3; a++ {
        for b := 1; b <= 3; b++ {
            for c := 1; c <= 3; c++ {
                list = append(list, a+b+c)
            }
        }
    }
    return list
}
